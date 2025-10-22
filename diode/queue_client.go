package diode

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	queueClientName       = "diode-sdk-go-queue"
	defaultQueueTimeout   = 10 * time.Second
	defaultQueuePath      = "/ingest"
	queueRequestErr       = "queue request failed"
	defaultQueueUserAgent = queueClientName + "/"
)

// QueueClientOption configures the QueueClient.
type QueueClientOption func(*QueueClient) error

// QueueClient implements Client and forwards ingest payloads to orb-agent via HTTP(S).
type QueueClient struct {
	appName    string
	appVersion string
	queue      string
	stream     string

	logger     *slog.Logger
	httpClient *http.Client
	transport  *http.Transport

	targetURL *url.URL
	rawTarget string

	timeout  time.Duration
	certFile string

	headers http.Header

	sdkName    string
	sdkVersion string

	metadata map[string]string
}

// WithQueueName configures the target queue field.
func WithQueueName(queue string) QueueClientOption {
	return func(c *QueueClient) error {
		c.queue = queue
		return nil
	}
}

// WithQueueTimeout overrides the default HTTP timeout.
func WithQueueTimeout(timeout time.Duration) QueueClientOption {
	return func(c *QueueClient) error {
		if timeout <= 0 {
			return fmt.Errorf("queue timeout must be greater than zero")
		}
		c.timeout = timeout
		return nil
	}
}

// WithQueueHeaders merges custom headers onto the request headers.
func WithQueueHeaders(headers http.Header) QueueClientOption {
	return func(c *QueueClient) error {
		for key, values := range headers {
			c.headers.Del(key)
			for _, v := range values {
				c.headers.Add(key, v)
			}
		}
		return nil
	}
}

// WithQueueCertFile overrides the certificate file used when verifying HTTPS targets.
func WithQueueCertFile(certFile string) QueueClientOption {
	return func(c *QueueClient) error {
		c.certFile = certFile
		return nil
	}
}

// WithQueueStream overrides the stream metadata that is sent alongside entities.
func WithQueueStream(stream string) QueueClientOption {
	return func(c *QueueClient) error {
		if stream == "" {
			return fmt.Errorf("queue stream must not be empty")
		}
		c.stream = stream
		return nil
	}
}

// NewQueueClient creates a new QueueClient that serializes payloads and sends them over HTTP(S).
func NewQueueClient(target string, appName string, appVersion string, opts ...QueueClientOption) (Client, error) {
	if target == "" {
		return nil, fmt.Errorf("target is required")
	}
	parsedTarget, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target: %w", err)
	}
	if parsedTarget.Scheme != "http" && parsedTarget.Scheme != "https" {
		return nil, fmt.Errorf("QueueClient target should start with http:// or https://")
	}
	if parsedTarget.Host == "" {
		return nil, fmt.Errorf("QueueClient target must include a hostname")
	}

	if parsedTarget.Path == "" {
		parsedTarget.Path = defaultQueuePath
	} else if !strings.HasPrefix(parsedTarget.Path, "/") {
		parsedTarget.Path = "/" + parsedTarget.Path
	}

	qc := &QueueClient{
		appName:    appName,
		appVersion: appVersion,
		targetURL:  parsedTarget,
		rawTarget:  target,
		timeout:    defaultQueueTimeout,
		headers:    make(http.Header),
		stream:     defaultStreamName,
		logger:     newLogger(),
		sdkName:    queueClientName,
		sdkVersion: getSDKVersion(),
		metadata: map[string]string{
			"platform":   runtime.GOOS + "/" + runtime.GOARCH,
			"go_version": runtime.Version(),
		},
	}

	qc.headers.Set("Content-Type", "application/json")
	qc.headers.Set("Accept", "application/json")
	qc.headers.Set("User-Agent", fmt.Sprintf("%s%s %s/%s", defaultQueueUserAgent, qc.sdkVersion, qc.appName, qc.appVersion))

	for _, opt := range opts {
		if err := opt(qc); err != nil {
			return nil, err
		}
	}

	if qc.certFile == "" {
		qc.certFile = getCertFile("")
	}

	if err := qc.configureHTTPClient(); err != nil {
		return nil, err
	}

	return qc, nil
}

func (c *QueueClient) configureHTTPClient() error {
	if c.httpClient != nil {
		return nil
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}

	if c.targetURL.Scheme == "https" {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: skipTLSVerify(),
		}
		if !tlsConfig.InsecureSkipVerify {
			rootCAs, err := loadCerts(c.certFile)
			if err != nil {
				return fmt.Errorf("failed to load certificates: %w", err)
			}
			if rootCAs != nil {
				tlsConfig.RootCAs = rootCAs
			}
		}
		transport.TLSClientConfig = tlsConfig
	}

	c.transport = transport
	c.httpClient = &http.Client{
		Timeout:   c.timeout,
		Transport: transport,
	}

	return nil
}

// Close closes idle connections maintained by the client's transport.
func (c *QueueClient) Close() error {
	if c.transport != nil {
		c.transport.CloseIdleConnections()
	}
	return nil
}

// Ingest converts domain entities to proto entities before invoking IngestProto.
func (c *QueueClient) Ingest(ctx context.Context, entities []Entity) (*diodepb.IngestResponse, error) {
	return c.IngestProto(ctx, convertEntitiesToProto(entities))
}

// IngestProto serializes proto entities and sends them to orb-agent via HTTP(S).
func (c *QueueClient) IngestProto(ctx context.Context, entities []*diodepb.Entity) (*diodepb.IngestResponse, error) {
	payload, err := c.serializePayload(entities)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.targetURL.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create queue request: %w", err)
	}

	for key, values := range c.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("queue request failed: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.logger.Debug("Failed to close queue response body", "error", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read queue response: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, &QueueClientError{
			StatusCode:   resp.StatusCode,
			Message:      queueRequestErr,
			ResponseBody: string(body),
		}
	}

	response := &diodepb.IngestResponse{}
	if len(body) == 0 {
		return response, nil
	}

	if err := protojson.Unmarshal(body, response); err != nil {
		c.logger.Debug("Failed to parse queue response into IngestResponse", "error", err)
	}

	return response, nil
}

func (c *QueueClient) serializePayload(entities []*diodepb.Entity) ([]byte, error) {
	jsonEntities := make([]json.RawMessage, 0, len(entities))

	marshalOpts := protojson.MarshalOptions{
		UseProtoNames: true,
	}

	for _, entity := range entities {
		if entity == nil {
			continue
		}
		data, err := marshalOpts.Marshal(entity)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal entity: %w", err)
		}
		jsonEntities = append(jsonEntities, json.RawMessage(data))
	}

	payload := queuePayload{
		ID:     uuid.NewString(),
		Stream: c.stream,
		SDK: queueSDKInfo{
			Name:    c.sdkName,
			Version: c.sdkVersion,
		},
		Producer: queueProducerInfo{
			AppName:    c.appName,
			AppVersion: c.appVersion,
		},
		Metadata: c.metadata,
		Entities: jsonEntities,
	}

	if c.queue != "" {
		payload.Queue = c.queue
	}

	return json.Marshal(payload)
}

// QueueClientError represents a non-successful HTTP response from the queue endpoint.
type QueueClientError struct {
	StatusCode   int
	Message      string
	ResponseBody string
}

// Error implements the error interface.
func (e *QueueClientError) Error() string {
	if e.ResponseBody != "" {
		return fmt.Sprintf("%d %s: %s", e.StatusCode, e.Message, e.ResponseBody)
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.Message)
}

type queuePayload struct {
	ID       string            `json:"id"`
	Queue    string            `json:"queue,omitempty"`
	Stream   string            `json:"stream"`
	SDK      queueSDKInfo      `json:"sdk"`
	Producer queueProducerInfo `json:"producer"`
	Metadata map[string]string `json:"metadata"`
	Entities []json.RawMessage `json:"entities"`
}

type queueSDKInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type queueProducerInfo struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
}
