// Package main demonstrates ingesting CircuitGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "circuit_group-example"
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
	circuitGroup := CircuitGroupMinimal()
	// circuitGroup := CircuitGroupExtended()
	// circuitGroup := CircuitGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{circuitGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CircuitGroup ingested successfully")
	}
}

// CircuitGroupMinimal Creates a CircuitGroup with only required fields.
func CircuitGroupMinimal() *diode.CircuitGroup {
	return &diode.CircuitGroup{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CircuitGroupExtended Creates a CircuitGroup with common optional fields.
func CircuitGroupExtended() *diode.CircuitGroup {
	return &diode.CircuitGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// CircuitGroupExplicit Creates a CircuitGroup with fully nested objects and all common fields.
func CircuitGroupExplicit() *diode.CircuitGroup {
	return &diode.CircuitGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
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
