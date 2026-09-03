// Package main demonstrates ingesting VMInterface entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vm_interface-example"
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
	vmInterface := VMInterfaceMinimal()
	// vmInterface := VMInterfaceExtended()
	// vmInterface := VMInterfaceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vmInterface})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VMInterface ingested successfully")
	}
}

// VMInterfaceMinimal Creates a VMInterface with only required fields.
func VMInterfaceMinimal() *diode.VMInterface {
	return &diode.VMInterface{
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VMInterfaceExtended Creates a VMInterface with common optional fields.
func VMInterfaceExtended() *diode.VMInterface {
	return &diode.VMInterface{
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Enabled:     diode.Bool(true),
		Mtu:         diode.Int64(1),
		Mode:        diode.String("access"),
		MacAddress:  diode.String("00:11:22:33:44:55"),
	}
}

// VMInterfaceExplicit Creates a VMInterface with fully nested objects and all common fields.
func VMInterfaceExplicit() *diode.VMInterface {
	return &diode.VMInterface{
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Enabled:     diode.Bool(true),
		Mtu:         diode.Int64(1),
		Mode:        diode.String("access"),
		MacAddress:  diode.String("00:11:22:33:44:55"),
		Parent: &diode.VMInterface{
			VirtualMachine: &diode.VirtualMachine{
				Name:     diode.String("Example Name"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Bridge: &diode.VMInterface{
			VirtualMachine: &diode.VirtualMachine{
				Name:     diode.String("Example Name"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		PrimaryMacAddress: &diode.MACAddress{
			MacAddress: diode.String("00:11:22:33:44:55"),
			Metadata:   diode.Metadata{"source": "example"},
		},
		UntaggedVlan: &diode.VLAN{
			Vid:      diode.Int64(1),
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		QinqSvlan: &diode.VLAN{
			Vid:      diode.Int64(1),
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		VlanTranslationPolicy: &diode.VLANTranslationPolicy{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Vrf: &diode.VRF{
			Name:     diode.String("Example Name"),
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
