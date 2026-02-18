// Package main demonstrates ingesting ModuleTypeProfile entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "module_type_profile-example"
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
	moduleTypeProfile := ModuleTypeProfileMinimal()
	// moduleTypeProfile := ModuleTypeProfileExtended()
	// moduleTypeProfile := ModuleTypeProfileExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{moduleTypeProfile})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ModuleTypeProfile ingested successfully")
	}
}

// ModuleTypeProfileMinimal Creates a ModuleTypeProfile with only required fields.
func ModuleTypeProfileMinimal() *diode.ModuleTypeProfile {
	return &diode.ModuleTypeProfile{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ModuleTypeProfileExtended Creates a ModuleTypeProfile with common optional fields.
func ModuleTypeProfileExtended() *diode.ModuleTypeProfile {
	return &diode.ModuleTypeProfile{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Schema:      diode.String("Example Schema"),
		Comments:    diode.String("Example comments"),
	}
}

// ModuleTypeProfileExplicit Creates a ModuleTypeProfile with fully nested objects and all common fields.
func ModuleTypeProfileExplicit() *diode.ModuleTypeProfile {
	return &diode.ModuleTypeProfile{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Schema:      diode.String("Example Schema"),
		Owner: &diode.Owner{
			Name: diode.String("Example Name"),
			Group: &diode.OwnerGroup{
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
