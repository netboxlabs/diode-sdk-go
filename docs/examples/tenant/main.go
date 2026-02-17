// Package main demonstrates ingesting Tenant entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tenant-example"
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
	tenant := TenantMinimal()
	// tenant := TenantExtended()
	// tenant := TenantExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tenant})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Tenant ingested successfully")
	}
}

// TenantMinimal Creates a Tenant with only required fields.
func TenantMinimal() *diode.Tenant {
	return &diode.Tenant{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// TenantExtended Creates a Tenant with common optional fields.
func TenantExtended() *diode.Tenant {
	return &diode.Tenant{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// TenantExplicit Creates a Tenant with fully nested objects and all common fields.
func TenantExplicit() *diode.Tenant {
	return &diode.Tenant{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
