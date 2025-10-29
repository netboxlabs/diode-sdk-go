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

func TestNewOTLPClientRejectsUnsupportedScheme(t *testing.T) {
	_, err := NewOTLPClient("http://localhost:4318", "app", "1.0.0")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OTLPClient target should start with grpc:// or grpcs://")
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
