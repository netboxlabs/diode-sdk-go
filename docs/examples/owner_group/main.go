// Package main demonstrates ingesting OwnerGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "owner_group-example"
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
	ownerGroup := OwnerGroupMinimal()
	// ownerGroup := OwnerGroupExtended()
	// ownerGroup := OwnerGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ownerGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("OwnerGroup ingested successfully")
	}
}

// OwnerGroupMinimal Creates a OwnerGroup with only required fields.
func OwnerGroupMinimal() *diode.OwnerGroup {
	return &diode.OwnerGroup{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// OwnerGroupExtended Creates a OwnerGroup with common optional fields.
func OwnerGroupExtended() *diode.OwnerGroup {
	return &diode.OwnerGroup{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
	}
}

// OwnerGroupExplicit Creates a OwnerGroup with fully nested objects and all common fields.
func OwnerGroupExplicit() *diode.OwnerGroup {
	return &diode.OwnerGroup{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
	}
}
