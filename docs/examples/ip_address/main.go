// Package main demonstrates ingesting IPAddress entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ip_address-example"
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
	ipAddress := IPAddressMinimal()
	// ipAddress := IPAddressExtended()
	// ipAddress := IPAddressExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ipAddress})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IPAddress ingested successfully")
	}
}

// IPAddressMinimal Creates a IPAddress with only required fields.
func IPAddressMinimal() *diode.IPAddress {
	return &diode.IPAddress{
		Address:  diode.String("192.0.2.1/32"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// IPAddressExtended Creates a IPAddress with common optional fields.
func IPAddressExtended() *diode.IPAddress {
	return &diode.IPAddress{
		Address:     diode.String("192.0.2.1/32"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Role:        diode.String("anycast"),
		DnsName:     diode.String("Example DnsName"),
		Comments:    diode.String("Example comments"),
	}
}

// IPAddressExplicit Creates a IPAddress with fully nested objects and all common fields.
func IPAddressExplicit() *diode.IPAddress {
	return &diode.IPAddress{
		Address:     diode.String("192.0.2.1/32"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Role:        diode.String("anycast"),
		DnsName:     diode.String("Example DnsName"),
		Vrf: &diode.VRF{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		NatInside: &diode.IPAddress{
			Address:  diode.String("192.0.2.1/32"),
			Status:   diode.String("active"),
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
