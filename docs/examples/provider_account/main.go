// Package main demonstrates ingesting ProviderAccount entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "provider_account-example"
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
	providerAccount := ProviderAccountMinimal()
	// providerAccount := ProviderAccountExtended()
	// providerAccount := ProviderAccountExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{providerAccount})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ProviderAccount ingested successfully")
	}
}

// ProviderAccountMinimal Creates a ProviderAccount with only required fields.
func ProviderAccountMinimal() *diode.ProviderAccount {
	return &diode.ProviderAccount{
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Account:  diode.String("Example Account"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ProviderAccountExtended Creates a ProviderAccount with common optional fields.
func ProviderAccountExtended() *diode.ProviderAccount {
	return &diode.ProviderAccount{
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Account:     diode.String("Example Account"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// ProviderAccountExplicit Creates a ProviderAccount with fully nested objects and all common fields.
func ProviderAccountExplicit() *diode.ProviderAccount {
	return &diode.ProviderAccount{
		Provider: &diode.Provider{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Account:     diode.String("Example Account"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
