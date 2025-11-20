package diode

import (
	"context"
	"fmt"
	"net"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	logsservicepb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type mockLogsService struct {
	logsservicepb.UnimplementedLogsServiceServer

	mu        sync.Mutex
	requests  []*logsservicepb.ExportLogsServiceRequest
	metadata  []metadata.MD
	failWith  error
	failAfter int
	counter   int
}

func (m *mockLogsService) Export(ctx context.Context, req *logsservicepb.ExportLogsServiceRequest) (*logsservicepb.ExportLogsServiceResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.counter++
	if m.failWith != nil && (m.failAfter == 0 || m.counter >= m.failAfter) {
		return nil, m.failWith
	}

	m.requests = append(m.requests, req)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		m.metadata = append(m.metadata, md)
	}

	return &logsservicepb.ExportLogsServiceResponse{}, nil
}

func startMockOTLPServer(t *testing.T, svc logsservicepb.LogsServiceServer, opts ...grpc.ServerOption) (string, func()) {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	server := grpc.NewServer(opts...)
	logsservicepb.RegisterLogsServiceServer(server, svc)

	go func() {
		_ = server.Serve(listener)
	}()

	cleanup := func() {
		server.GracefulStop()
		_ = listener.Close()
	}

	return fmt.Sprintf("grpc://%s", listener.Addr().String()), cleanup
}

func TestOTLPClientExportsEntitiesAsLogs(t *testing.T) {
	service := &mockLogsService{}
	target, cleanup := startMockOTLPServer(t, service)
	defer cleanup()

	client, err := NewOTLPClient(target, "otlp-producer", "1.2.3")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, client.Close())
	})

	ctx := context.Background()
	resp, err := client.Ingest(ctx, []Entity{
		&Site{Name: String("Test Site")},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	service.mu.Lock()
	defer service.mu.Unlock()
	require.Len(t, service.requests, 1)

	req := service.requests[0]
	require.Len(t, req.ResourceLogs, 1)

	resourceLogs := req.ResourceLogs[0]
	require.NotNil(t, resourceLogs.Resource)
	resourceAttr := attributesToMap(resourceLogs.Resource.Attributes)
	assert.Equal(t, "otlp-producer", resourceAttr["service.name"])
	assert.Equal(t, "1.2.3", resourceAttr["service.version"])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, resourceAttr["os.description"])
	assert.Equal(t, runtime.Version(), resourceAttr["process.runtime.version"])
	assert.Equal(t, defaultStreamName, resourceAttr["diode.stream"])
	assert.NotContains(t, resourceAttr, "telemetry.sdk.name")

	require.Len(t, resourceLogs.ScopeLogs, 1)
	scopeLogs := resourceLogs.ScopeLogs[0]
	require.NotNil(t, scopeLogs.Scope)
	assert.Equal(t, otlpClientName, scopeLogs.Scope.Name)
	assert.Equal(t, getSDKVersion(), scopeLogs.Scope.Version)

	require.Len(t, scopeLogs.LogRecords, 1)
	record := scopeLogs.LogRecords[0]

	assert.Equal(t, logspb.SeverityNumber_SEVERITY_NUMBER_INFO, record.SeverityNumber)
	assert.Equal(t, "INFO", record.SeverityText)

	body := record.Body.GetStringValue()
	assert.Contains(t, body, `"site"`)
	assert.Contains(t, body, `"Test Site"`)

	attrMap := attributesToMap(record.Attributes)
	assert.Equal(t, "site", attrMap["diode.entity"])
	assert.NotContains(t, attrMap, "diode.stream")
}

func TestOTLPClientWrapsExportErrors(t *testing.T) {
	service := &mockLogsService{
		failWith: status.Error(codes.Internal, "collector exploded"),
	}
	target, cleanup := startMockOTLPServer(t, service)
	defer cleanup()

	client, err := NewOTLPClient(target, "otlp-producer", "1.2.3")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, client.Close())
	})

	_, err = client.Ingest(context.Background(), []Entity{
		&Site{Name: String("Test Site")},
	})
	require.Error(t, err)

	var otlpErr *OTLPClientError
	require.ErrorAs(t, err, &otlpErr)
	assert.Equal(t, codes.Internal, otlpErr.StatusCode)
	assert.Contains(t, otlpErr.Error(), "OTLP export failed")
	assert.Contains(t, otlpErr.Details, "collector exploded")
}

func attributesToMap(attributes []*commonpb.KeyValue) map[string]string {
	result := make(map[string]string, len(attributes))
	for _, attr := range attributes {
		if attr == nil || attr.Value == nil {
			continue
		}
		result[attr.Key] = attr.Value.GetStringValue()
	}
	return result
}

func TestOTLPClientIngestWithMetadata(t *testing.T) {
	service := &mockLogsService{}
	target, cleanup := startMockOTLPServer(t, service)
	defer cleanup()

	client, err := NewOTLPClient(target, "otlp-producer", "1.2.3")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, client.Close())
	})

	ctx := context.Background()

	md := Metadata{
		"batch_id": "batch-999",
		"source":   "otlp_test",
		"priority": 5,
		"enabled":  true,
		// Additional type tests
		"int64_value":   int64(9223372036854775807),
		"float_value":   3.14159,
		"bool_false":    false,
		"empty_string":  "",
		"zero_int":      0,
		"zero_float":    0.0,
		"array_value":   []any{"item1", "item2", 123},
		"nested_object": map[string]any{"key": "value"},
		"nil_value":     nil,
	}

	resp, err := client.Ingest(ctx, []Entity{
		&Site{Name: String("Test Site"), Metadata: Metadata{"foo": "bar"}},
	}, WithIngestMetadata(md))
	require.NoError(t, err)
	require.NotNil(t, resp)

	service.mu.Lock()
	require.Len(t, service.requests, 1)

	req := service.requests[0]
	require.Len(t, req.ResourceLogs, 1)

	resourceLogs := req.ResourceLogs[0]
	require.NotNil(t, resourceLogs.Resource)

	// Verify standard resource attributes
	resourceAttr := attributesToMap(resourceLogs.Resource.Attributes)
	assert.Equal(t, "otlp-producer", resourceAttr["service.name"])
	assert.Equal(t, "1.2.3", resourceAttr["service.version"])

	// Verify metadata is added to resource attributes with diode.metadata prefix
	assert.Equal(t, "batch-999", resourceAttr["diode.metadata.batch_id"])
	assert.Equal(t, "otlp_test", resourceAttr["diode.metadata.source"])

	// Verify all metadata types are handled correctly
	foundTypes := make(map[string]bool)
	for _, attr := range resourceLogs.Resource.Attributes {
		switch attr.Key {
		case "diode.metadata.priority":
			assert.Equal(t, int64(5), attr.Value.GetIntValue())
			foundTypes["priority"] = true
		case "diode.metadata.enabled":
			assert.Equal(t, true, attr.Value.GetBoolValue())
			foundTypes["enabled"] = true
		case "diode.metadata.int64_value":
			assert.Equal(t, int64(9223372036854775807), attr.Value.GetIntValue())
			foundTypes["int64"] = true
		case "diode.metadata.float_value":
			assert.InDelta(t, 3.14159, attr.Value.GetDoubleValue(), 0.00001)
			foundTypes["float"] = true
		case "diode.metadata.bool_false":
			assert.Equal(t, false, attr.Value.GetBoolValue())
			foundTypes["bool_false"] = true
		case "diode.metadata.zero_int":
			assert.Equal(t, int64(0), attr.Value.GetIntValue())
			foundTypes["zero_int"] = true
		case "diode.metadata.zero_float":
			assert.Equal(t, float64(0), attr.Value.GetDoubleValue())
			foundTypes["zero_float"] = true
		}
	}

	// Verify all expected primitive types were found
	assert.True(t, foundTypes["priority"], "int metadata should be present")
	assert.True(t, foundTypes["enabled"], "bool true metadata should be present")
	assert.True(t, foundTypes["int64"], "int64 metadata should be present")
	assert.True(t, foundTypes["float"], "float metadata should be present")
	assert.True(t, foundTypes["bool_false"], "bool false metadata should be present")
	assert.True(t, foundTypes["zero_int"], "zero int metadata should be present")
	assert.True(t, foundTypes["zero_float"], "zero float metadata should be present")

	// Verify string types (including empty string)
	assert.Equal(t, "", resourceAttr["diode.metadata.empty_string"])

	// Verify complex types are converted to strings (using fmt.Sprintf("%v", value))
	assert.Contains(t, resourceAttr["diode.metadata.array_value"], "item1")
	assert.Contains(t, resourceAttr["diode.metadata.array_value"], "item2")
	assert.Contains(t, resourceAttr["diode.metadata.nested_object"], "key")
	assert.Contains(t, resourceAttr["diode.metadata.nested_object"], "value")

	// Verify nil is converted to string "<nil>"
	assert.Equal(t, "<nil>", resourceAttr["diode.metadata.nil_value"])

	service.mu.Unlock()

	resp, err = client.Ingest(ctx, []Entity{
		&Site{Name: String("Test Site 2")},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	service.mu.Lock()
	defer service.mu.Unlock()
	require.Len(t, service.requests, 2)

	req2 := service.requests[1]
	resourceAttr2 := attributesToMap(req2.ResourceLogs[0].Resource.Attributes)
	assert.NotContains(t, resourceAttr2, "diode.metadata.batch_id")
	assert.NotContains(t, resourceAttr2, "diode.metadata.source")
}
