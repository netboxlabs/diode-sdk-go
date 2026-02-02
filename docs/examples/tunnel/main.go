// Package main demonstrates ingesting Tunnel entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tunnel-example"
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
	tunnel := TunnelMinimal()
	// tunnel := TunnelExtended()
	// tunnel := TunnelExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tunnel})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Tunnel ingested successfully")
	}
}

// TunnelMinimal Creates a Tunnel with only required fields.
func TunnelMinimal() *diode.Tunnel {
	return &diode.Tunnel{
		Name:          diode.String("Example Name"),
		Status:        diode.String("active"),
		Encapsulation: diode.String("gre"),
	}
}

// TunnelExtended Creates a Tunnel with common optional fields.
func TunnelExtended() *diode.Tunnel {
	return &diode.Tunnel{
		Name:          diode.String("Example Name"),
		Status:        diode.String("active"),
		Encapsulation: diode.String("gre"),
		Description:   diode.String("Example description"),
	}
}

// TunnelExplicit Creates a Tunnel with fully nested objects and all common fields.
func TunnelExplicit() *diode.Tunnel {
	return &diode.Tunnel{
		Name:          diode.String("Example Name"),
		Status:        diode.String("active"),
		Encapsulation: diode.String("gre"),
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
