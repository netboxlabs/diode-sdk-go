package diode

import (
	"context"
	"crypto/tls"
	"fmt"
	"runtime"
	"strings"
	"time"

	logsservicepb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	otlpClientName        = "diode-sdk-go-otlp"
	defaultOTLPTimeout    = 10 * time.Second
	otlpExportErrorPrefix = "OTLP export failed"
)

// OTLPClientOption configures the OTLPClient.
type OTLPClientOption func(*OTLPClient) error

// OTLPClient implements Client and exports entities as OTLP log records.
type OTLPClient struct {
	appName    string
	appVersion string

	timeout time.Duration
	stream  string

	certFile string

	metadata metadata.MD

	conn   *grpc.ClientConn
	client logsservicepb.LogsServiceClient

	sdkName    string
	sdkVersion string

	platform  string
	goVersion string
}

// WithOTLPTimeout overrides the default export timeout.
func WithOTLPTimeout(timeout time.Duration) OTLPClientOption {
	return func(c *OTLPClient) error {
		if timeout <= 0 {
			return fmt.Errorf("OTLP timeout must be greater than zero")
		}
		c.timeout = timeout
		return nil
	}
}

// WithOtlpCertFile configures the certificate file to trust for secure endpoints.
func WithOtlpCertFile(certFile string) OTLPClientOption {
	return func(c *OTLPClient) error {
		c.certFile = certFile
		return nil
	}
}

// WithOtlpStream overrides the default stream value associated with exported entities.
func WithOtlpStream(stream string) OTLPClientOption {
	return func(c *OTLPClient) error {
		if strings.TrimSpace(stream) == "" {
			return fmt.Errorf("OTLP stream must not be empty")
		}
		c.stream = stream
		return nil
	}
}

// NewOTLPClient creates a new OTLPClient that exports entities as OTLP log records.
func NewOTLPClient(target, appName, appVersion string, opts ...OTLPClientOption) (Client, error) {
	if target == "" {
		return nil, fmt.Errorf("target is required")
	}

	if !strings.HasPrefix(target, "grpc://") && !strings.HasPrefix(target, "grpcs://") {
		return nil, fmt.Errorf("OTLPClient target should start with grpc:// or grpcs://")
	}

	authority, path, isPlaintext, tlsVerify, err := parseTarget(target)
	if err != nil {
		return nil, err
	}

	client := &OTLPClient{
		appName:    appName,
		appVersion: appVersion,
		timeout:    defaultOtlpTimeout,
		stream:     defaultStreamName,
		sdkName:    otlpClientName,
		sdkVersion: getSDKVersion(),
		platform:   runtime.GOOS + "/" + runtime.GOARCH,
		goVersion:  runtime.Version(),
	}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	if client.certFile == "" {
		client.certFile = getCertFile("")
	}

	conn, err := client.dial(authority, path, isPlaintext, tlsVerify)
	if err != nil {
		return nil, err
	}
	client.conn = conn
	client.client = logsservicepb.NewLogsServiceClient(conn)

	return client, nil
}

func (c *OTLPClient) dial(authority, path string, isPlaintext bool, tlsVerify bool) (*grpc.ClientConn, error) {
	var dialOpts []grpc.DialOption

	userAgent := fmt.Sprintf("%s/%s %s/%s", c.sdkName, c.sdkVersion, c.appName, c.appVersion)
	dialOpts = append(dialOpts, grpc.WithUserAgent(userAgent))

	if path != "" {
		dialOpts = append(dialOpts, methodUnaryInterceptor(path))
	}

	var transportCreds credentials.TransportCredentials

	if isPlaintext {
		transportCreds = insecure.NewCredentials()
	} else {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: !tlsVerify,
		}

		// Only attempt to load certificates when verification is enabled or a custom file is provided.
		if tlsVerify || c.certFile != "" {
			rootCAs, err := loadCerts(c.certFile)
			if err != nil {
				return nil, fmt.Errorf("failed to load certificates: %w", err)
			}
			tlsConfig.RootCAs = rootCAs
		}

		transportCreds = credentials.NewTLS(tlsConfig)
	}

	dialOpts = append(dialOpts, grpc.WithTransportCredentials(transportCreds))

	return grpc.NewClient(authority, dialOpts...)
}

// Close closes the underlying gRPC connection.
func (c *OTLPClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Ingest converts the provided entities to proto messages before exporting them.
func (c *OTLPClient) Ingest(ctx context.Context, entities []Entity) (*diodepb.IngestResponse, error) {
	return c.IngestProto(ctx, convertEntitiesToProto(entities))
}

// IngestProto exports proto entities as OTLP log records.
func (c *OTLPClient) IngestProto(ctx context.Context, entities []*diodepb.Entity) (*diodepb.IngestResponse, error) {
	logRecords := make([]*logspb.LogRecord, 0, len(entities))
	for _, entity := range entities {
		if entity == nil {
			continue
		}
		record, err := c.entityToLogRecord(entity)
		if err != nil {
			return nil, err
		}
		logRecords = append(logRecords, record)
	}

	if len(logRecords) == 0 {
		return &diodepb.IngestResponse{}, nil
	}

	request := c.buildExportRequest(logRecords)

	exportCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	if len(c.metadata) > 0 {
		exportCtx = metadata.NewOutgoingContext(exportCtx, c.metadata)
	}

	if _, err := c.client.Export(exportCtx, request); err != nil {
		return nil, newOTLPClientError(err, otlpExportErrorPrefix)
	}

	return &diodepb.IngestResponse{}, nil
}

func (c *OTLPClient) entityToLogRecord(entity *diodepb.Entity) (*logspb.LogRecord, error) {
	body, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(entity)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal entity: %w", err)
	}
	entityType := c.resolveEntityType(entity)
	return &logspb.LogRecord{
		TimeUnixNano:   uint64(time.Now().UTC().UnixNano()),
		SeverityNumber: logspb.SeverityNumber_SEVERITY_NUMBER_INFO,
		SeverityText:   "INFO",
		Body: &commonpb.AnyValue{
			Value: &commonpb.AnyValue_StringValue{
				StringValue: string(body),
			},
		},
		Attributes: []*commonpb.KeyValue{
			stringKV("diode.entity", entityType),
		},
	}, nil
}

func (c *OTLPClient) buildExportRequest(logRecords []*logspb.LogRecord) *logsservicepb.ExportLogsServiceRequest {
	return &logsservicepb.ExportLogsServiceRequest{
		ResourceLogs: []*logspb.ResourceLogs{
			{
				Resource: &resourcepb.Resource{
					Attributes: c.resourceAttributes(),
				},
				ScopeLogs: []*logspb.ScopeLogs{
					{
						Scope: &commonpb.InstrumentationScope{
							Name:    c.sdkName,
							Version: c.sdkVersion,
						},
						LogRecords: logRecords,
					},
				},
			},
		},
	}
}

func (c *OTLPClient) resourceAttributes() []*commonpb.KeyValue {
	return []*commonpb.KeyValue{
		stringKV("service.name", c.appName),
		stringKV("service.version", c.appVersion),
		stringKV("os.description", c.platform),
		stringKV("process.runtime.version", c.goVersion),
		stringKV("diode.stream", c.stream),
	}
}

func (c *OTLPClient) resolveEntityType(entity *diodepb.Entity) string {
	message := entity.ProtoReflect()
	oneofs := message.Descriptor().Oneofs()
	if oneofs.Len() == 0 {
		return "unknown"
	}

	field := message.WhichOneof(oneofs.Get(0))
	if field == nil {
		return "unknown"
	}

	return string(field.Name())
}

func stringKV(key, value string) *commonpb.KeyValue {
	return &commonpb.KeyValue{
		Key: key,
		Value: &commonpb.AnyValue{
			Value: &commonpb.AnyValue_StringValue{
				StringValue: value,
			},
		},
	}
}

// OTLPClientError represents a failure while exporting log records.
type OTLPClientError struct {
	Message    string
	StatusCode codes.Code
	Details    string
	Err        error
}

func newOTLPClientError(err error, message string) *OTLPClientError {
	e := &OTLPClientError{
		Message: message,
		Err:     err,
	}

	if statusErr, ok := status.FromError(err); ok {
		e.StatusCode = statusErr.Code()
		e.Details = statusErr.Message()
	} else {
		e.Details = err.Error()
	}

	return e
}

// Error implements the error interface.
func (e *OTLPClientError) Error() string {
	parts := []string{e.Message}

	if e.StatusCode != codes.OK && e.StatusCode != codes.Unknown {
		parts = append(parts, fmt.Sprintf("status=%s", e.StatusCode.String()))
	}

	if e.Details != "" {
		parts = append(parts, fmt.Sprintf("details=%s", e.Details))
	}

	return strings.Join(parts, ", ")
}

// Unwrap exposes the underlying error.
func (e *OTLPClientError) Unwrap() error {
	return e.Err
}
