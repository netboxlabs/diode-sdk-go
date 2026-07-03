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
			Encapsulation: diode.String("gre"),
			Metadata:      diode.Metadata{"source": "example"},
		},
		Role:     diode.String("hub"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// TunnelTerminationExtended Creates a TunnelTermination with common optional fields.
func TunnelTerminationExtended() *diode.TunnelTermination {
	return &diode.TunnelTermination{
		Tunnel: &diode.Tunnel{
			Name:          diode.String("Example Name"),
			Status:        diode.String("active"),
			Encapsulation: diode.String("gre"),
			Metadata:      diode.Metadata{"source": "example"},
		},
		Role:     diode.String("hub"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value"},
	}
}

// TunnelTerminationExplicit Creates a TunnelTermination with fully nested objects and all common fields.
func TunnelTerminationExplicit() *diode.TunnelTermination {
	return &diode.TunnelTermination{
		Tunnel: &diode.Tunnel{
			Name:          diode.String("Example Name"),
			Status:        diode.String("active"),
			Encapsulation: diode.String("gre"),
			Metadata:      diode.Metadata{"source": "example"},
		},
		Role:     diode.String("hub"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		OutsideIp: &diode.IPAddress{
			Address:  diode.String("192.0.2.1/32"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
