package diode

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"runtime/debug"
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
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// SDKName is the name of the Diode SDK
	SDKName = "diode-sdk-go"

	// DiodeCertFileEnvVarName is the environment variable name for the custom certificate file path
	DiodeCertFileEnvVarName = "DIODE_CERT_FILE"

	// DiodeClientIDEnvVarName is the environment variable name for the Diode Client ID
	DiodeClientIDEnvVarName = "DIODE_CLIENT_ID"

	// DiodeClientSecretEnvVarName is the environment variable name for the Diode Client Secret
	DiodeClientSecretEnvVarName = "DIODE_CLIENT_SECRET"

	// DiodeMaxAuthRetriesEnvVarName is the environment variable name for the maximum number of authentication retries
	DiodeMaxAuthRetriesEnvVarName = "DIODE_MAX_AUTH_RETRIES"

	// DiodeOAuth2IngestScope is the OAuth2 scope for the data ingestion
	DiodeOAuth2IngestScope = "diode:ingest"

	// DiodeSDKLogLevelEnvVarName is the environment variable name for the Diode SDK log level
	DiodeSDKLogLevelEnvVarName = "DIODE_SDK_LOG_LEVEL"

	// DiodeSkipTLSVerifyEnvVarName is the environment variable name to skip TLS verification
	DiodeSkipTLSVerifyEnvVarName = "DIODE_SKIP_TLS_VERIFY"

	defaultStreamName = "latest"

	authInitialRetryDelay = 1 * time.Second
	authMaxRetryDelay     = 30 * time.Second
)

var (
	// ErrInvalidTargetScheme is returned when the target URL does not start with a valid scheme.
	ErrInvalidTargetScheme = errors.New("target should start with grpc://, grpcs://, http:// or https://")

	allowedSchemesRe = regexp.MustCompile(`grpc|grpcs|http|https`)
)

// loadCerts loads the x509 cert pool from custom cert file or system certs
func loadCerts(certFile string) (*x509.CertPool, error) {
	if certFile != "" {
		certData, err := os.ReadFile(certFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read certificate file: %w", err)
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(certData) {
			return nil, fmt.Errorf("failed to parse certificate file")
		}
		return certPool, nil
	}
	return x509.SystemCertPool()
}

// skipTLSVerify determines if TLS certificate verification should be skipped for secure schemes
// This function should only be called for secure schemes (grpcs://, https://)
func skipTLSVerify() bool {
	// Check environment variable to skip TLS verification
	skipTLSEnv := strings.ToLower(os.Getenv(DiodeSkipTLSVerifyEnvVarName))
	skipTLSFromEnv := skipTLSEnv == "true" || skipTLSEnv == "1" || skipTLSEnv == "yes" || skipTLSEnv == "on"

	// TLS verification is enabled by default for secure schemes, disabled only by env var
	return skipTLSFromEnv
}

// parseTarget parses the target string into authority, path, isPlaintext, and tlsVerify
func parseTarget(target string) (string, string, bool, bool, error) {
	u, err := url.Parse(target)
	if err != nil {
		return "", "", false, false, err
	}

	if !allowedSchemesRe.MatchString(u.Scheme) {
		return "", "", false, false, ErrInvalidTargetScheme
	}

	authority := u.Host
	if u.Port() == "" {
		switch u.Scheme {
		case "grpc", "http":
			authority += ":80"
		case "grpcs", "https":
			authority += ":443"
		default:
			return "", "", false, false, fmt.Errorf("missing port with unsupported scheme: %s: %w", u.Scheme, ErrInvalidTargetScheme)
		}
	}

	path := u.Path
	if path == "/" {
		path = ""
	}

	// Determine if this is a plaintext connection
	isPlaintext := u.Scheme == "grpc" || u.Scheme == "http"

	// Default to TLS verification for secure schemes, disable for plaintext or when explicitly skipped
	tlsVerify := !isPlaintext && !skipTLSVerify()

	return authority, path, isPlaintext, tlsVerify, nil
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

// getCertFile returns the cert file path either from provided value or environment variable
func getCertFile(certFile string) string {
	if certFile == "" {
		certFile = os.Getenv(DiodeCertFileEnvVarName)
	}
	return certFile
}

// getSDKVersion returns the SDK version from build info or a fallback
func getSDKVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	// Check if this module is in the dependency list (when used as a dependency)
	for _, mod := range info.Deps {
		if strings.HasSuffix(mod.Path, "/diode-sdk-go") || mod.Path == "github.com/netboxlabs/diode-sdk-go" {
			return mod.Version
		}
	}

	return "dev"
}

// Client is an interface that defines the methods available from Diode API
type Client interface {
	// Close closes the connection to the API service
	Close() error

	// Ingest sends an ingest request to the ingester service
	Ingest(context.Context, []Entity, ...IngestOption) (*diodepb.IngestResponse, error)

	// IngestProto sends an ingest request to the ingester service with proto entities
	IngestProto(context.Context, []*diodepb.Entity, ...IngestOption) (*diodepb.IngestResponse, error)
}

// IngestOption is a functional option for ingest operations
type IngestOption func(*ingestConfig)

// ingestConfig holds configuration for ingest operations
type ingestConfig struct {
	metadata         Metadata
	maxChunkSizeMB   float64
	enableChunking   bool
	returnAllResults bool
}

// WithIngestMetadata adds optional metadata to the IngestRequest
func WithIngestMetadata(metadata Metadata) IngestOption {
	return func(c *ingestConfig) {
		c.metadata = metadata
	}
}

// WithChunking enables automatic message chunking for large entity lists.
// When enabled, entities will be automatically split into size-appropriate chunks
// before being sent to the server. Only the result from the last chunk is returned
// unless WithChunkingReturnAllResults is also used.
//
// The maxChunkSizeMB parameter sets the maximum size per chunk in megabytes.
// Use 0 to apply the default of 3.0 MB, which provides a safe margin below
// the gRPC 4 MB message size limit.
//
// Example:
//
//	// Use default 3.0 MB chunks
//	client.Ingest(ctx, entities, WithChunking(0))
//
//	// Use custom 3.5 MB chunks
//	client.Ingest(ctx, entities, WithChunking(3.5))
func WithChunking(maxChunkSizeMB float64) IngestOption {
	return func(c *ingestConfig) {
		c.enableChunking = true
		c.maxChunkSizeMB = maxChunkSizeMB
	}
}

// WithChunkingReturnAllResults modifies chunking behavior to return all chunk results.
// By default, only the last chunk's result is returned. Use this option when you need
// to collect results from all chunks.
//
// This option only has an effect when used with WithChunking.
//
// Example:
//
//	client.Ingest(ctx, entities, WithChunking(3.0), WithChunkingReturnAllResults())
func WithChunkingReturnAllResults() IngestOption {
	return func(c *ingestConfig) {
		c.returnAllResults = true
	}
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

	// Custom certificate file path
	certFile string

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

	// root CAs
	rootCAs *x509.CertPool

	// Plaintext connection (grpc://, http://)
	isPlaintext bool

	// TLS verify (only meaningful for secure connections)
	tlsVerify bool

	// Platform name
	platform string

	// Go version
	goVersion string

	// SDK name
	sdkName string

	// SDK version
	sdkVersion string

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

// WithCertFile sets the certificate file path for the GRPCClient
func WithCertFile(certFile string) ClientOption {
	return func(c *GRPCClient) {
		c.certFile = certFile
	}
}

// WithSkipTLSVerify disables TLS certificate verification
func WithSkipTLSVerify() ClientOption {
	return func(c *GRPCClient) {
		c.tlsVerify = false
	}
}

func formatClientUserAgent(sdkName, sdkVersion, appName, appVersion string) string {
	return fmt.Sprintf("%s/%s %s/%s", sdkName, sdkVersion, appName, appVersion)
}

// authenticate fetches an OAuth2 token using client credentials and updates the metadata with the token.
func (g *GRPCClient) authenticate(ctx context.Context) error {
	authClient := newDiodeAuthentication(
		g.target, g.path, g.isPlaintext, g.tlsVerify, g.rootCAs,
		g.clientID, g.clientSecret,
		g.sdkName, g.sdkVersion, g.appName, g.appVersion,
	)
	accessToken, err := authClient.authenticate(ctx, g.logger, []string{DiodeOAuth2IngestScope}, g.maxAuthRetries)
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
	rootCAs      *x509.CertPool
	isPlaintext  bool
	tlsVerify    bool
	clientID     string
	clientSecret string
	sdkName      string
	sdkVersion   string
	appName      string
	appVersion   string

	// Test hooks; zero values use production defaults in authenticate().
	initialRetryDelay time.Duration
	maxRetryDelay     time.Duration
}

// NewDiodeAuthentication creates a new instance of DiodeAuthentication.
func newDiodeAuthentication(
	target string,
	path string,
	isPlaintext bool,
	tlsVerify bool,
	rootCAs *x509.CertPool,
	clientID, clientSecret string,
	sdkName, sdkVersion, appName, appVersion string,
) *diodeAuthentication {
	return &diodeAuthentication{
		target:       target,
		path:         path,
		isPlaintext:  isPlaintext,
		tlsVerify:    tlsVerify,
		rootCAs:      rootCAs,
		clientID:     clientID,
		clientSecret: clientSecret,
		sdkName:      sdkName,
		sdkVersion:   sdkVersion,
		appName:      appName,
		appVersion:   appVersion,
	}
}

func isRetriableAuthHTTPStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable:
		return true
	default:
		return false
	}
}

const maxRetryAfterSeconds = int64((1<<63 - 1) / int64(time.Second))

func parseRetryAfterHeader(value string) (time.Duration, bool) {
	if value == "" {
		return 0, false
	}
	if seconds, err := strconv.Atoi(value); err == nil {
		if seconds < 0 {
			return 0, false
		}
		sec := int64(seconds)
		if sec > maxRetryAfterSeconds {
			sec = maxRetryAfterSeconds
		}
		return time.Duration(sec) * time.Second, true
	}
	if t, err := http.ParseTime(value); err == nil {
		delay := time.Until(t)
		if delay < 0 {
			delay = 0
		}
		return delay, true
	}
	return 0, false
}

func authRetryDelay(attempt int, statusCode int, retryAfter string, initialDelay, maxDelay time.Duration) time.Duration {
	var delay time.Duration
	if statusCode == http.StatusTooManyRequests || statusCode == http.StatusServiceUnavailable {
		if parsed, ok := parseRetryAfterHeader(retryAfter); ok {
			delay = parsed
		}
	}
	if delay == 0 {
		delay = initialDelay
		for i := 1; i < attempt; i++ {
			delay *= 2
			if delay >= maxDelay {
				delay = maxDelay
				break
			}
		}
	}
	if delay > maxDelay {
		delay = maxDelay
	}
	jitter := time.Duration(rand.Int63n(int64(delay)/4 + 1))
	total := delay + jitter
	if total > maxDelay {
		total = maxDelay
	}
	return total
}

func waitForAuthRetry(ctx context.Context, delay time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if delay <= 0 {
		return nil
	}

	timer := time.NewTimer(delay)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// Authenticate requests an OAuth2 token using client credentials and returns it.
func (d *diodeAuthentication) authenticate(ctx context.Context, logger *slog.Logger, scopes []string, maxRetries int) (string, error) {
	scheme := "https"
	if d.isPlaintext {
		scheme = "http"
	}
	authURL := fmt.Sprintf("%s://%s/auth/token", scheme, d.target)
	if d.path != "" {
		authURL = fmt.Sprintf("%s://%s%s/auth/token", scheme, d.target, d.path)
	}
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", d.clientID)
	formData.Set("client_secret", d.clientSecret)
	formData.Set("scope", strings.Join(scopes, " "))

	initialDelay := d.initialRetryDelay
	if initialDelay == 0 {
		initialDelay = authInitialRetryDelay
	}
	maxDelay := d.maxRetryDelay
	if maxDelay == 0 {
		maxDelay = authMaxRetryDelay
	}
	client := &http.Client{}
	if d.isPlaintext {
		client.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		}
	} else {
		client.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs:            d.rootCAs,
				InsecureSkipVerify: !d.tlsVerify,
			},
		}
	}

	var lastStatus string
	for attempt := 1; attempt <= maxRetries; attempt++ {
		if err := ctx.Err(); err != nil {
			return "", fmt.Errorf("authentication canceled: %w", err)
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, authURL, strings.NewReader(formData.Encode()))
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", formatClientUserAgent(d.sdkName, d.sdkVersion, d.appName, d.appVersion))

		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("failed to send request: %w", err)
		}

		if resp.StatusCode == http.StatusOK {
			var result struct {
				AccessToken string `json:"access_token"`
			}
			decodeErr := json.NewDecoder(resp.Body).Decode(&result)
			closeErr := resp.Body.Close()
			if decodeErr != nil {
				return "", fmt.Errorf("failed to parse response: %w", decodeErr)
			}
			if closeErr != nil {
				logger.Error("failed to close response body", "error", closeErr)
			}
			if result.AccessToken == "" {
				return "", errors.New("access token not found in response")
			}
			return result.AccessToken, nil
		}

		lastStatus = resp.Status
		retryAfter := resp.Header.Get("Retry-After")
		_, _ = io.Copy(io.Discard, resp.Body)
		if closeErr := resp.Body.Close(); closeErr != nil {
			logger.Error("failed to close response body", "error", closeErr)
		}

		if !isRetriableAuthHTTPStatus(resp.StatusCode) || attempt >= maxRetries {
			return "", fmt.Errorf("authentication failed: %s", lastStatus)
		}

		delay := authRetryDelay(attempt, resp.StatusCode, retryAfter, initialDelay, maxDelay)
		logger.Debug(
			"Auth token request failed, retrying",
			"status", resp.StatusCode,
			"attempt", attempt,
			"retry_in", delay,
		)
		if err := waitForAuthRetry(ctx, delay); err != nil {
			return "", fmt.Errorf("authentication canceled: %w", err)
		}
	}

	return "", fmt.Errorf("authentication failed: %s", lastStatus)
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

	target, path, isPlaintext, tlsVerify, err := parseTarget(target)
	if err != nil {
		return nil, err
	}

	platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	goVersion := runtime.Version()
	sdkVersion := getSDKVersion()

	c := &GRPCClient{
		logger:         logger,
		appName:        appName,
		appVersion:     appVersion,
		target:         target,
		path:           path,
		isPlaintext:    isPlaintext,
		tlsVerify:      tlsVerify,
		platform:       platform,
		goVersion:      goVersion,
		sdkName:        SDKName,
		sdkVersion:     sdkVersion,
		maxAuthRetries: 3,
	}

	var clientID string
	var clientSecret string

	for _, o := range opts {
		o(c)
	}

	certFile := getCertFile(c.certFile)
	c.certFile = certFile

	// Load certificates
	rootCAs, err := loadCerts(c.certFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificates: %w", err)
	}
	c.rootCAs = rootCAs

	userAgent := formatClientUserAgent(c.sdkName, c.sdkVersion, c.appName, c.appVersion)
	dialOpts := []grpc.DialOption{
		grpc.WithUserAgent(userAgent),
		defaultClientKeepaliveDialOption(),
	}

	if path != "" {
		logger.Debug("Setting up gRPC interceptor for path", "path", path)
		dialOpts = append(dialOpts, methodUnaryInterceptor(path))
	}

	// Setup transport credentials based on connection type
	if c.isPlaintext {
		// Use plaintext for grpc:// and http://
		logger.Debug("Setting up gRPC plaintext channel")
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// Always use TLS for secure schemes (grpcs://, https://)
		logger.Debug("Setting up gRPC secure channel")
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			RootCAs:            rootCAs,
			InsecureSkipVerify: !c.tlsVerify, // Skip verification if tlsVerify is false
		})))
	}

	conn, err := grpc.NewClient(target, dialOpts...)
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.client = diodepb.NewIngesterServiceClient(conn)

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

	if err = c.authenticate(context.Background()); err != nil {
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
func (g *GRPCClient) Ingest(ctx context.Context, entities []Entity, opts ...IngestOption) (*diodepb.IngestResponse, error) {
	return g.IngestProto(ctx, convertEntitiesToProto(entities), opts...)
}

// IngestProto sends an ingest request to the ingester service with proto entities
func (g *GRPCClient) IngestProto(ctx context.Context, entities []*diodepb.Entity, opts ...IngestOption) (*diodepb.IngestResponse, error) {
	// Apply options
	cfg := &ingestConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	// Handle chunking if enabled
	if cfg.enableChunking {
		return g.ingestWithChunking(ctx, entities, cfg)
	}

	// Standard single-request ingestion
	return g.ingestSingleRequest(ctx, entities, cfg)
}

// ingestSingleRequest sends a single ingest request without chunking
func (g *GRPCClient) ingestSingleRequest(ctx context.Context, entities []*diodepb.Entity, cfg *ingestConfig) (*diodepb.IngestResponse, error) {
	stream := defaultStreamName

	req := &diodepb.IngestRequest{
		Id:                 uuid.NewString(),
		Entities:           entities,
		Stream:             stream,
		ProducerAppName:    g.appName,
		ProducerAppVersion: g.appVersion,
		SdkName:            g.sdkName,
		SdkVersion:         g.sdkVersion,
	}

	// Add metadata to request if provided
	if len(cfg.metadata) > 0 {
		req.Metadata, _ = structpb.NewStruct(cfg.metadata)
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
				if err := g.authenticate(ctx); err != nil {
					if ctx.Err() != nil {
						return nil, fmt.Errorf("re-authentication canceled: %w", ctx.Err())
					}
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

// ingestWithChunking sends entities in chunks
func (g *GRPCClient) ingestWithChunking(ctx context.Context, entities []*diodepb.Entity, cfg *ingestConfig) (*diodepb.IngestResponse, error) {
	chunks := CreateMessageChunks(entities, cfg.maxChunkSizeMB)

	g.logger.Debug("Chunking enabled", "total_entities", len(entities), "chunks", len(chunks), "max_chunk_size_mb", cfg.maxChunkSizeMB)

	var lastResponse *diodepb.IngestResponse
	var allResponses []*diodepb.IngestResponse

	for i, chunk := range chunks {
		g.logger.Debug("Ingesting chunk", "chunk", i+1, "of", len(chunks), "entities", len(chunk))

		res, err := g.ingestSingleRequest(ctx, chunk, cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to ingest chunk %d of %d: %w", i+1, len(chunks), err)
		}

		lastResponse = res
		if cfg.returnAllResults {
			allResponses = append(allResponses, res)
		}
	}

	// If returnAllResults is enabled, we need to combine the responses
	// For now, we'll return the last response as the primary one
	// In the future, this could be enhanced to merge error counts, etc.
	if cfg.returnAllResults && len(allResponses) > 0 {
		// TODO: Consider merging response statistics from all chunks
		// For now, just return the last response
		g.logger.Debug("Completed chunked ingestion", "total_chunks", len(allResponses))
	}

	return lastResponse, nil
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
