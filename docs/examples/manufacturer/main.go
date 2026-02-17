// Package main demonstrates ingesting Manufacturer entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "manufacturer-example"
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
	manufacturer := ManufacturerMinimal()
	// manufacturer := ManufacturerExtended()
	// manufacturer := ManufacturerExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{manufacturer})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Manufacturer ingested successfully")
	}
}

// ManufacturerMinimal Creates a Manufacturer with only required fields.
func ManufacturerMinimal() *diode.Manufacturer {
	return &diode.Manufacturer{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ManufacturerExtended Creates a Manufacturer with common optional fields.
func ManufacturerExtended() *diode.Manufacturer {
	return &diode.Manufacturer{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// ManufacturerExplicit Creates a Manufacturer with fully nested objects and all common fields.
func ManufacturerExplicit() *diode.Manufacturer {
	return &diode.Manufacturer{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
