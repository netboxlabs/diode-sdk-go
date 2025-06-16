package diode

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// DiodeDryRunOutpurDirEnvVarName is the environment variable name for the dry run directory
	DiodeDryRunOutpurDirEnvVarName = "DIODE_DRY_RUN_OUTPUT_DIR"
)

// DryRunClient implements Client and writes ingest payloads to stdout or a file.
type DryRunClient struct {
	appName   string
	dryRunDir string
}

// NewDryRunClient creates a new DryRunClient. If dryRunFile is empty it falls
// back to the DIODE_DRY_RUN_FILE environment variable. When no file is
// specified the output is written to STDOUT.
func NewDryRunClient(appName string, dryRunDir string) (Client, error) {
	if appName == "" {
		appName = "dryrun"
	}
	if dryRunDir == "" {
		dryRunDir = os.Getenv(DiodeDryRunOutpurDirEnvVarName)
	}
	if dryRunDir != "" {
		if err := os.MkdirAll(dryRunDir, 0o755); err != nil {
			return nil, err
		}
	}
	return &DryRunClient{appName: appName, dryRunDir: dryRunDir}, nil
}

// Close closes the DryRunClient writer if necessary.
func (d *DryRunClient) Close() error {
	return nil
}

// Ingest writes the given entities as JSON to the configured writer.
// If writing to a file, it appends to an existing JSON array or creates a new one.
func (d *DryRunClient) Ingest(_ context.Context, entities []Entity) (*diodepb.IngestResponse, error) {
	protoEntities := convertEntitiesToProto(entities)
	wrapper := &diodepb.IngestRequest{
		Id:         uuid.New().String(),
		Entities:   protoEntities,
		SdkName:    SDKName,
		SdkVersion: SDKVersion,
		// Add Stream field if it exists in your proto definition
		// Stream: d.stream,
	}

	data, err := protojson.MarshalOptions{UseProtoNames: true, Indent: "  "}.Marshal(wrapper)
	if err != nil {
		return nil, err
	}

	// If writer is a file, handle JSON array appending
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
func LoadDryRunEntities(path string) ([]Entity, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	wrapper := &diodepb.IngestRequest{}
	if err := protojson.Unmarshal(b, wrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}

	return convertProtoEntitiesToEntity(wrapper.Entities), nil
}

// convertProtoEntitiesToEntity converts protobuf entities to Entity implementations
func convertProtoEntitiesToEntity(protoEntities []*diodepb.Entity) []Entity {
	entities := make([]Entity, 0, len(protoEntities))
	for _, protoEntity := range protoEntities {
		entities = append(entities, ProtoEntity{PB: protoEntity})
	}
	return entities
}

// sanitizeAppName sanitizes the application name by replacing
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
