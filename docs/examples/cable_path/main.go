// Package main demonstrates ingesting CablePath entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cable_path-example"
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
	cablePath := CablePathMinimal()
	// cablePath := CablePathExtended()
	// cablePath := CablePathExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{cablePath})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CablePath ingested successfully")
	}
}

// CablePathMinimal Creates a CablePath with only required fields.
func CablePathMinimal() *diode.CablePath {
	return &diode.CablePath{}
}

// CablePathExtended Creates a CablePath with common optional fields.
func CablePathExtended() *diode.CablePath {
	return &diode.CablePath{}
}

// CablePathExplicit Creates a CablePath with fully nested objects and all common fields.
func CablePathExplicit() *diode.CablePath {
	return &diode.CablePath{}
}
