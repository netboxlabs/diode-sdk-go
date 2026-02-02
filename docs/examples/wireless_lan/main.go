// Package main demonstrates ingesting WirelessLAN entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "wireless_lan-example"
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
	wirelessLan := WirelessLANMinimal()
	// wirelessLan := WirelessLANExtended()
	// wirelessLan := WirelessLANExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{wirelessLan})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("WirelessLAN ingested successfully")
	}
}

// WirelessLANMinimal Creates a WirelessLAN with only required fields.
func WirelessLANMinimal() *diode.WirelessLAN {
	return &diode.WirelessLAN{
		Ssid: diode.String("ExampleSSID"),
	}
}

// WirelessLANExtended Creates a WirelessLAN with common optional fields.
func WirelessLANExtended() *diode.WirelessLAN {
	return &diode.WirelessLAN{
		Ssid:        diode.String("ExampleSSID"),
		Description: diode.String("Example description"),
		Status:      diode.String("active"),
	}
}

// WirelessLANExplicit Creates a WirelessLAN with fully nested objects and all common fields.
func WirelessLANExplicit() *diode.WirelessLAN {
	return &diode.WirelessLAN{
		Ssid:        diode.String("ExampleSSID"),
		Description: diode.String("Example description"),
		Status:      diode.String("active"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
