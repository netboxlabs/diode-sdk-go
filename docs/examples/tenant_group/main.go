// Package main demonstrates ingesting TenantGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tenant_group-example"
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
	tenantGroup := TenantGroupMinimal()
	// tenantGroup := TenantGroupExtended()
	// tenantGroup := TenantGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tenantGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("TenantGroup ingested successfully")
	}
}

// TenantGroupMinimal Creates a TenantGroup with only required fields.
func TenantGroupMinimal() *diode.TenantGroup {
	return &diode.TenantGroup{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// TenantGroupExtended Creates a TenantGroup with common optional fields.
func TenantGroupExtended() *diode.TenantGroup {
	return &diode.TenantGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// TenantGroupExplicit Creates a TenantGroup with fully nested objects and all common fields.
func TenantGroupExplicit() *diode.TenantGroup {
	return &diode.TenantGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
