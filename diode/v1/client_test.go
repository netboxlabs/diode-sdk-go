package diode

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

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

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		desc              string
		apiKey            string
		apiKeyEnvVarValue string
		wantApiKey        string
		wantErr           error
	}{
		{
			desc:              "API key provided explicitly",
			apiKey:            "foobar",
			apiKeyEnvVarValue: "",
			wantApiKey:        "foobar",
			wantErr:           nil,
		},
		{
			desc:              "API key provided with environment variable",
			apiKey:            "",
			apiKeyEnvVarValue: "barfoo",
			wantApiKey:        "barfoo",
			wantErr:           nil,
		},
		{
			desc:              "API key not provided either explicitly or with environment variable",
			apiKey:            "",
			apiKeyEnvVarValue: "",
			wantApiKey:        "",
			wantErr:           errors.New("api_key param or DIODE_API_KEY environment variable required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.apiKeyEnvVarValue != "" {
				_ = os.Setenv(DiodeAPIKeyEnvVarName, tt.apiKeyEnvVarValue)
				defer os.Unsetenv(DiodeAPIKeyEnvVarName)
			}
			apiKey, err := getAPIKey(tt.apiKey)
			require.Equal(t, tt.wantApiKey, apiKey)
			require.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		desc                string
		target              string
		appName             string
		appVersion          string
		apiKey              string
		apiKeyEnvVarValue   string
		logLevelEnvVarValue string
		wantErr             error
	}{
		{
			desc:                "explicit arguments provided",
			target:              "grpc://localhost:8081",
			appName:             "my-producer",
			appVersion:          "0.1.0",
			apiKey:              "foobar",
			apiKeyEnvVarValue:   "",
			logLevelEnvVarValue: "",
			wantErr:             nil,
		},
		{
			desc:                "API key provided with environment variable",
			target:              "grpc://localhost:8081",
			appName:             "my-producer",
			appVersion:          "0.1.0",
			apiKey:              "",
			apiKeyEnvVarValue:   "foo.bar",
			logLevelEnvVarValue: "",
			wantErr:             nil,
		},
		{
			desc:                "target with path",
			target:              "grpc://localhost:8081/abcdef",
			appName:             "my-producer",
			appVersion:          "0.1.0",
			apiKey:              "",
			apiKeyEnvVarValue:   "foo.bar",
			logLevelEnvVarValue: "",
			wantErr:             nil,
		},
		{
			desc:                "target with grpcs scheme",
			target:              "grpcs://localhost:8081",
			appName:             "my-producer",
			appVersion:          "0.1.0",
			apiKey:              "",
			apiKeyEnvVarValue:   "foo.bar",
			logLevelEnvVarValue: "",
			wantErr:             nil,
		},
		{
			desc:                "app name not provided",
			target:              "grpc://localhost:8081",
			appName:             "",
			appVersion:          "0.1.0",
			apiKey:              "foobar",
			apiKeyEnvVarValue:   "",
			logLevelEnvVarValue: "",
			wantErr:             errors.New("app name is required"),
		},
		{
			desc:                "app version not provided",
			target:              "grpc://localhost:8081",
			appName:             "my-producer",
			appVersion:          "",
			apiKey:              "foobar",
			apiKeyEnvVarValue:   "",
			logLevelEnvVarValue: "",
			wantErr:             errors.New("app version is required"),
		},
		{
			desc:                "invalid target",
			target:              "http://localhost:8081",
			appName:             "my-producer",
			appVersion:          "0.1.0",
			apiKey:              "foobar",
			apiKeyEnvVarValue:   "",
			logLevelEnvVarValue: "",
			wantErr:             errors.New("target should start with grpc:// or grpcs://"),
		},
		{
			desc:              "missing API key",
			target:            "grpc://localhost:8081",
			appName:           "my-producer",
			appVersion:        "0.1.0",
			apiKey:            "",
			apiKeyEnvVarValue: "",
			wantErr:           errors.New("api_key param or DIODE_API_KEY environment variable required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			defer os.Unsetenv(DiodeAPIKeyEnvVarName)
			defer os.Unsetenv(DiodeSDKLogLevelEnvVarName)

			if tt.apiKeyEnvVarValue != "" {
				_ = os.Setenv(DiodeAPIKeyEnvVarName, tt.apiKeyEnvVarValue)
			}

			client, err := NewClient(tt.target, tt.appName, tt.appVersion, tt.apiKey)
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
}

func (MockIngesterServiceServer) Ingest(_ context.Context, req *diodepb.IngestRequest) (*diodepb.IngestResponse, error) {
	return &diodepb.IngestResponse{Errors: nil}, nil
}

func startMockServer() (net.Listener, error) {
	port, _ := getFreePort()
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen on port %s: %v", port, err)
	}

	server := grpc.NewServer()

	diodepb.RegisterIngesterServiceServer(server, &MockIngesterServiceServer{})

	go func() {
		if err := server.Serve(grpcListener); err != nil {
			log.Fatal(err)
		}
	}()

	return grpcListener, nil
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
			listener, err := startMockServer()
			require.NoError(t, err)

			target := fmt.Sprintf("grpc://%s/%s", listener.Addr().String(), tt.path)
			appName := "my-producer"
			appVersion := "0.1.0"
			apiKey := "abcde"

			client, err := NewClient(target, appName, appVersion, apiKey)
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
			wantLogLevel:        slog.LevelDebug,
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
			defer os.Unsetenv(DiodeSDKLogLevelEnvVarName)

			if tt.logLevelEnvVarValue != "" {
				_ = os.Setenv(DiodeSDKLogLevelEnvVarName, tt.logLevelEnvVarValue)
			}

			logger := newLogger()
			require.NotNil(t, logger)
			assert.True(t, logger.Enabled(context.Background(), tt.wantLogLevel))
		})
	}
}
