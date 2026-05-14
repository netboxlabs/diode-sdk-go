// Package main demonstrates ingesting VirtualMachineType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_machine_type-example"
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
	virtualMachineType := VirtualMachineTypeMinimal()
	// virtualMachineType := VirtualMachineTypeExtended()
	// virtualMachineType := VirtualMachineTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualMachineType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualMachineType ingested successfully")
	}
}

// VirtualMachineTypeMinimal Creates a VirtualMachineType with only required fields.
func VirtualMachineTypeMinimal() *diode.VirtualMachineType {
	return &diode.VirtualMachineType{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VirtualMachineTypeExtended Creates a VirtualMachineType with common optional fields.
func VirtualMachineTypeExtended() *diode.VirtualMachineType {
	return &diode.VirtualMachineType{
		Name:          diode.String("Example Name"),
		Slug:          diode.String("example-slug"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		DefaultVcpus:  diode.Float64(1.0),
		DefaultMemory: diode.Int64(1),
		Comments:      diode.String("Example comments"),
	}
}

// VirtualMachineTypeExplicit Creates a VirtualMachineType with fully nested objects and all common fields.
func VirtualMachineTypeExplicit() *diode.VirtualMachineType {
	return &diode.VirtualMachineType{
		Name:          diode.String("Example Name"),
		Slug:          diode.String("example-slug"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		DefaultVcpus:  diode.Float64(1.0),
		DefaultMemory: diode.Int64(1),
		DefaultPlatform: &diode.Platform{
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
