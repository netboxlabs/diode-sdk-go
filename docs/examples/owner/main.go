// Package main demonstrates ingesting Owner entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "owner-example"
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
	owner := OwnerMinimal()
	// owner := OwnerExtended()
	// owner := OwnerExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{owner})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Owner ingested successfully")
	}
}

// OwnerMinimal Creates a Owner with only required fields.
func OwnerMinimal() *diode.Owner {
	return &diode.Owner{
		Name: diode.String("Example Name"),
		Group: &diode.OwnerGroup{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// OwnerExtended Creates a Owner with common optional fields.
func OwnerExtended() *diode.Owner {
	return &diode.Owner{
		Name: diode.String("Example Name"),
		Group: &diode.OwnerGroup{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
	}
}

// OwnerExplicit Creates a Owner with fully nested objects and all common fields.
func OwnerExplicit() *diode.Owner {
	return &diode.Owner{
		Name: diode.String("Example Name"),
		Group: &diode.OwnerGroup{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
	}
}
