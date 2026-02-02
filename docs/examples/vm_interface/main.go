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
			Name: diode.String("Example Name"),
		},
		Name: diode.String("Example Name"),
	}
}

// VMInterfaceExtended Creates a VMInterface with common optional fields.
func VMInterfaceExtended() *diode.VMInterface {
	return &diode.VMInterface{
		VirtualMachine: &diode.VirtualMachine{
			Name: diode.String("Example Name"),
		},
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
	}
}

// VMInterfaceExplicit Creates a VMInterface with fully nested objects and all common fields.
func VMInterfaceExplicit() *diode.VMInterface {
	return &diode.VMInterface{
		VirtualMachine: &diode.VirtualMachine{
			Name:   diode.String("Example Name"),
			Status: diode.String("active"),
		},
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
