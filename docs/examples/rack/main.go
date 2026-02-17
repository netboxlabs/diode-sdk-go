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
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
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
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		AssetTag:    diode.String("ASSET-001"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
