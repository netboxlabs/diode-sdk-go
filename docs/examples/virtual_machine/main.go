// Package main demonstrates ingesting VirtualMachine entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_machine-example"
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
	virtualMachine := VirtualMachineMinimal()
	// virtualMachine := VirtualMachineExtended()
	// virtualMachine := VirtualMachineExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualMachine})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualMachine ingested successfully")
	}
}

// VirtualMachineMinimal Creates a VirtualMachine with only required fields.
func VirtualMachineMinimal() *diode.VirtualMachine {
	return &diode.VirtualMachine{
		Name: diode.String("Example Name"),
	}
}

// VirtualMachineExtended Creates a VirtualMachine with common optional fields.
func VirtualMachineExtended() *diode.VirtualMachine {
	return &diode.VirtualMachine{
		Name:        diode.String("Example Name"),
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
	}
}

// VirtualMachineExplicit Creates a VirtualMachine with fully nested objects and all common fields.
func VirtualMachineExplicit() *diode.VirtualMachine {
	return &diode.VirtualMachine{
		Name:        diode.String("Example Name"),
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Platform: &diode.Platform{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
