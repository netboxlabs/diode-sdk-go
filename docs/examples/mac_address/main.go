// Package main demonstrates ingesting MACAddress entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "mac_address-example"
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
	macAddress := MACAddressMinimal()
	// macAddress := MACAddressExtended()
	// macAddress := MACAddressExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{macAddress})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("MACAddress ingested successfully")
	}
}

// MACAddressMinimal Creates a MACAddress with only required fields.
func MACAddressMinimal() *diode.MACAddress {
	return &diode.MACAddress{
		MacAddress: diode.String("00:11:22:33:44:55"),
		Metadata:   diode.Metadata{"source": "example"},
	}
}

// MACAddressExtended Creates a MACAddress with common optional fields.
func MACAddressExtended() *diode.MACAddress {
	return &diode.MACAddress{
		MacAddress:  diode.String("00:11:22:33:44:55"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}

// MACAddressExplicit Creates a MACAddress with fully nested objects and all common fields.
func MACAddressExplicit() *diode.MACAddress {
	return &diode.MACAddress{
		MacAddress:  diode.String("00:11:22:33:44:55"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Owner: &diode.Owner{
			Name: diode.String("Example Name"),
			Group: &diode.OwnerGroup{
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		// Polymorphic 'assigned_object' — choose ONE variant for AssignedObject:
		AssignedObject: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name:     diode.String("Example Name"),
						Slug:     diode.String("example-slug"),
						Metadata: diode.Metadata{"source": "example"},
					},
					Model:    diode.String("Model X"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Role: &diode.DeviceRole{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Color:    diode.String("0000ff"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Site: &diode.Site{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Status:   diode.String("active"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Type:     diode.String("1000base-t"),
			Metadata: diode.Metadata{"source": "example"},
		},
		// AssignedObject: &diode.VMInterface{ VirtualMachine: &diode.VirtualMachine{ Name: diode.String("Example Name"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
