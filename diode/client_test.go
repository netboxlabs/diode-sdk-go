package diode

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

func TestLoadCerts(t *testing.T) {
	tests := []struct {
		desc     string
		certFile string
		wantErr  bool
	}{
		{
			desc:     "load system certificates",
			certFile: "",
			wantErr:  false,
		},
		{
			desc:     "load custom certificate file",
			certFile: "testdata/test-cert.pem",
			wantErr:  false,
		},
		{
			desc:     "load certificate chain file",
			certFile: "testdata/cert-chain.pem",
			wantErr:  false,
		},
		{
			desc:     "load non-existent certificate file",
			certFile: "testdata/non-existent.pem",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			certPool, err := loadCerts(tt.certFile)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, certPool)
			} else {
				require.NoError(t, err)
				require.NotNil(t, certPool)
			}
		})
	}
}

func TestParseTarget(t *testing.T) {
	tests := []struct {
		desc      string
		target    string
		authority string
		path      string
		tlsVerify bool
		wantErr   error
	}{
		{
			desc:      "valid target without path and tls verification",
			target:    "grpc://localhost:8081",
			authority: "localhost:8081",
			path:      "",
			tlsVerify: false,
			wantErr:   nil,
		},
		{
			desc:      "valid target with path",
			target:    "grpc://localhost:8081/fsfsd",
			authority: "localhost:8081",
			path:      "/fsfsd",
			tlsVerify: false,
			wantErr:   nil,
		},
		{
			desc:      "valid target with tls",
			target:    "grpcs://localhost:8081",
			authority: "localhost:8081",
			path:      "",
			tlsVerify: true,
			wantErr:   nil,
		},
		{
			desc:      "valid HTTP target",
			target:    "http://localhost:8081",
			authority: "localhost:8081",
			tlsVerify: false,
			wantErr:   nil,
		},
		{
			desc:      "valid HTTP target with tls",
			target:    "https://localhost:8081",
			authority: "localhost:8081",
			tlsVerify: true,
			wantErr:   nil,
		},
		{
			desc:      "valid target empty path on grpc://localhost:8081/",
			target:    "grpc://localhost:8081/",
			authority: "localhost:8081",
			path:      "",
			tlsVerify: false,
			wantErr:   nil,
		},
		{
			desc:      "valid target without port having 443 appended",
			target:    "grpcs://localhost",
			authority: "localhost:443",
			path:      "",
			tlsVerify: true,
			wantErr:   nil,
		},
		{
			desc:      "invalid scheme in target",
			target:    "ftp://localhost:8081",
			authority: "",
			path:      "",
			tlsVerify: false,
			wantErr:   ErrInvalidTargetScheme,
		},
		{
			desc:      "invalid target",
			target:    "grpc://local%host:8081",
			authority: "",
			path:      "",
			tlsVerify: false,
			wantErr:   &url.Error{Op: "parse", URL: "grpc://local%host:8081", Err: url.EscapeError("%ho")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			authority, path, _, tlsVerify, err := parseTarget(tt.target)
			assert.Equal(t, tt.authority, authority)
			assert.Equal(t, tt.path, path)
			assert.Equal(t, tt.tlsVerify, tlsVerify)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetCertFile(t *testing.T) {
	tests := []struct {
		desc                string
		certFile            string
		certFileEnvVarValue string
		wantCertFile        string
	}{
		{
			desc:                "cert file provided explicitly",
			certFile:            "path/to/cert.pem",
			certFileEnvVarValue: "",
			wantCertFile:        "path/to/cert.pem",
		},
		{
			desc:                "cert file provided via environment variable",
			certFile:            "",
			certFileEnvVarValue: "env/cert.pem",
			wantCertFile:        "env/cert.pem",
		},
		{
			desc:                "no cert file provided",
			certFile:            "",
			certFileEnvVarValue: "",
			wantCertFile:        "",
		},
		{
			desc:                "explicit cert file takes precedence over env var",
			certFile:            "explicit/cert.pem",
			certFileEnvVarValue: "env/cert.pem",
			wantCertFile:        "explicit/cert.pem",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.certFileEnvVarValue != "" {
				_ = os.Setenv(DiodeCertFileEnvVarName, tt.certFileEnvVarValue)
				defer func() {
					_ = os.Unsetenv(DiodeCertFileEnvVarName)
				}()
			}

			certFile := getCertFile(tt.certFile)
			require.Equal(t, tt.wantCertFile, certFile)
		})
	}
}

func TestGetClientCredentials(t *testing.T) {
	tests := []struct {
		desc                    string
		clientID                string
		clientSecret            string
		clientIDEnvVarValue     string
		clientSecretEnvVarValue string
		wantClientID            string
		wantClientSecret        string
		wantIDErr               error
		wantSecretErr           error
	}{
		{
			desc:                    "Client credentials provided explicitly",
			clientID:                "client-id-123",
			clientSecret:            "client-secret-456",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantClientID:            "client-id-123",
			wantClientSecret:        "client-secret-456",
			wantIDErr:               nil,
			wantSecretErr:           nil,
		},
		{
			desc:                    "Client credentials provided via environment variables",
			clientID:                "",
			clientSecret:            "",
			clientIDEnvVarValue:     "env-client-id",
			clientSecretEnvVarValue: "env-client-secret",
			wantClientID:            "env-client-id",
			wantClientSecret:        "env-client-secret",
			wantIDErr:               nil,
			wantSecretErr:           nil,
		},
		{
			desc:                    "Missing clientID and clientSecret",
			clientID:                "",
			clientSecret:            "",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantClientID:            "",
			wantClientSecret:        "",
			wantIDErr:               errors.New("client_id param or DIODE_CLIENT_ID environment variable required"),
			wantSecretErr:           errors.New("client_secret param or DIODE_CLIENT_SECRET environment variable required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.clientIDEnvVarValue != "" {
				_ = os.Setenv(DiodeClientIDEnvVarName, tt.clientIDEnvVarValue)
				defer func() {
					_ = os.Unsetenv(DiodeClientIDEnvVarName)
				}()
			}
			if tt.clientSecretEnvVarValue != "" {
				_ = os.Setenv(DiodeClientSecretEnvVarName, tt.clientSecretEnvVarValue)
				defer func() {
					_ = os.Unsetenv(DiodeClientSecretEnvVarName)
				}()
			}

			clientID, err := getClientID(tt.clientID)
			require.Equal(t, tt.wantClientID, clientID)
			require.Equal(t, tt.wantIDErr, err)

			clientSecret, err := getClientSecret(tt.clientSecret)
			require.Equal(t, tt.wantClientSecret, clientSecret)
			require.Equal(t, tt.wantSecretErr, err)
		})
	}
}

type MockAuthServer struct {
	listener   net.Listener
	httpServer *http.Server
	grpcServer *grpc.Server
}

func (m *MockAuthServer) Close() {
	m.grpcServer.GracefulStop()
	_ = m.httpServer.Shutdown(context.Background())
	_ = m.listener.Close()
}

func startMockAuthServer(port string, path string, authErrorGrpc bool) (*MockAuthServer, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	diodepb.RegisterIngesterServiceServer(grpcServer, &MockIngesterServiceServer{unauthenticatedError: authErrorGrpc})

	// HTTP handler
	httpMux := http.NewServeMux()

	newPath := "/auth/token"
	if path != "" {
		newPath = fmt.Sprintf("/%s/auth/token", path)
	}
	httpMux.HandleFunc(newPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "invalid method", http.StatusMethodNotAllowed)
			return
		}
		_ = r.ParseForm()
		clientID := r.FormValue("client_id")
		clientSecret := r.FormValue("client_secret")
		scope := r.FormValue("scope")
		if strings.Contains(clientID, "client-id") && strings.Contains(clientSecret, "client-secret") && strings.Contains(scope, DiodeOAuth2IngestScope) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token": "mock-token"}`))
		} else {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		}
	})

	// Wrap the mux with h2c-compatible handler that detects gRPC
	handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpMux.ServeHTTP(w, r)
		}
	}), &http2.Server{})

	httpServer := &http.Server{
		Handler: handler,
	}

	go func() {
		if err := httpServer.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	return &MockAuthServer{
		listener:   listener,
		httpServer: httpServer,
		grpcServer: grpcServer,
	}, nil
}

func TestNewClient(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err)

	authServer, err := startMockAuthServer(port, "", false)
	require.NoError(t, err)
	defer authServer.Close()

	tests := []struct {
		desc                    string
		target                  string
		appName                 string
		appVersion              string
		clientID                string
		clientSecret            string
		clientIDEnvVarValue     string
		clientSecretEnvVarValue string
		wantErr                 error
	}{
		{
			desc:                    "explicit arguments provided",
			target:                  fmt.Sprintf("grpc://localhost:%s", port),
			appName:                 "my-producer",
			appVersion:              "0.1.0",
			clientID:                "client-id-123",
			clientSecret:            "client-secret-456",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantErr:                 nil,
		},
		{
			desc:                    "Client credentials provided via environment variables",
			target:                  fmt.Sprintf("grpc://localhost:%s", port),
			appName:                 "my-producer",
			appVersion:              "0.1.0",
			clientID:                "",
			clientSecret:            "",
			clientIDEnvVarValue:     "env-client-id",
			clientSecretEnvVarValue: "env-client-secret",
			wantErr:                 nil,
		},
		{
			desc:                    "app name not provided",
			target:                  "grpc://localhost:8081",
			appName:                 "",
			appVersion:              "0.1.0",
			clientID:                "client-id-123",
			clientSecret:            "client-secret-456",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantErr:                 errors.New("app name is required"),
		},
		{
			desc:                    "app version not provided",
			target:                  "grpc://localhost:8081",
			appName:                 "my-producer",
			appVersion:              "",
			clientID:                "client-id-123",
			clientSecret:            "client-secret-456",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantErr:                 errors.New("app version is required"),
		},
		{
			desc:                    "invalid target",
			target:                  "ftp://localhost:8081",
			appName:                 "my-producer",
			appVersion:              "0.1.0",
			clientID:                "client-id-123",
			clientSecret:            "client-secret-456",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantErr:                 ErrInvalidTargetScheme,
		},
		{
			desc:                    "missing clientID and clientSecret",
			target:                  "grpc://localhost:8081",
			appName:                 "my-producer",
			appVersion:              "0.1.0",
			clientID:                "",
			clientSecret:            "",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantErr:                 errors.New("client_id param or DIODE_CLIENT_ID environment variable required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
				_ = os.Unsetenv(DiodeSDKLogLevelEnvVarName)
			}()

			if tt.clientIDEnvVarValue != "" {
				_ = os.Setenv(DiodeClientIDEnvVarName, tt.clientIDEnvVarValue)
			}
			if tt.clientSecretEnvVarValue != "" {
				_ = os.Setenv(DiodeClientSecretEnvVarName, tt.clientSecretEnvVarValue)
			}

			opts := []ClientOption{}
			if tt.clientID != "" {
				opts = append(opts, WithClientID(tt.clientID))
			}
			if tt.clientSecret != "" {
				opts = append(opts, WithClientSecret(tt.clientSecret))
			}

			client, err := NewClient(tt.target, tt.appName, tt.appVersion, opts...)
			require.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				require.NotNil(t, client)
				require.NoError(t, client.Close())
			}
		})
	}
}

func TestNewClientWithCertFile(t *testing.T) {
	tests := []struct {
		desc                string
		target              string
		certFile            string
		certFileEnvVarValue string
		wantTLSVerify       bool
		wantErr             error
		expectAuthError     bool
	}{
		{
			desc:            "cert file provided via option with insecure target",
			target:          "grpc://localhost:8080",
			certFile:        "testdata/test-cert.pem",
			wantTLSVerify:   false, // grpc:// scheme = insecure
			wantErr:         nil,
			expectAuthError: true, // No server running, will fail auth
		},
		{
			desc:                "cert file provided via environment variable with insecure target",
			target:              "grpc://localhost:8080",
			certFile:            "",
			certFileEnvVarValue: "testdata/test-cert.pem",
			wantTLSVerify:       false, // grpc:// scheme = insecure
			wantErr:             nil,
			expectAuthError:     true, // No server running, will fail auth
		},
		{
			desc:                "explicit cert file takes precedence over env var with insecure target",
			target:              "grpc://localhost:8080",
			certFile:            "testdata/test-cert.pem",
			certFileEnvVarValue: "testdata/invalid.pem",
			wantTLSVerify:       false, // grpc:// scheme = insecure
			wantErr:             nil,
			expectAuthError:     true, // No server running, will fail auth
		},
		{
			desc:            "invalid cert file",
			target:          "grpc://localhost:8080",
			certFile:        "testdata/non-existent.pem",
			wantTLSVerify:   false,
			wantErr:         errors.New("failed to load certificates"),
			expectAuthError: false, // Will fail during cert loading, before auth
		},
		{
			desc:            "grpcs target with custom cert file",
			target:          "grpcs://localhost:8080",
			certFile:        "testdata/test-cert.pem",
			wantTLSVerify:   true,
			wantErr:         nil,
			expectAuthError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
				_ = os.Unsetenv(DiodeCertFileEnvVarName)
			}()

			_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
			_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")
			if tt.certFileEnvVarValue != "" {
				_ = os.Setenv(DiodeCertFileEnvVarName, tt.certFileEnvVarValue)
			}

			opts := []ClientOption{}
			if tt.certFile != "" {
				opts = append(opts, WithCertFile(tt.certFile))
			}

			client, err := NewClient(tt.target, "my-producer", "0.1.0", opts...)
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else if tt.expectAuthError {
				// We expect the client creation to fail during authentication
				require.Error(t, err)
				assert.Contains(t, err.Error(), "authentication failed")
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)

				grpcClient := client.(*GRPCClient)
				assert.Equal(t, tt.wantTLSVerify, grpcClient.tlsVerify)
				if tt.certFile != "" || tt.certFileEnvVarValue != "" {
					expectedCertFile := tt.certFile
					if expectedCertFile == "" {
						expectedCertFile = tt.certFileEnvVarValue
					}
					assert.Equal(t, expectedCertFile, grpcClient.certFile)
				}

				require.NoError(t, client.Close())
			}
		})
	}
}

func getFreePort() (string, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return strconv.Itoa(0), err
	}

	addr := listener.Addr().(*net.TCPAddr)

	if err = listener.Close(); err != nil {
		return strconv.Itoa(0), err
	}
	return strconv.Itoa(addr.Port), nil
}

type MockIngesterServiceServer struct {
	diodepb.UnimplementedIngesterServiceServer
	unauthenticatedError bool
}

func (m MockIngesterServiceServer) Ingest(_ context.Context, _ *diodepb.IngestRequest) (*diodepb.IngestResponse, error) {
	if m.unauthenticatedError {
		return nil, status.Error(codes.Unauthenticated, "mock auth failure")
	}
	return &diodepb.IngestResponse{Errors: nil}, nil
}

func TestMethodUnaryInterceptor(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		wantErr error
	}{
		{
			desc:    "empty path",
			path:    "",
			wantErr: nil,
		},
		{
			desc:    "non-empty path",
			path:    "foobar",
			wantErr: errors.New("rpc error: code = Unimplemented desc = unknown service foobar/diode.v1.IngesterService"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			port, err := getFreePort()
			require.NoError(t, err)

			authServer, err := startMockAuthServer(port, tt.path, false)
			require.NoError(t, err)
			defer authServer.Close()

			target := fmt.Sprintf("grpc://localhost:%s/%s", port, tt.path)
			client, err := NewClient(target, "my-producer", "0.1.0", WithClientID("client-id"), WithClientSecret("client-secret"))
			require.NoError(t, err)
			require.NotNil(t, client)

			_, err = client.Ingest(context.Background(), nil)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
			require.NoError(t, client.Close())
		})
	}
}

func TestGetSDKVersion(t *testing.T) {
	tests := []struct {
		desc        string
		wantVersion string
	}{
		{
			desc:        "SDK version detection",
			wantVersion: "dev", // When running tests, build info won't have module dependencies
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			version := getSDKVersion()
			// Since we're running in a test environment without proper module deps,
			// we should at least get a non-empty version string
			assert.NotEmpty(t, version)
			// When running tests locally, it should return "dev"
			if version != "dev" {
				// If not "dev", it should be a valid version string (starts with v or is a commit hash)
				assert.True(t, version[0] == 'v' || len(version) >= 7,
					"version should be 'dev', start with 'v', or be a commit hash (7+ chars): %s", version)
			}
		})
	}
}

func TestClientSDKVersionCaching(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err)

	authServer, err := startMockAuthServer(port, "", false)
	require.NoError(t, err)
	defer authServer.Close()

	_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
	_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")
	defer func() {
		_ = os.Unsetenv(DiodeClientIDEnvVarName)
		_ = os.Unsetenv(DiodeClientSecretEnvVarName)
	}()

	client, err := NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
	require.NoError(t, err)
	defer func() {
		err := client.Close()
		require.NoError(t, err)
	}()

	grpcClient := client.(*GRPCClient)

	// Verify that SDK name and version are cached
	assert.Equal(t, SDKName, grpcClient.sdkName)
	assert.NotEmpty(t, grpcClient.sdkVersion)

	// The cached version should be consistent
	cachedVersion := grpcClient.sdkVersion

	// Create another client and verify it gets the same version
	client2, err := NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
	require.NoError(t, err)
	defer func() {
		err := client2.Close()
		require.NoError(t, err)
	}()

	grpcClient2 := client2.(*GRPCClient)
	assert.Equal(t, cachedVersion, grpcClient2.sdkVersion)
	assert.Equal(t, SDKName, grpcClient2.sdkName)
}

func TestNewLogger(t *testing.T) {
	tests := []struct {
		desc                string
		logLevelEnvVarValue string
		wantLogLevel        slog.Level
	}{
		{
			desc:                "log level not provided",
			logLevelEnvVarValue: "",
			wantLogLevel:        slog.LevelInfo,
		},
		{
			desc:                "debug log level provided",
			logLevelEnvVarValue: "debug",
			wantLogLevel:        slog.LevelDebug,
		},
		{
			desc:                "info log level provided",
			logLevelEnvVarValue: "info",
			wantLogLevel:        slog.LevelInfo,
		},
		{
			desc:                "warn log level provided",
			logLevelEnvVarValue: "warn",
			wantLogLevel:        slog.LevelWarn,
		},
		{
			desc:                "error log level provided",
			logLevelEnvVarValue: "error",
			wantLogLevel:        slog.LevelError,
		},
		{
			desc:                "invalid log level provided",
			logLevelEnvVarValue: "invalid",
			wantLogLevel:        slog.LevelDebug,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeSDKLogLevelEnvVarName)
			}()

			if tt.logLevelEnvVarValue != "" {
				_ = os.Setenv(DiodeSDKLogLevelEnvVarName, tt.logLevelEnvVarValue)
			}

			logger := newLogger()
			require.NotNil(t, logger)
			assert.True(t, logger.Enabled(context.Background(), tt.wantLogLevel))
		})
	}
}

func TestConvertEntitiesToProto(t *testing.T) {
	tests := []struct {
		desc     string
		entities []Entity
	}{
		{
			desc:     "empty entities",
			entities: nil,
		},
		{
			desc: "non-empty entities",
			entities: []Entity{
				&Device{Name: String("device-1")},
				&Site{Name: String("site-1")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			protoEntities := convertEntitiesToProto(tt.entities)
			require.NotNil(t, protoEntities)
			assert.Equal(t, len(tt.entities), len(protoEntities))
			for _, entityPb := range protoEntities {
				assert.NotNil(t, entityPb.Timestamp)
				assert.NotZero(t, entityPb.Timestamp.Seconds)
				assert.NotZero(t, entityPb.Timestamp.Nanos)
			}
		})
	}
}

func TestIngest(t *testing.T) {
	tests := []struct {
		desc            string
		entities        []Entity
		authRetries     int
		mockAuthFailure bool
		wantErr         error
	}{
		{
			desc: "successful ingest",
			entities: []Entity{
				&Device{Name: String("device-1")},
				&Site{Name: String("site-1")},
			},
			authRetries:     3,
			mockAuthFailure: false,
			wantErr:         nil,
		},
		{
			desc:            "authentication failure after retries",
			entities:        nil,
			authRetries:     2,
			mockAuthFailure: true,
			wantErr:         errors.New("authentication failed after 2 attempts: rpc error: code = Unauthenticated desc = mock auth failure"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
			}()

			_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
			_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")

			port, err := getFreePort()
			require.NoError(t, err)

			authServer, err := startMockAuthServer(port, "", tt.mockAuthFailure)
			require.NoError(t, err)
			defer authServer.Close()

			client, err := NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
			require.NoError(t, err)
			defer func() {
				err := client.Close()
				require.NoError(t, err)
			}()

			grpcClient := client.(*GRPCClient)
			grpcClient.maxAuthRetries = tt.authRetries

			_, err = grpcClient.Ingest(context.Background(), tt.entities)
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIngestProto(t *testing.T) {
	defer func() {
		_ = os.Unsetenv(DiodeClientIDEnvVarName)
		_ = os.Unsetenv(DiodeClientSecretEnvVarName)
	}()

	_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
	_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")

	port, err := getFreePort()
	require.NoError(t, err)

	authServer, err := startMockAuthServer(port, "", false)
	require.NoError(t, err)
	defer authServer.Close()

	client, err := NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
	require.NoError(t, err)
	defer func() {
		err := client.Close()
		require.NoError(t, err)
	}()

	grpcClient := client.(*GRPCClient)

	entities := []*diodepb.Entity{
		{
			Entity: &diodepb.Entity_Device{
				Device: &diodepb.Device{
					Name:        String("device-1"),
					Description: String("Test device"),
				},
			},
		},
	}

	resp, err := grpcClient.IngestProto(context.Background(), entities)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	require.Empty(t, resp.Errors)
}

func TestHTTPAuthError(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err)

	authServer, err := startMockAuthServer(port, "", false)
	require.NoError(t, err)
	defer authServer.Close()

	_ = os.Setenv(DiodeClientIDEnvVarName, "invalid-client")
	_ = os.Setenv(DiodeClientSecretEnvVarName, "invalid-client-sct")
	_ = os.Setenv(DiodeMaxAuthRetriesEnvVarName, "2")
	defer func() {
		_ = os.Unsetenv(DiodeClientIDEnvVarName)
		_ = os.Unsetenv(DiodeClientSecretEnvVarName)
		_ = os.Unsetenv(DiodeMaxAuthRetriesEnvVarName)
	}()

	_, err = NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
	require.Error(t, err)
}

func TestWithSkipTLSVerify(t *testing.T) {
	tests := []struct {
		desc          string
		target        string
		withSkipTLS   bool
		wantTLSVerify bool
		expectError   bool
	}{
		{
			desc:          "grpcs target with WithSkipTLSVerify option",
			target:        "grpcs://localhost:8080",
			withSkipTLS:   true,
			wantTLSVerify: false, // WithSkipTLSVerify should override default
			expectError:   true,  // No server running, will fail auth
		},
		{
			desc:          "grpcs target without WithSkipTLSVerify option",
			target:        "grpcs://localhost:8080",
			withSkipTLS:   false,
			wantTLSVerify: true, // Default for secure schemes
			expectError:   true, // No server running, will fail auth
		},
		{
			desc:          "https target with WithSkipTLSVerify option",
			target:        "https://localhost:8080",
			withSkipTLS:   true,
			wantTLSVerify: false, // WithSkipTLSVerify should override default
			expectError:   true,  // No server running, will fail auth
		},
		{
			desc:          "grpc target with WithSkipTLSVerify option (should have no effect)",
			target:        "grpc://localhost:8080",
			withSkipTLS:   true,
			wantTLSVerify: false, // Plaintext schemes always false
			expectError:   true,  // No server running, will fail auth
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
			}()

			_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
			_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")

			opts := []ClientOption{}
			if tt.withSkipTLS {
				opts = append(opts, WithSkipTLSVerify())
			}

			client, err := NewClient(tt.target, "my-producer", "0.1.0", opts...)
			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "authentication failed")
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)

				grpcClient := client.(*GRPCClient)
				assert.Equal(t, tt.wantTLSVerify, grpcClient.tlsVerify)

				require.NoError(t, client.Close())
			}
		})
	}
}

func TestSkipTLSVerifyEnv(t *testing.T) {
	tests := []struct {
		desc            string
		target          string
		skipTLSEnvValue string
		wantTLSVerify   bool
		expectError     bool
	}{
		{
			desc:            "grpcs target with DIODE_SKIP_TLS_VERIFY=true",
			target:          "grpcs://localhost:8080",
			skipTLSEnvValue: "true",
			wantTLSVerify:   false, // Should skip TLS verification
			expectError:     true,  // No server running, will fail auth
		},
		{
			desc:            "grpcs target with DIODE_SKIP_TLS_VERIFY=1",
			target:          "grpcs://localhost:8080",
			skipTLSEnvValue: "1",
			wantTLSVerify:   false, // Should skip TLS verification
			expectError:     true,  // No server running, will fail auth
		},
		{
			desc:            "grpcs target with DIODE_SKIP_TLS_VERIFY=yes",
			target:          "grpcs://localhost:8080",
			skipTLSEnvValue: "yes",
			wantTLSVerify:   false, // Should skip TLS verification
			expectError:     true,  // No server running, will fail auth
		},
		{
			desc:            "grpcs target with DIODE_SKIP_TLS_VERIFY=false",
			target:          "grpcs://localhost:8080",
			skipTLSEnvValue: "false",
			wantTLSVerify:   true, // Should verify TLS (false is not a skip value)
			expectError:     true, // No server running, will fail auth
		},
		{
			desc:            "grpcs target with no DIODE_SKIP_TLS_VERIFY env var",
			target:          "grpcs://localhost:8080",
			skipTLSEnvValue: "",
			wantTLSVerify:   true, // Default to verify TLS
			expectError:     true, // No server running, will fail auth
		},
		{
			desc:            "grpc target with DIODE_SKIP_TLS_VERIFY=true (should have no effect)",
			target:          "grpc://localhost:8080",
			skipTLSEnvValue: "true",
			wantTLSVerify:   false, // Plaintext schemes always false
			expectError:     true,  // No server running, will fail auth
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
				_ = os.Unsetenv(DiodeSkipTLSVerifyEnvVarName)
			}()

			_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
			_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")
			if tt.skipTLSEnvValue != "" {
				_ = os.Setenv(DiodeSkipTLSVerifyEnvVarName, tt.skipTLSEnvValue)
			}

			client, err := NewClient(tt.target, "my-producer", "0.1.0")
			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "authentication failed")
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)

				grpcClient := client.(*GRPCClient)
				assert.Equal(t, tt.wantTLSVerify, grpcClient.tlsVerify)

				require.NoError(t, client.Close())
			}
		})
	}
}

func TestIsPlaintextField(t *testing.T) {
	tests := []struct {
		desc            string
		target          string
		wantIsPlaintext bool
		wantTLSVerify   bool
		expectError     bool
	}{
		{
			desc:            "grpc scheme should be plaintext",
			target:          "grpc://localhost:8080",
			wantIsPlaintext: true,
			wantTLSVerify:   false,
			expectError:     true, // No server running, will fail auth
		},
		{
			desc:            "http scheme should be plaintext",
			target:          "http://localhost:8080",
			wantIsPlaintext: true,
			wantTLSVerify:   false,
			expectError:     true, // No server running, will fail auth
		},
		{
			desc:            "grpcs scheme should not be plaintext",
			target:          "grpcs://localhost:8080",
			wantIsPlaintext: false,
			wantTLSVerify:   true,
			expectError:     true, // No server running, will fail auth
		},
		{
			desc:            "https scheme should not be plaintext",
			target:          "https://localhost:8080",
			wantIsPlaintext: false,
			wantTLSVerify:   true,
			expectError:     true, // No server running, will fail auth
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
			}()

			_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
			_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")

			client, err := NewClient(tt.target, "my-producer", "0.1.0")
			if tt.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "authentication failed")
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)

				grpcClient := client.(*GRPCClient)
				assert.Equal(t, tt.wantIsPlaintext, grpcClient.isPlaintext)
				assert.Equal(t, tt.wantTLSVerify, grpcClient.tlsVerify)

				require.NoError(t, client.Close())
			}
		})
	}
}

func TestAuthRetryEnv(t *testing.T) {
	tests := []struct {
		desc       string
		retryValue string
		wantErr    error
	}{
		{
			desc:       "valid value",
			retryValue: "2",
			wantErr:    nil,
		},
		{
			desc:       "empty value",
			retryValue: "",
			wantErr:    nil,
		},
		{
			desc:       "invalid value",
			retryValue: "invalid",
			wantErr:    errors.New("invalid value for DIODE_MAX_AUTH_RETRIES: strconv.Atoi: parsing \"invalid\": invalid syntax"),
		},
		{
			desc:       "negative value",
			retryValue: "-1",
			wantErr:    errors.New("max_auth_retries param or DIODE_MAX_AUTH_RETRIES environment variable must be greater than 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			port, err := getFreePort()
			require.NoError(t, err)

			authServer, err := startMockAuthServer(port, "", false)
			require.NoError(t, err)
			defer authServer.Close()

			_ = os.Setenv(DiodeClientIDEnvVarName, "client-id")
			_ = os.Setenv(DiodeClientSecretEnvVarName, "client-secret")
			_ = os.Setenv(DiodeMaxAuthRetriesEnvVarName, tt.retryValue)
			defer func() {
				_ = os.Unsetenv(DiodeClientIDEnvVarName)
				_ = os.Unsetenv(DiodeClientSecretEnvVarName)
				_ = os.Unsetenv(DiodeMaxAuthRetriesEnvVarName)
			}()

			_, err = NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
