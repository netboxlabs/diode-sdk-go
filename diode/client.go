package diode

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// SDKName is the name of the Diode SDK
	SDKName = "diode-sdk-go"

	// SDKVersion is the version of the Diode SDK
	SDKVersion = "0.2.0"

	// DiodeAPIKeyEnvVarName is the environment variable name for the Diode API key
	DiodeAPIKeyEnvVarName = "DIODE_API_KEY"

	// DiodeSDKLogLevelEnvVarName is the environment variable name for the Diode SDK log level
	DiodeSDKLogLevelEnvVarName = "DIODE_SDK_LOG_LEVEL"

	defaultStreamName = "latest"

	authAPIKeyName = "diode-api-key"
)

var allowedSchemesRe = regexp.MustCompile(`grpc|grpcs`)

// loadCerts loads the system x509 cert pool
func loadCerts() *x509.CertPool {
	certPool, _ := x509.SystemCertPool()
	return certPool
}

// parseTarget parses the target string into authority, path, and tlsVerify
func parseTarget(target string) (string, string, bool, error) {
	u, err := url.Parse(target)
	if err != nil {
		return "", "", false, err
	}

	if !allowedSchemesRe.MatchString(u.Scheme) {
		return "", "", false, errors.New("target should start with grpc:// or grpcs://")
	}

	authority := u.Host
	if u.Port() == "" {
		authority += ":443"
	}

	path := u.Path
	if path == "/" {
		path = ""
	}

	tlsVerify := u.Scheme == "grpcs"

	return authority, path, tlsVerify, nil
}

// getAPIKey returns the API key either from provided value or environment variable
func getAPIKey(apiKey string) (string, error) {
	if apiKey == "" {
		apiKey = os.Getenv(DiodeAPIKeyEnvVarName)
	}

	if apiKey == "" {
		return "", fmt.Errorf("api_key param or %s environment variable required", DiodeAPIKeyEnvVarName)
	}

	return apiKey, nil
}

// Client is an interface that defines the methods available from Diode API
type Client interface {
	// Close closes the connection to the API service
	Close() error

	// Ingest sends an ingest request to the ingester service
	Ingest(context.Context, []Entity) (*diodepb.IngestResponse, error)
}

// GRPCClient is a gRPC implementation of the ingester service
type GRPCClient struct {
	// The logger for the client
	logger *slog.Logger

	// gRPC virtual connection
	conn *grpc.ClientConn

	// The gRPC API client
	client diodepb.IngesterServiceClient

	// Producer's application name
	appName string

	// Producer's application version
	appVersion string

	// An API key for the Diode API
	apiKey string

	// GRPC target
	target string

	// GRPC path
	path string

	// TLS verify
	tlsVerify bool

	// Platform name
	platform string

	// Go version
	goVersion string

	// Metadata
	metadata metadata.MD
}

// ClientOption is a functional option for the GRPCClient
type ClientOption func(*GRPCClient)

// WithAPIKey sets the API key for the client
func WithAPIKey(apiKey string) ClientOption {
	return func(c *GRPCClient) {
		c.apiKey = apiKey
	}
}

// NewClient creates a new diode client based on gRPC
func NewClient(target string, appName string, appVersion string, opts ...ClientOption) (Client, error) {
	logger := newLogger()

	if appName == "" {
		return nil, fmt.Errorf("app name is required")
	}

	if appVersion == "" {
		return nil, fmt.Errorf("app version is required")
	}

	target, path, tlsVerify, err := parseTarget(target)
	if err != nil {
		return nil, err
	}

	dialOpts := []grpc.DialOption{
		grpc.WithUserAgent(userAgent()),
	}

	if path != "" {
		logger.Debug("Setting up gRPC interceptor for path", "path", path)
		dialOpts = append(dialOpts, methodUnaryInterceptor(path))
	}

	if tlsVerify {
		logger.Debug("Setting up gRPC secure channel")
		rootCAs := loadCerts()
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: rootCAs})))
	} else {
		logger.Debug("Setting up gRPC insecure channel")
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(target, dialOpts...)
	if err != nil {
		return nil, err
	}

	platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	goVersion := runtime.Version()

	c := &GRPCClient{
		logger:     logger,
		conn:       conn,
		client:     diodepb.NewIngesterServiceClient(conn),
		appName:    appName,
		appVersion: appVersion,
		target:     target,
		path:       path,
		tlsVerify:  tlsVerify,
		platform:   platform,
		goVersion:  goVersion,
	}

	var apiKey string

	for _, o := range opts {
		o(c)
	}

	apiKey, err = getAPIKey(c.apiKey)
	if err != nil {
		return nil, err
	}

	c.apiKey = apiKey
	c.metadata = metadata.Pairs(authAPIKeyName, c.apiKey, "platform", platform, "go-version", goVersion)

	return c, nil
}

// Close closes the connection to the API service
func (g *GRPCClient) Close() error {
	if g.conn != nil {
		return g.conn.Close()
	}
	return nil
}

// Ingest sends an ingest request to the ingester service
func (g *GRPCClient) Ingest(ctx context.Context, entities []Entity) (*diodepb.IngestResponse, error) {
	stream := defaultStreamName

	protoEntities := make([]*diodepb.Entity, 0)
	for _, entity := range entities {
		protoEntities = append(protoEntities, entity.ConvertToProtoEntity())
	}

	req := &diodepb.IngestRequest{
		Id:                 uuid.NewString(),
		Entities:           protoEntities,
		Stream:             stream,
		ProducerAppName:    g.appName,
		ProducerAppVersion: g.appVersion,
		SdkName:            SDKName,
		SdkVersion:         SDKVersion,
	}

	ctx = metadata.NewOutgoingContext(ctx, g.metadata)

	return g.client.Ingest(ctx, req)
}

// methodUnaryInterceptor returns a gRPC dial option with a unary interceptor
//
// It's used to intercept the client calls and modify the method details.
//
// Diode's default method generated from Protocol Buffers definition is /diode.v1.IngesterService/Ingest and in order
// to use Diode targets with path (i.e. localhost:8081/this/is/custom/path), this interceptor is used to modify the
// method details, by prepending the generated method name with the path extracted from initial target.
func methodUnaryInterceptor(path string) grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		method = fmt.Sprintf("%s%s", path, method)
		return invoker(ctx, method, req, reply, cc, opts...)
	})
}

// userAgent returns the user agent string for the SDK
func userAgent() string {
	return fmt.Sprintf("%s/%s", SDKName, SDKVersion)
}

// newLogger creates a new logger for the SDK
func newLogger() *slog.Logger {
	level, ok := os.LookupEnv(DiodeSDKLogLevelEnvVarName)
	if !ok {
		level = "INFO"
	}

	var l slog.Level
	switch strings.ToUpper(level) {
	case "DEBUG":
		l = slog.LevelDebug
	case "INFO":
		l = slog.LevelInfo
	case "WARN":
		l = slog.LevelWarn
	case "ERROR":
		l = slog.LevelError
	default:
		l = slog.LevelDebug
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l, AddSource: false})

	return slog.New(h)
}
