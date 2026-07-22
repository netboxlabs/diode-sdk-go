// Package main demonstrates ingesting Rack entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "rack-example"
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
	rack := RackMinimal()
	// rack := RackExtended()
	// rack := RackExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{rack})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Rack ingested successfully")
	}
}

// RackMinimal Creates a Rack with only required fields.
func RackMinimal() *diode.Rack {
	return &diode.Rack{
		Name: diode.String("Example Name"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// RackExtended Creates a Rack with common optional fields.
func RackExtended() *diode.Rack {
	return &diode.Rack{
		Name: diode.String("Example Name"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:        diode.String("active"),
		Serial:        diode.String("SN-001234"),
		Description:   diode.String("Example description"),
		FacilityId:    diode.String("Example FacilityId"),
		AssetTag:      diode.String("ASSET-001"),
		FormFactor:    diode.String("2-post-frame"),
		Width:         diode.Int64(10),
		UHeight:       diode.Int64(1),
		StartingUnit:  diode.Int64(1),
		Weight:        diode.Float64(1.0),
		MaxWeight:     diode.Int64(1),
		WeightUnit:    diode.String("g"),
		DescUnits:     diode.Bool(true),
		OuterWidth:    diode.Int64(1),
		OuterDepth:    diode.Int64(1),
		OuterUnit:     diode.String("in"),
		MountingDepth: diode.Int64(1),
		Airflow:       diode.String("front-to-rear"),
		Comments:      diode.String("Example comments"),
		OuterHeight:   diode.Int64(1),
	}
}

// RackExplicit Creates a Rack with fully nested objects and all common fields.
func RackExplicit() *diode.Rack {
	return &diode.Rack{
		Name: diode.String("Example Name"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:        diode.String("active"),
		Serial:        diode.String("SN-001234"),
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		AssetTag:      diode.String("ASSET-001"),
		FacilityId:    diode.String("Example FacilityId"),
		FormFactor:    diode.String("2-post-frame"),
		Width:         diode.Int64(10),
		UHeight:       diode.Int64(1),
		StartingUnit:  diode.Int64(1),
		Weight:        diode.Float64(1.0),
		MaxWeight:     diode.Int64(1),
		WeightUnit:    diode.String("g"),
		DescUnits:     diode.Bool(true),
		OuterWidth:    diode.Int64(1),
		OuterDepth:    diode.Int64(1),
		OuterUnit:     diode.String("in"),
		MountingDepth: diode.Int64(1),
		Airflow:       diode.String("front-to-rear"),
		OuterHeight:   diode.Int64(1),
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
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Role: &diode.RackRole{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Color:    diode.String("0000ff"),
			Metadata: diode.Metadata{"source": "example"},
		},
		RackType: &diode.RackType{
			Manufacturer: &diode.Manufacturer{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Model:    diode.String("Model X"),
			Slug:     diode.String("example-slug"),
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
		Group: &diode.RackGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
