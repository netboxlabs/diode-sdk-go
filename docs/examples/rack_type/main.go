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
		Model:         diode.String("Model X"),
		Slug:          diode.String("example-slug"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		FormFactor:    diode.String("2-post-frame"),
		Width:         diode.Int64(10),
		UHeight:       diode.Int64(1),
		StartingUnit:  diode.Int64(1),
		DescUnits:     diode.Bool(true),
		OuterWidth:    diode.Int64(1),
		OuterDepth:    diode.Int64(1),
		OuterUnit:     diode.String("in"),
		Weight:        diode.Float64(1.0),
		MaxWeight:     diode.Int64(1),
		WeightUnit:    diode.String("g"),
		MountingDepth: diode.Int64(1),
		Comments:      diode.String("Example comments"),
		OuterHeight:   diode.Int64(1),
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
		Model:         diode.String("Model X"),
		Slug:          diode.String("example-slug"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		FormFactor:    diode.String("2-post-frame"),
		Width:         diode.Int64(10),
		UHeight:       diode.Int64(1),
		StartingUnit:  diode.Int64(1),
		DescUnits:     diode.Bool(true),
		OuterWidth:    diode.Int64(1),
		OuterDepth:    diode.Int64(1),
		OuterUnit:     diode.String("in"),
		Weight:        diode.Float64(1.0),
		MaxWeight:     diode.Int64(1),
		WeightUnit:    diode.String("g"),
		MountingDepth: diode.Int64(1),
		OuterHeight:   diode.Int64(1),
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
