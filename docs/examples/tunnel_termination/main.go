// Package main demonstrates ingesting TunnelTermination entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tunnel_termination-example"
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
	tunnelTermination := TunnelTerminationMinimal()
	// tunnelTermination := TunnelTerminationExtended()
	// tunnelTermination := TunnelTerminationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tunnelTermination})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("TunnelTermination ingested successfully")
	}
}

// TunnelTerminationMinimal Creates a TunnelTermination with only required fields.
func TunnelTerminationMinimal() *diode.TunnelTermination {
	return &diode.TunnelTermination{
		Tunnel: &diode.Tunnel{
			Name:          diode.String("Example Name"),
			Status:        diode.String("active"),
			Encapsulation: diode.String("Example Encapsulation"),
		},
		Role: diode.String("hub"),
	}
}

// TunnelTerminationExtended Creates a TunnelTermination with common optional fields.
func TunnelTerminationExtended() *diode.TunnelTermination {
	return &diode.TunnelTermination{
		Tunnel: &diode.Tunnel{
			Name:          diode.String("Example Name"),
			Status:        diode.String("active"),
			Encapsulation: diode.String("Example Encapsulation"),
		},
		Role: diode.String("hub"),
	}
}

// TunnelTerminationExplicit Creates a TunnelTermination with fully nested objects and all common fields.
func TunnelTerminationExplicit() *diode.TunnelTermination {
	return &diode.TunnelTermination{
		Tunnel: &diode.Tunnel{
			Name:          diode.String("Example Name"),
			Status:        diode.String("active"),
			Encapsulation: diode.String("Example Encapsulation"),
		},
		Role: diode.String("hub"),
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
