// Package main demonstrates ingesting ModuleType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "module_type-example"
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
	moduleType := ModuleTypeMinimal()
	// moduleType := ModuleTypeExtended()
	// moduleType := ModuleTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{moduleType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ModuleType ingested successfully")
	}
}

// ModuleTypeMinimal Creates a ModuleType with only required fields.
func ModuleTypeMinimal() *diode.ModuleType {
	return &diode.ModuleType{
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Model: diode.String("Model X"),
	}
}

// ModuleTypeExtended Creates a ModuleType with common optional fields.
func ModuleTypeExtended() *diode.ModuleType {
	return &diode.ModuleType{
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Model:       diode.String("Model X"),
		Description: diode.String("Example description"),
	}
}

// ModuleTypeExplicit Creates a ModuleType with fully nested objects and all common fields.
func ModuleTypeExplicit() *diode.ModuleType {
	return &diode.ModuleType{
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Model:       diode.String("Model X"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
