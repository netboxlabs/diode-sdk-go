package diode

import (
	"context"
	"io"
	"os"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
)

// DryRunClient implements Client and writes ingest payloads to stdout or a file.
type DryRunClient struct {
	writer io.WriteCloser
}

// nopWriteCloser wraps an io.Writer to satisfy io.WriteCloser without closing.
type nopWriteCloser struct{ io.Writer }

func (nopWriteCloser) Close() error { return nil }

// NewDryRunClient creates a new DryRunClient. If dryRunFile is empty the
// output is written to STDOUT otherwise to the specified file.
func NewDryRunClient(dryRunFile string) (Client, error) {
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
func (d *DryRunClient) Ingest(_ context.Context, entities []Entity) (*diodepb.IngestResponse, error) {
	protoEntities := convertEntitiesToProto(entities)
	wrapper := &diodepb.IngestRequest{Entities: protoEntities}
	data, err := protojson.MarshalOptions{Indent: "  "}.Marshal(wrapper)
	if err != nil {
		return nil, err
	}
	if _, err = d.writer.Write(data); err != nil {
		return nil, err
	}
	if _, err = d.writer.Write([]byte("\n")); err != nil {
		return nil, err
	}
	return &diodepb.IngestResponse{}, nil
}

// LoadDryRunEntities loads entities written by DryRunClient from the file path
// and returns them as protobuf entities.
func LoadDryRunEntities(path string) ([]*diodepb.Entity, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	wrapper := &diodepb.IngestRequest{}
	if err := protojson.Unmarshal(b, wrapper); err != nil {
		return nil, err
	}
	return wrapper.Entities, nil
}
