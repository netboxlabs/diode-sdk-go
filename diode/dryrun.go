package diode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

const (
	// DiodeDryRunFileEnvVarName is the environment variable name for the dry run file path
	DiodeDryRunFileEnvVarName = "DIODE_DRY_RUN_FILE"
)

// DryRunClient implements Client and writes ingest payloads to stdout or a file.
type DryRunClient struct {
	writer io.WriteCloser
}

// nopWriteCloser wraps an io.Writer to satisfy io.WriteCloser without closing.
type nopWriteCloser struct{ io.Writer }

func (nopWriteCloser) Close() error { return nil }

// NewDryRunClient creates a new DryRunClient. If dryRunFile is empty it falls
// back to the DIODE_DRY_RUN_FILE environment variable. When no file is
// specified the output is written to STDOUT.
func NewDryRunClient(dryRunFile string) (Client, error) {
	if dryRunFile == "" {
		dryRunFile = os.Getenv(DiodeDryRunFileEnvVarName)
	}
	if dryRunFile == "" {
		return &DryRunClient{writer: nopWriteCloser{os.Stdout}}, nil
	}
	f, err := os.Create(dryRunFile)
	if err != nil {
		return nil, err
	}
	return &DryRunClient{writer: f}, nil
}

// Close closes the DryRunClient writer if necessary.
func (d *DryRunClient) Close() error {
	if d.writer != nil {
		return d.writer.Close()
	}
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
	if file, ok := d.writer.(*os.File); ok {
		stat, err := file.Stat()
		if err != nil {
			return nil, err
		}

		if stat.Size() == 0 {
			// New file, start JSON array
			if _, err := file.Write([]byte("[\n")); err != nil {
				return nil, err
			}
			if _, err := file.Write(data); err != nil {
				return nil, err
			}
			if _, err := file.Write([]byte("\n]")); err != nil {
				return nil, err
			}
		} else {
			// Existing file, append to JSON array
			// Seek to check last 2 bytes
			if _, err := file.Seek(-2, io.SeekEnd); err != nil {
				return nil, err
			}

			trailer := make([]byte, 2)
			if _, err := file.Read(trailer); err != nil {
				return nil, err
			}

			if string(trailer) != "\n]" {
				return &diodepb.IngestResponse{
					Errors: []string{"Invalid JSON trailer in dry run output file"},
				}, nil
			}

			// Seek back and overwrite the closing bracket
			if _, err := file.Seek(-2, io.SeekEnd); err != nil {
				return nil, err
			}

			if _, err := file.Write([]byte(",\n")); err != nil {
				return nil, err
			}
			if _, err := file.Write(data); err != nil {
				return nil, err
			}
			if _, err := file.Write([]byte("\n]")); err != nil {
				return nil, err
			}
		}
	} else {
		// Not a file (e.g., stdout), just write the JSON
		if _, err := d.writer.Write(data); err != nil {
			return nil, err
		}
		if _, err := d.writer.Write([]byte("\n")); err != nil {
			return nil, err
		}
	}

	return &diodepb.IngestResponse{}, nil
}

// LoadDryRunEntities loads entities written by DryRunClient from the file path
// and returns them as protobuf entities from all IngestRequests in the JSON array.
func LoadDryRunEntities(path string) ([]*diodepb.Entity, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse as JSON array
	var requestsJSON []json.RawMessage
	if err := json.Unmarshal(b, &requestsJSON); err != nil {
		return nil, err
	}

	var allEntities []*diodepb.Entity

	// Convert each JSON object to IngestRequest and collect entities
	for _, reqJSON := range requestsJSON {
		wrapper := &diodepb.IngestRequest{}
		if err := protojson.Unmarshal(reqJSON, wrapper); err != nil {
			return nil, fmt.Errorf("failed to unmarshal request: %w", err)
		}
		allEntities = append(allEntities, wrapper.Entities...)
	}

	return allEntities, nil
}
