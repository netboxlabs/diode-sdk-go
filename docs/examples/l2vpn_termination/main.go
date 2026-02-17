// Package main demonstrates ingesting L2VPNTermination entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "l2vpn_termination-example"
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
	l2vpnTermination := L2VPNTerminationMinimal()
	// l2vpnTermination := L2VPNTerminationExtended()
	// l2vpnTermination := L2VPNTerminationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{l2vpnTermination})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("L2VPNTermination ingested successfully")
	}
}

// L2VPNTerminationMinimal Creates a L2VPNTermination with only required fields.
func L2VPNTerminationMinimal() *diode.L2VPNTermination {
	return &diode.L2VPNTermination{
		L2Vpn: &diode.L2VPN{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// L2VPNTerminationExtended Creates a L2VPNTermination with common optional fields.
func L2VPNTerminationExtended() *diode.L2VPNTermination {
	return &diode.L2VPNTermination{
		L2Vpn: &diode.L2VPN{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// L2VPNTerminationExplicit Creates a L2VPNTermination with fully nested objects and all common fields.
func L2VPNTerminationExplicit() *diode.L2VPNTermination {
	return &diode.L2VPNTermination{
		L2Vpn: &diode.L2VPN{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
		Tags:     []*diode.Tag{{Name: diode.String("production")}},
	}
}
