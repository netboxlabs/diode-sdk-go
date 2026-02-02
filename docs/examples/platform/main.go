// Package main demonstrates ingesting Platform entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "platform-example"
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
	platform := PlatformMinimal()
	// platform := PlatformExtended()
	// platform := PlatformExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{platform})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Platform ingested successfully")
	}
}

// PlatformMinimal Creates a Platform with only required fields.
func PlatformMinimal() *diode.Platform {
	return &diode.Platform{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// PlatformExtended Creates a Platform with common optional fields.
func PlatformExtended() *diode.Platform {
	return &diode.Platform{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// PlatformExplicit Creates a Platform with fully nested objects and all common fields.
func PlatformExplicit() *diode.Platform {
	return &diode.Platform{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
