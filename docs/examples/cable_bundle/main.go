// Package main demonstrates ingesting CableBundle entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cable_bundle-example"
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
	cableBundle := CableBundleMinimal()
	// cableBundle := CableBundleExtended()
	// cableBundle := CableBundleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{cableBundle})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CableBundle ingested successfully")
	}
}

// CableBundleMinimal Creates a CableBundle with only required fields.
func CableBundleMinimal() *diode.CableBundle {
	return &diode.CableBundle{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CableBundleExtended Creates a CableBundle with common optional fields.
func CableBundleExtended() *diode.CableBundle {
	return &diode.CableBundle{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}

// CableBundleExplicit Creates a CableBundle with fully nested objects and all common fields.
func CableBundleExplicit() *diode.CableBundle {
	return &diode.CableBundle{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
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
