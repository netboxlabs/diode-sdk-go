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
		Ssid:     diode.String("ExampleSSID"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// WirelessLANExtended Creates a WirelessLAN with common optional fields.
func WirelessLANExtended() *diode.WirelessLAN {
	return &diode.WirelessLAN{
		Ssid:        diode.String("ExampleSSID"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		AuthType:    diode.String("open"),
		AuthCipher:  diode.String("aes"),
		AuthPsk:     diode.String("Example AuthPsk"),
		Comments:    diode.String("Example comments"),
	}
}

// WirelessLANExplicit Creates a WirelessLAN with fully nested objects and all common fields.
func WirelessLANExplicit() *diode.WirelessLAN {
	return &diode.WirelessLAN{
		Ssid:        diode.String("ExampleSSID"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		AuthType:    diode.String("open"),
		AuthCipher:  diode.String("aes"),
		AuthPsk:     diode.String("Example AuthPsk"),
		Group: &diode.WirelessLANGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Vlan: &diode.VLAN{
			Vid:      diode.Int64(1),
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
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
