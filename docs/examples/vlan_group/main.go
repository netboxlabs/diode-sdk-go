// Package main demonstrates ingesting VLANGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vlan_group-example"
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
	vlanGroup := VLANGroupMinimal()
	// vlanGroup := VLANGroupExtended()
	// vlanGroup := VLANGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vlanGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VLANGroup ingested successfully")
	}
}

// VLANGroupMinimal Creates a VLANGroup with only required fields.
func VLANGroupMinimal() *diode.VLANGroup {
	return &diode.VLANGroup{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VLANGroupExtended Creates a VLANGroup with common optional fields.
func VLANGroupExtended() *diode.VLANGroup {
	return &diode.VLANGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// VLANGroupExplicit Creates a VLANGroup with fully nested objects and all common fields.
func VLANGroupExplicit() *diode.VLANGroup {
	return &diode.VLANGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
