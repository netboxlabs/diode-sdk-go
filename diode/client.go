package diode

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// SDKName is the name of the Diode SDK
	SDKName = "diode-sdk-go"

	// SDKVersion is the version of the Diode SDK
	SDKVersion = "0.2.0"

	// DiodeClientIDEnvVarName is the environment variable name for the Diode Client ID
	DiodeClientIDEnvVarName = "DIODE_CLIENT_ID"

	// DiodeClientSecretEnvVarName is the environment variable name for the Diode Client Secret
	DiodeClientSecretEnvVarName = "DIODE_CLIENT_SECRET"

	// DiodeSDKLogLevelEnvVarName is the environment variable name for the Diode SDK log level
	DiodeSDKLogLevelEnvVarName = "DIODE_SDK_LOG_LEVEL"

	// DiodeMaxAuthRetriesEnvVarName is the environment variable name for the maximum number of authentication retries
	DiodeMaxAuthRetriesEnvVarName = "DIODE_MAX_AUTH_RETRIES"

	// DiodeOAuth2IngestScope is the OAuth2 scope for the data ingestion
	DiodeOAuth2IngestScope = "diode:ingest"

	defaultStreamName = "latest"
)

var allowedSchemesRe = regexp.MustCompile(`grpc|grpcs|http|https`)

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
		return "", "", false, errors.New("target should start with grpc:// or grpcs:// or http:// or https://")
	}

	authority := u.Host
	if u.Port() == "" {
		switch u.Scheme {
		case "grpc", "http":
			authority += ":80"
		case "grpcs", "https":
			authority += ":443"
		default:
			return "", "", false, fmt.Errorf("missing port with unsupported scheme: %s", u.Scheme)
		}
	}

	path := u.Path
	if path == "/" {
		path = ""
	}

	tlsVerify := u.Scheme == "grpcs"

	return authority, path, tlsVerify, nil
}

// getClientID returns the client ID either from provided value or environment variable
func getClientID(clientID string) (string, error) {
	if clientID == "" {
		clientID = os.Getenv(DiodeClientIDEnvVarName)
	}

	if clientID == "" {
		return "", fmt.Errorf("client_id param or %s environment variable required", DiodeClientIDEnvVarName)
	}

	return clientID, nil
}

// getClientSecret returns the client secret either from provided value or environment variable
func getClientSecret(clientSecret string) (string, error) {
	if clientSecret == "" {
		clientSecret = os.Getenv(DiodeClientSecretEnvVarName)
	}

	if clientSecret == "" {
		return "", fmt.Errorf("client_secret param or %s environment variable required", DiodeClientSecretEnvVarName)
	}

	return clientSecret, nil
}

// getAuthRetries returns the maximum number of authentication retries
func getAuthRetries(maxAuthRetries int) (int, error) {
	maxAuthRetriesStr := os.Getenv(DiodeMaxAuthRetriesEnvVarName)
	if maxAuthRetriesStr != "" {
		retries, err := strconv.Atoi(maxAuthRetriesStr)
		if err != nil {
			return 0, fmt.Errorf("invalid value for %s: %w", DiodeMaxAuthRetriesEnvVarName, err)
		}
		maxAuthRetries = retries
	}
	if maxAuthRetries <= 0 {
		return 0, fmt.Errorf("max_auth_retries param or %s environment variable must be greater than 0", DiodeMaxAuthRetriesEnvVarName)
	}
	return maxAuthRetries, nil
}

// Client is an interface that defines the methods available from Diode API
type Client interface {
	// Close closes the connection to the API service
	Close() error

	// Ingest sends an ingest request to the ingester service
	Ingest(context.Context, []Entity) (*diodepb.IngestResponse, error)

	// IngestProto sends an ingest request to the ingester service with proto entities
	IngestProto(context.Context, []*diodepb.Entity) (*diodepb.IngestResponse, error)
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

	// The client ID for the API
	clientID string

	// The client secret for the API
	clientSecret string

	// The maximum number of authentication retries
	maxAuthRetries int

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

// WithClientID sets the client ID for the GRPCClient
func WithClientID(clientID string) ClientOption {
	return func(c *GRPCClient) {
		c.clientID = clientID
	}
}

// WithClientSecret sets the client secret for the GRPCClient
func WithClientSecret(clientSecret string) ClientOption {
	return func(c *GRPCClient) {
		c.clientSecret = clientSecret
	}
}

// authenticate fetches an OAuth2 token using client credentials and updates the metadata with the token.
func (g *GRPCClient) authenticate() error {
	authClient := newDiodeAuthentication(g.target, g.path, g.tlsVerify, g.clientID, g.clientSecret)
	accessToken, err := authClient.authenticate(g.logger, []string{DiodeOAuth2IngestScope})
	if err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	// Update metadata with the new authorization token
	g.metadata.Set("authorization", fmt.Sprintf("Bearer %s", accessToken))
	return nil
}

// DiodeAuthentication handles OAuth2 authentication for the Diode API.
type diodeAuthentication struct {
	target       string
	path         string
	tlsVerify    bool
	clientID     string
	clientSecret string
}

// NewDiodeAuthentication creates a new instance of DiodeAuthentication.
func newDiodeAuthentication(target string, path string, tlsVerify bool, clientID, clientSecret string) *diodeAuthentication {
	return &diodeAuthentication{
		target:       target,
		path:         path,
		tlsVerify:    tlsVerify,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// Authenticate requests an OAuth2 token using client credentials and returns it.
func (d *diodeAuthentication) authenticate(logger *slog.Logger, scopes []string) (string, error) {
	scheme := "http"
	if d.tlsVerify {
		scheme = "https"
	}
	authURL := fmt.Sprintf("%s://%s/auth/token", scheme, d.target)
	if d.path != "" {
		authURL = fmt.Sprintf("%s://%s%s/auth/token", scheme, d.target, d.path)
	}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", d.clientID)
	data.Set("client_secret", d.clientSecret)
	data.Set("scope", strings.Join(scopes, " "))
	req, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	if !d.tlsVerify {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Error("failed to close response body", "error", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authentication failed: %s", resp.Status)
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if result.AccessToken == "" {
		return "", errors.New("access token not found in response")
	}

	return result.AccessToken, nil
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
		logger:         logger,
		conn:           conn,
		client:         diodepb.NewIngesterServiceClient(conn),
		appName:        appName,
		appVersion:     appVersion,
		target:         target,
		path:           path,
		tlsVerify:      tlsVerify,
		platform:       platform,
		goVersion:      goVersion,
		maxAuthRetries: 3,
	}

	var clientID string
	var clientSecret string

	for _, o := range opts {
		o(c)
	}

	c.metadata = metadata.Pairs("platform", platform, "go-version", goVersion)

	c.maxAuthRetries, err = getAuthRetries(c.maxAuthRetries)
	if err != nil {
		return nil, err
	}

	clientID, err = getClientID(c.clientID)
	if err != nil {
		return nil, err
	}
	clientSecret, err = getClientSecret(c.clientSecret)
	if err != nil {
		return nil, err
	}

	c.clientID = clientID
	c.clientSecret = clientSecret

	if err = c.authenticate(); err != nil {
		return nil, err
	}

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
	return g.IngestProto(ctx, convertEntitiesToProto(entities))
}

// IngestProto sends an ingest request to the ingester service with proto entities
func (g *GRPCClient) IngestProto(ctx context.Context, entities []*diodepb.Entity) (*diodepb.IngestResponse, error) {
	stream := defaultStreamName

	req := &diodepb.IngestRequest{
		Id:                 uuid.NewString(),
		Entities:           entities,
		Stream:             stream,
		ProducerAppName:    g.appName,
		ProducerAppVersion: g.appVersion,
		SdkName:            SDKName,
		SdkVersion:         SDKVersion,
	}

	ctx = metadata.NewOutgoingContext(ctx, g.metadata)

	var err error
	var res *diodepb.IngestResponse

	attempt := 0
	for {
		res, err = g.client.Ingest(ctx, req)
		if err != nil {
			if status.Code(err) == codes.Unauthenticated {
				attempt++
				if attempt >= g.maxAuthRetries {
					return nil, fmt.Errorf("authentication failed after %d attempts: %w", attempt, err)
				}
				g.logger.Debug("Authentication failed, retrying...", "attempt", attempt)
				if err := g.authenticate(); err != nil {
					g.logger.Error("Failed to re-authenticate", "error", err)
				}
				continue
			}
			return nil, err
		}
		break
	}
	return res, nil
}

// convertEntitiesToProto converts entities to proto entities
func convertEntitiesToProto(entities []Entity) []*diodepb.Entity {
	protoEntities := make([]*diodepb.Entity, 0)
	for _, entity := range entities {
		entityPb := entity.ConvertToProtoEntity()
		entityPb.Timestamp = timestamppb.New(time.Now().UTC())
		protoEntities = append(protoEntities, entityPb)
	}
	return protoEntities
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
