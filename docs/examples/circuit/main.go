// Package main demonstrates ingesting Circuit entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "circuit-example"
	appVersion = "1.0.0"
)

func main() {
	client, err := diode.NewClient(
		target,
		appName,
		appVersion,
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Choose one of the three patterns by uncommenting:
	circuit := CircuitMinimal()
	// circuit := CircuitExtended()
	// circuit := CircuitExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{circuit})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Circuit ingested successfully")
	}
}

// CircuitMinimal Creates a Circuit with only required fields.
func CircuitMinimal() *diode.Circuit {
	return &diode.Circuit{
		Cid: diode.String("CID-001"),
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Type: &diode.CircuitType{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CircuitExtended Creates a Circuit with common optional fields.
func CircuitExtended() *diode.Circuit {
	return &diode.Circuit{
		Cid: diode.String("CID-001"),
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Type: &diode.CircuitType{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
	}
}

// CircuitExplicit Creates a Circuit with fully nested objects and all common fields.
func CircuitExplicit() *diode.Circuit {
	return &diode.Circuit{
		Cid: diode.String("CID-001"),
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Type: &diode.CircuitType{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Color:    diode.String("0000ff"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
