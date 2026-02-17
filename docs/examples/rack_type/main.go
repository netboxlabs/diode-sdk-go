// Package main demonstrates ingesting RackType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "rack_type-example"
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
	rackType := RackTypeMinimal()
	// rackType := RackTypeExtended()
	// rackType := RackTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{rackType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("RackType ingested successfully")
	}
}

// RackTypeMinimal Creates a RackType with only required fields.
func RackTypeMinimal() *diode.RackType {
	return &diode.RackType{
		Manufacturer: &diode.Manufacturer{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Model:    diode.String("Model X"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// RackTypeExtended Creates a RackType with common optional fields.
func RackTypeExtended() *diode.RackType {
	return &diode.RackType{
		Manufacturer: &diode.Manufacturer{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Model:       diode.String("Model X"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// RackTypeExplicit Creates a RackType with fully nested objects and all common fields.
func RackTypeExplicit() *diode.RackType {
	return &diode.RackType{
		Manufacturer: &diode.Manufacturer{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Model:       diode.String("Model X"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
