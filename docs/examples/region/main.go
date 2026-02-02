// Package main demonstrates ingesting Region entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "region-example"
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
	region := RegionMinimal()
	// region := RegionExtended()
	// region := RegionExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{region})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Region ingested successfully")
	}
}

// RegionMinimal Creates a Region with only required fields.
func RegionMinimal() *diode.Region {
	return &diode.Region{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// RegionExtended Creates a Region with common optional fields.
func RegionExtended() *diode.Region {
	return &diode.Region{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// RegionExplicit Creates a Region with fully nested objects and all common fields.
func RegionExplicit() *diode.Region {
	return &diode.Region{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
