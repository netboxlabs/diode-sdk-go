// Package main demonstrates ingesting Provider entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "provider-example"
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
	provider := ProviderMinimal()
	// provider := ProviderExtended()
	// provider := ProviderExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{provider})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Provider ingested successfully")
	}
}

// ProviderMinimal Creates a Provider with only required fields.
func ProviderMinimal() *diode.Provider {
	return &diode.Provider{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// ProviderExtended Creates a Provider with common optional fields.
func ProviderExtended() *diode.Provider {
	return &diode.Provider{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// ProviderExplicit Creates a Provider with fully nested objects and all common fields.
func ProviderExplicit() *diode.Provider {
	return &diode.Provider{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
