// Package main demonstrates ingesting CircuitType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "circuit_type-example"
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
	circuitType := CircuitTypeMinimal()
	// circuitType := CircuitTypeExtended()
	// circuitType := CircuitTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{circuitType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CircuitType ingested successfully")
	}
}

// CircuitTypeMinimal Creates a CircuitType with only required fields.
func CircuitTypeMinimal() *diode.CircuitType {
	return &diode.CircuitType{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CircuitTypeExtended Creates a CircuitType with common optional fields.
func CircuitTypeExtended() *diode.CircuitType {
	return &diode.CircuitType{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}

// CircuitTypeExplicit Creates a CircuitType with fully nested objects and all common fields.
func CircuitTypeExplicit() *diode.CircuitType {
	return &diode.CircuitType{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
