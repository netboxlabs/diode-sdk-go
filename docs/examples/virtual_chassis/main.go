// Package main demonstrates ingesting VirtualChassis entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_chassis-example"
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
	virtualChassis := VirtualChassisMinimal()
	// virtualChassis := VirtualChassisExtended()
	// virtualChassis := VirtualChassisExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualChassis})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualChassis ingested successfully")
	}
}

// VirtualChassisMinimal Creates a VirtualChassis with only required fields.
func VirtualChassisMinimal() *diode.VirtualChassis {
	return &diode.VirtualChassis{
		Name: diode.String("Example Name"),
	}
}

// VirtualChassisExtended Creates a VirtualChassis with common optional fields.
func VirtualChassisExtended() *diode.VirtualChassis {
	return &diode.VirtualChassis{
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
	}
}

// VirtualChassisExplicit Creates a VirtualChassis with fully nested objects and all common fields.
func VirtualChassisExplicit() *diode.VirtualChassis {
	return &diode.VirtualChassis{
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
