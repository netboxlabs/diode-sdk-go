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
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Name:        diode.String("Example Name"),
		Comments:    diode.String("Example comments"),
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
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Name:        diode.String("Example Name"),
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
