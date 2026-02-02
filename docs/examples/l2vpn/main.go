// Package main demonstrates ingesting L2VPN entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "l2vpn-example"
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
	l2vpn := L2VPNMinimal()
	// l2vpn := L2VPNExtended()
	// l2vpn := L2VPNExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{l2vpn})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("L2VPN ingested successfully")
	}
}

// L2VPNMinimal Creates a L2VPN with only required fields.
func L2VPNMinimal() *diode.L2VPN {
	return &diode.L2VPN{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// L2VPNExtended Creates a L2VPN with common optional fields.
func L2VPNExtended() *diode.L2VPN {
	return &diode.L2VPN{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
	}
}

// L2VPNExplicit Creates a L2VPN with fully nested objects and all common fields.
func L2VPNExplicit() *diode.L2VPN {
	return &diode.L2VPN{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
