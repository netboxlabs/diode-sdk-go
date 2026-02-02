// Package main demonstrates ingesting DeviceType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "device_type-example"
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
	deviceType := DeviceTypeMinimal()
	// deviceType := DeviceTypeExtended()
	// deviceType := DeviceTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{deviceType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("DeviceType ingested successfully")
	}
}

// DeviceTypeMinimal Creates a DeviceType with only required fields.
func DeviceTypeMinimal() *diode.DeviceType {
	return &diode.DeviceType{
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Model: diode.String("Model X"),
		Slug:  diode.String("example-slug"),
	}
}

// DeviceTypeExtended Creates a DeviceType with common optional fields.
func DeviceTypeExtended() *diode.DeviceType {
	return &diode.DeviceType{
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Model:       diode.String("Model X"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// DeviceTypeExplicit Creates a DeviceType with fully nested objects and all common fields.
func DeviceTypeExplicit() *diode.DeviceType {
	return &diode.DeviceType{
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Model:       diode.String("Model X"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
