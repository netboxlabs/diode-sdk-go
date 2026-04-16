// Package main demonstrates ingesting WirelessLANGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "wireless_lan_group-example"
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
	wirelessLanGroup := WirelessLANGroupMinimal()
	// wirelessLanGroup := WirelessLANGroupExtended()
	// wirelessLanGroup := WirelessLANGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{wirelessLanGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("WirelessLANGroup ingested successfully")
	}
}

// WirelessLANGroupMinimal Creates a WirelessLANGroup with only required fields.
func WirelessLANGroupMinimal() *diode.WirelessLANGroup {
	return &diode.WirelessLANGroup{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// WirelessLANGroupExtended Creates a WirelessLANGroup with common optional fields.
func WirelessLANGroupExtended() *diode.WirelessLANGroup {
	return &diode.WirelessLANGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}

// WirelessLANGroupExplicit Creates a WirelessLANGroup with fully nested objects and all common fields.
func WirelessLANGroupExplicit() *diode.WirelessLANGroup {
	return &diode.WirelessLANGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Parent: &diode.WirelessLANGroup{
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
