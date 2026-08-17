package diode

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

// makeJWT builds an unsigned JWT carrying only an exp claim. The SDK reads exp
// to schedule renewal and never verifies the signature, so a placeholder is
// enough.
func makeJWT(exp time.Time) string {
	enc := func(v any) string {
		b, _ := json.Marshal(v)
		return base64.RawURLEncoding.EncodeToString(b)
	}
	header := enc(map[string]string{"alg": "RS256", "typ": "JWT"})
	payload := enc(map[string]int64{"exp": exp.Unix()})
	return header + "." + payload + ".not-a-real-signature"
}

// countingIngester fails every Ingest with the configured status.
type countingIngester struct {
	diodepb.UnimplementedIngesterServiceServer
	calls atomic.Int64
	err   error
}

func (m *countingIngester) Ingest(_ context.Context, _ *diodepb.IngestRequest) (*diodepb.IngestResponse, error) {
	m.calls.Add(1)
	if m.err != nil {
		return nil, m.err
	}
	return &diodepb.IngestResponse{}, nil
}

// refreshTestServer serves the token endpoint and the ingester on one cleartext
// listener, counting token requests so a test can assert renewal happened.
type refreshTestServer struct {
	tokenRequests atomic.Int64
	ingester      *countingIngester
	httpServer    *http.Server
	port          string
}

// startRefreshTestServer issues tokens whose exp is now+tokenTTL. A negative or
// small tokenTTL produces a token the SDK should consider stale.
func startRefreshTestServer(t *testing.T, tokenTTL time.Duration, ingestErr error) *refreshTestServer {
	t.Helper()

	port, err := getFreePort()
	require.NoError(t, err)

	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	require.NoError(t, err)

	s := &refreshTestServer{
		ingester: &countingIngester{err: ingestErr},
		port:     port,
	}

	grpcServer := grpc.NewServer()
	diodepb.RegisterIngesterServiceServer(grpcServer, s.ingester)

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/token", func(w http.ResponseWriter, _ *http.Request) {
		s.tokenRequests.Add(1)
		w.Header().Set("Content-Type", "application/json")
		body := fmt.Sprintf(`{"access_token": %q}`, makeJWT(time.Now().Add(tokenTTL)))
		_, _ = w.Write([]byte(body))
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		}
		mux.ServeHTTP(w, r)
	})

	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	s.httpServer = &http.Server{Handler: handler, Protocols: protocols}

	go func() { _ = s.httpServer.Serve(listener) }()
	t.Cleanup(func() { _ = s.httpServer.Close() })

	return s
}

func newRefreshTestClient(t *testing.T, s *refreshTestServer) *GRPCClient {
	t.Helper()

	t.Setenv(DiodeClientIDEnvVarName, "client-id")
	t.Setenv(DiodeClientSecretEnvVarName, "client-secret")
	t.Cleanup(func() {
		_ = os.Unsetenv(DiodeClientIDEnvVarName)
		_ = os.Unsetenv(DiodeClientSecretEnvVarName)
	})

	client, err := NewClient(fmt.Sprintf("grpc://localhost:%s", s.port), "my-producer", "0.1.0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = client.Close() })

	return client.(*GRPCClient)
}

func TestAccessTokenExpiry(t *testing.T) {
	exp := time.Now().Add(42 * time.Minute).Truncate(time.Second)

	tests := []struct {
		desc  string
		token string
		want  time.Time
	}{
		{"jwt with exp", makeJWT(exp), exp},
		{"not a jwt", "opaque-token", time.Time{}},
		{"wrong segment count", "only.two", time.Time{}},
		{"payload not base64", "aaa.!!!not-base64!!!.ccc", time.Time{}},
		{"payload not json", "aaa." + base64.RawURLEncoding.EncodeToString([]byte("nope")) + ".ccc", time.Time{}},
		{"no exp claim", "aaa." + base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"x"}`)) + ".ccc", time.Time{}},
		{"zero exp claim", "aaa." + base64.RawURLEncoding.EncodeToString([]byte(`{"exp":0}`)) + ".ccc", time.Time{}},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := accessTokenExpiry(tt.token)
			if tt.want.IsZero() {
				assert.True(t, got.IsZero(), "expected zero time, got %v", got)
				return
			}
			assert.True(t, got.Equal(tt.want), "expected %v, got %v", tt.want, got)
		})
	}
}

// A token that is still comfortably valid must not be renewed on every ingest.
func TestFreshTokenIsNotRefreshed(t *testing.T) {
	s := startRefreshTestServer(t, time.Hour, nil)
	client := newRefreshTestClient(t, s)

	require.Equal(t, int64(1), s.tokenRequests.Load(), "construction should authenticate once")

	_, err := client.Ingest(context.Background(), []Entity{&Site{Name: String("site-1")}})
	require.NoError(t, err)

	assert.Equal(t, int64(1), s.tokenRequests.Load(), "a fresh token should be reused")
	assert.Equal(t, int64(1), s.ingester.calls.Load())
}

// The token is renewed before it expires, so the request never carries a token
// the edge would reject.
func TestStaleTokenIsRefreshedBeforeIngest(t *testing.T) {
	// Inside tokenRefreshWindow the moment it is issued.
	s := startRefreshTestServer(t, 10*time.Second, nil)
	client := newRefreshTestClient(t, s)

	require.Equal(t, int64(1), s.tokenRequests.Load())

	_, err := client.Ingest(context.Background(), []Entity{&Site{Name: String("site-1")}})
	require.NoError(t, err)

	assert.Equal(t, int64(2), s.tokenRequests.Load(), "a stale token should be renewed before ingesting")
	assert.Equal(t, int64(1), s.ingester.calls.Load(), "the ingest itself should not be retried")
}

// An intermediary that rejects an expired token with a plain HTTP response can
// leave the client holding a generic Internal error instead of Unauthenticated.
// A stale token must still trigger renewal in that case.
func TestStaleTokenRetriesOnNonUnauthenticatedError(t *testing.T) {
	opaque := status.Error(codes.Internal, "server closed the stream without sending trailers")
	s := startRefreshTestServer(t, 10*time.Second, opaque)
	client := newRefreshTestClient(t, s)
	client.maxAuthRetries = 3

	_, err := client.Ingest(context.Background(), []Entity{&Site{Name: String("site-1")}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "authentication failed after 3 attempts")

	assert.Greater(t, s.tokenRequests.Load(), int64(1), "an opaque failure on a stale token should renew it")
}

// The counterpart hazard: with a healthy token, an unrelated failure must stay
// fatal rather than being mistaken for an auth problem and retried.
func TestFreshTokenDoesNotRetryOnNonUnauthenticatedError(t *testing.T) {
	s := startRefreshTestServer(t, time.Hour, status.Error(codes.Internal, "boom"))
	client := newRefreshTestClient(t, s)
	client.maxAuthRetries = 3

	_, err := client.Ingest(context.Background(), []Entity{&Site{Name: String("site-1")}})
	require.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.NotContains(t, err.Error(), "authentication failed after")

	assert.Equal(t, int64(1), s.tokenRequests.Load(), "no renewal for an unrelated failure")
	assert.Equal(t, int64(1), s.ingester.calls.Load(), "no retry for an unrelated failure")
}

// An opaque token carries no exp, so behaviour must be unchanged: no proactive
// renewal, and only Unauthenticated drives a retry.
func TestOpaqueTokenLeavesBehaviourUnchanged(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err)
	authServer, err := startMockAuthServer(port, "", false)
	require.NoError(t, err)
	t.Cleanup(func() { authServer.Close() })

	t.Setenv(DiodeClientIDEnvVarName, "client-id")
	t.Setenv(DiodeClientSecretEnvVarName, "client-secret")

	client, err := NewClient(fmt.Sprintf("grpc://localhost:%s", port), "my-producer", "0.1.0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = client.Close() })

	grpcClient := client.(*GRPCClient)
	assert.True(t, grpcClient.tokenExpiry.IsZero(), "an opaque token yields no expiry")
	assert.False(t, grpcClient.tokenStale(), "unknown expiry must not read as stale")

	_, err = grpcClient.Ingest(context.Background(), []Entity{&Site{Name: String("site-1")}})
	require.NoError(t, err)
}
