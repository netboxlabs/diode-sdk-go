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
	certPool := loadCerts()
	assert.NotNil(t, certPool)
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
			target:    "http://localhost:8081",
			authority: "",
			path:      "",
			tlsVerify: false,
			wantErr:   errors.New("target should start with grpc:// or grpcs://"),
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
			authority, path, tlsVerify, err := parseTarget(tt.target)
			assert.Equal(t, tt.authority, authority)
			assert.Equal(t, tt.path, path)
			assert.Equal(t, tt.tlsVerify, tlsVerify)
			assert.Equal(t, tt.wantErr, err)
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
			target:                  "http://localhost:8081",
			appName:                 "my-producer",
			appVersion:              "0.1.0",
			clientID:                "client-id-123",
			clientSecret:            "client-secret-456",
			clientIDEnvVarValue:     "",
			clientSecretEnvVarValue: "",
			wantErr:                 errors.New("target should start with grpc:// or grpcs://"),
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
