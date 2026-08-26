// Package main demonstrates ingesting CoolingSource entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cooling_source-example"
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
	coolingSource := CoolingSourceMinimal()
	// coolingSource := CoolingSourceExtended()
	// coolingSource := CoolingSourceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{coolingSource})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CoolingSource ingested successfully")
	}
}

// CoolingSourceMinimal Creates a CoolingSource with only required fields.
func CoolingSourceMinimal() *diode.CoolingSource {
	return &diode.CoolingSource{
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Type:     diode.String("chiller"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CoolingSourceExtended Creates a CoolingSource with common optional fields.
func CoolingSourceExtended() *diode.CoolingSource {
	return &diode.CoolingSource{
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:            diode.String("Example Name"),
		Type:            diode.String("chiller"),
		Metadata:        diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:          diode.String("active"),
		Description:     diode.String("Example description"),
		FluidType:       diode.String("dielectric"),
		CoolingCapacity: diode.Float64(1.0),
		Comments:        diode.String("Example comments"),
	}
}

// CoolingSourceExplicit Creates a CoolingSource with fully nested objects and all common fields.
func CoolingSourceExplicit() *diode.CoolingSource {
	return &diode.CoolingSource{
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:            diode.String("Example Name"),
		Type:            diode.String("chiller"),
		Metadata:        diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:          diode.String("active"),
		Description:     diode.String("Example description"),
		Comments:        diode.String("Example comments"),
		FluidType:       diode.String("dielectric"),
		CoolingCapacity: diode.Float64(1.0),
		Location: &diode.Location{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
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
