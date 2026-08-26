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
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Model:    diode.String("Model X"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ModuleTypeExtended Creates a ModuleType with common optional fields.
func ModuleTypeExtended() *diode.ModuleType {
	return &diode.ModuleType{
		Manufacturer: &diode.Manufacturer{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Model:         diode.String("Model X"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		PartNumber:    diode.String("Example PartNumber"),
		Airflow:       diode.String("front-to-rear"),
		Weight:        diode.Float64(1.0),
		WeightUnit:    diode.String("g"),
		Comments:      diode.String("Example comments"),
		Attributes:    diode.String("Example Attributes"),
		CoolingMethod: diode.String("air"),
	}
}

// ModuleTypeExplicit Creates a ModuleType with fully nested objects and all common fields.
func ModuleTypeExplicit() *diode.ModuleType {
	return &diode.ModuleType{
		Manufacturer: &diode.Manufacturer{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Model:         diode.String("Model X"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		PartNumber:    diode.String("Example PartNumber"),
		Airflow:       diode.String("front-to-rear"),
		Weight:        diode.Float64(1.0),
		WeightUnit:    diode.String("g"),
		Attributes:    diode.String("Example Attributes"),
		CoolingMethod: diode.String("air"),
		Profile: &diode.ModuleTypeProfile{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
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
