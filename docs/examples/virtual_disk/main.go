// Package main demonstrates ingesting VirtualDisk entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_disk-example"
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
	virtualDisk := VirtualDiskMinimal()
	// virtualDisk := VirtualDiskExtended()
	// virtualDisk := VirtualDiskExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualDisk})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualDisk ingested successfully")
	}
}

// VirtualDiskMinimal Creates a VirtualDisk with only required fields.
func VirtualDiskMinimal() *diode.VirtualDisk {
	return &diode.VirtualDisk{
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Size:     diode.Int64(1),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VirtualDiskExtended Creates a VirtualDisk with common optional fields.
func VirtualDiskExtended() *diode.VirtualDisk {
	return &diode.VirtualDisk{
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Size:        diode.Int64(1),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
	}
}

// VirtualDiskExplicit Creates a VirtualDisk with fully nested objects and all common fields.
func VirtualDiskExplicit() *diode.VirtualDisk {
	return &diode.VirtualDisk{
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Size:        diode.Int64(1),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
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
