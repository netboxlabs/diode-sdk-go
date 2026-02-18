// Package main demonstrates ingesting VRF entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vrf-example"
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
	vrf := VRFMinimal()
	// vrf := VRFExtended()
	// vrf := VRFExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vrf})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VRF ingested successfully")
	}
}

// VRFMinimal Creates a VRF with only required fields.
func VRFMinimal() *diode.VRF {
	return &diode.VRF{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VRFExtended Creates a VRF with common optional fields.
func VRFExtended() *diode.VRF {
	return &diode.VRF{
		Name:          diode.String("Example Name"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		Rd:            diode.String("Example Rd"),
		EnforceUnique: diode.Bool(true),
		Comments:      diode.String("Example comments"),
	}
}

// VRFExplicit Creates a VRF with fully nested objects and all common fields.
func VRFExplicit() *diode.VRF {
	return &diode.VRF{
		Name:          diode.String("Example Name"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		Rd:            diode.String("Example Rd"),
		EnforceUnique: diode.Bool(true),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
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
