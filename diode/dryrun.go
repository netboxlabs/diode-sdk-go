package diode

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// DiodeDryRunOutputDirEnvVarName is the environment variable name for the dry run directory
	DiodeDryRunOutputDirEnvVarName = "DIODE_DRY_RUN_OUTPUT_DIR"
)

// DryRunClient implements Client and writes ingest payloads to stdout or a file.
type DryRunClient struct {
	appName    string
	dryRunDir  string
	sdkName    string
	sdkVersion string
}

// NewDryRunClient creates a new DryRunClient.
func NewDryRunClient(appName string, dryRunDir string) (Client, error) {
	if appName == "" {
		appName = "dryrun"
	}
	if dryRunDir == "" {
		dryRunDir = os.Getenv(DiodeDryRunOutputDirEnvVarName)
	}
	if dryRunDir != "" {
		if err := os.MkdirAll(dryRunDir, 0o755); err != nil {
			return nil, err
		}
	}
	return &DryRunClient{
		appName:    appName,
		dryRunDir:  dryRunDir,
		sdkName:    SDKName,
		sdkVersion: getSDKVersion(),
	}, nil
}

// Close closes the DryRunClient writer if necessary.
func (d *DryRunClient) Close() error {
	return nil
}

// Ingest writes the given entities to stdout or a file depending on the configuration.
// This is a wrapper around IngestProto that converts the entities to protobuf first.
func (d *DryRunClient) Ingest(ctx context.Context, entities []Entity, opts ...IngestOption) (*diodepb.IngestResponse, error) {
	return d.IngestProto(ctx, convertEntitiesToProto(entities), opts...)
}

// IngestProto serializes entities as JSON and writes them to stdout or a file.
// If a directory is configured, it creates a new timestamped file in that directory.
// Otherwise, it writes to stdout.
func (d *DryRunClient) IngestProto(_ context.Context, entities []*diodepb.Entity, opts ...IngestOption) (*diodepb.IngestResponse, error) {
	// Apply options
	cfg := &ingestConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	wrapper := &diodepb.IngestRequest{
		Id:         uuid.New().String(),
		Entities:   entities,
		SdkName:    d.sdkName,
		SdkVersion: d.sdkVersion,
	}

	// Add metadata to request if provided
	if len(cfg.metadata) > 0 {
		wrapper.Metadata, _ = structpb.NewStruct(cfg.metadata)
	}

	data, err := protojson.MarshalOptions{UseProtoNames: true, Indent: "  "}.Marshal(wrapper)
	if err != nil {
		return nil, err
	}

	if d.dryRunDir != "" {
		fileName := fmt.Sprintf("%s_%d.json", sanitizeAppName(d.appName), time.Now().UnixNano())
		path := filepath.Join(d.dryRunDir, fileName)
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return nil, err
		}
	} else {
		// Not a file (e.g., stdout), just write the JSON
		if _, err := os.Stdout.Write(append(data, '\n')); err != nil {
			return nil, err
		}
	}

	return &diodepb.IngestResponse{}, nil
}

// LoadDryRunEntities loads entities written by DryRunClient from a single file
// produced by Ingest.
func LoadDryRunEntities(path string) ([]*diodepb.Entity, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	wrapper := &diodepb.IngestRequest{}
	if err := protojson.Unmarshal(b, wrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}

	return wrapper.Entities, nil
}

// sanitizeAppName sanitizes the application name by replacing invalid characters
// with underscores. Valid characters are letters, digits, underscores, and hyphens.
func sanitizeAppName(appName string) string {
	var b strings.Builder
	for _, c := range appName {
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || c == '-' {
			b.WriteRune(c)
		} else {
			b.WriteRune('_')
		}
	}
	return b.String()
}
