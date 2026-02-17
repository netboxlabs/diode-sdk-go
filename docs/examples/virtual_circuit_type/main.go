// Package main demonstrates ingesting VirtualCircuitType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_circuit_type-example"
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
	virtualCircuitType := VirtualCircuitTypeMinimal()
	// virtualCircuitType := VirtualCircuitTypeExtended()
	// virtualCircuitType := VirtualCircuitTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualCircuitType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualCircuitType ingested successfully")
	}
}

// VirtualCircuitTypeMinimal Creates a VirtualCircuitType with only required fields.
func VirtualCircuitTypeMinimal() *diode.VirtualCircuitType {
	return &diode.VirtualCircuitType{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VirtualCircuitTypeExtended Creates a VirtualCircuitType with common optional fields.
func VirtualCircuitTypeExtended() *diode.VirtualCircuitType {
	return &diode.VirtualCircuitType{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}

// VirtualCircuitTypeExplicit Creates a VirtualCircuitType with fully nested objects and all common fields.
func VirtualCircuitTypeExplicit() *diode.VirtualCircuitType {
	return &diode.VirtualCircuitType{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
