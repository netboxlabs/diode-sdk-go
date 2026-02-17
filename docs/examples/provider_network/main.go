// Package main demonstrates ingesting ProviderNetwork entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "provider_network-example"
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
	providerNetwork := ProviderNetworkMinimal()
	// providerNetwork := ProviderNetworkExtended()
	// providerNetwork := ProviderNetworkExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{providerNetwork})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ProviderNetwork ingested successfully")
	}
}

// ProviderNetworkMinimal Creates a ProviderNetwork with only required fields.
func ProviderNetworkMinimal() *diode.ProviderNetwork {
	return &diode.ProviderNetwork{
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ProviderNetworkExtended Creates a ProviderNetwork with common optional fields.
func ProviderNetworkExtended() *diode.ProviderNetwork {
	return &diode.ProviderNetwork{
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// ProviderNetworkExplicit Creates a ProviderNetwork with fully nested objects and all common fields.
func ProviderNetworkExplicit() *diode.ProviderNetwork {
	return &diode.ProviderNetwork{
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
