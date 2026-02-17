// Package main demonstrates ingesting Cable entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cable-example"
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
	cable := CableMinimal()
	// cable := CableExtended()
	// cable := CableExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{cable})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Cable ingested successfully")
	}
}

// CableMinimal Creates a Cable with only required fields.
func CableMinimal() *diode.Cable {
	return &diode.Cable{
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CableExtended Creates a Cable with common optional fields.
func CableExtended() *diode.Cable {
	return &diode.Cable{
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("connected"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}

// CableExplicit Creates a Cable with fully nested objects and all common fields.
func CableExplicit() *diode.Cable {
	return &diode.Cable{
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("connected"),
		Color:       diode.String("0000ff"),
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
