package diode

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueueClientIngestSerializesEntities(t *testing.T) {
	var (
		requestBody    []byte
		requestHeaders http.Header
		requestPath    string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_ = r.Body.Close()
		}()

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		requestBody = body
		requestHeaders = r.Header.Clone()
		requestPath = r.URL.Path

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"errors": []}`))
	}))
	defer server.Close()

	client, err := NewQueueClient(
		server.URL+"/queue",
		"orb-producer",
		"1.2.3",
		WithQueueName("orb"),
		WithQueueTimeout(2*time.Second),
	)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, client.Close())
	}()

	resp, err := client.Ingest(context.Background(), []Entity{
		&Site{Name: String("Site1")},
		&Device{Name: String("Device1")},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Empty(t, resp.GetErrors())

	assert.Equal(t, "/queue", requestPath)
	assert.Equal(t, "application/json", requestHeaders.Get("Content-Type"))

	var payload map[string]any
	require.NoError(t, json.Unmarshal(requestBody, &payload))

	assert.Equal(t, "orb", payload["queue"])
	assert.Equal(t, defaultStreamName, payload["stream"])

	entities, ok := payload["entities"].([]any)
	require.True(t, ok)
	require.Len(t, entities, 2)

	firstEntity, ok := entities[0].(map[string]any)
	require.True(t, ok)
	site, ok := firstEntity["site"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "Site1", site["name"])

	metadata, ok := payload["metadata"].(map[string]any)
	require.True(t, ok)
	assert.NotEmpty(t, metadata["go_version"])
}

func TestQueueClientReturnsDetailedError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"detail": "failed"}`))
	}))
	defer server.Close()

	client, err := NewQueueClient(server.URL, "orb-producer", "1.2.3")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, client.Close())
	}()

	_, err = client.Ingest(context.Background(), []Entity{&Site{Name: String("Site1")}})
	require.Error(t, err)

	var queueErr *QueueClientError
	require.ErrorAs(t, err, &queueErr)
	assert.Equal(t, http.StatusInternalServerError, queueErr.StatusCode)
	assert.Contains(t, queueErr.Error(), "queue request failed")
	assert.Contains(t, queueErr.ResponseBody, "failed")
}

func TestQueueClientHTTPSHonoursSkipTLSVerify(t *testing.T) {
	t.Setenv(DiodeSkipTLSVerifyEnvVarName, "1")

	var receivedPath string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := NewQueueClient(server.URL, "orb-producer", "1.2.3")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, client.Close())
	}()

	resp, err := client.Ingest(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, defaultQueuePath, receivedPath)
}
