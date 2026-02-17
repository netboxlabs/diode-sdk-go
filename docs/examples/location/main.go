// Package main demonstrates ingesting Location entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "location-example"
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
	location := LocationMinimal()
	// location := LocationExtended()
	// location := LocationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{location})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Location ingested successfully")
	}
}

// LocationMinimal Creates a Location with only required fields.
func LocationMinimal() *diode.Location {
	return &diode.Location{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// LocationExtended Creates a Location with common optional fields.
func LocationExtended() *diode.Location {
	return &diode.Location{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
	}
}

// LocationExplicit Creates a Location with fully nested objects and all common fields.
func LocationExplicit() *diode.Location {
	return &diode.Location{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
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
