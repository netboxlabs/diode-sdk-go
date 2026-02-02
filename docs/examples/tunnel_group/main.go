// Package main demonstrates ingesting TunnelGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tunnel_group-example"
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
	tunnelGroup := TunnelGroupMinimal()
	// tunnelGroup := TunnelGroupExtended()
	// tunnelGroup := TunnelGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tunnelGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("TunnelGroup ingested successfully")
	}
}

// TunnelGroupMinimal Creates a TunnelGroup with only required fields.
func TunnelGroupMinimal() *diode.TunnelGroup {
	return &diode.TunnelGroup{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// TunnelGroupExtended Creates a TunnelGroup with common optional fields.
func TunnelGroupExtended() *diode.TunnelGroup {
	return &diode.TunnelGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// TunnelGroupExplicit Creates a TunnelGroup with fully nested objects and all common fields.
func TunnelGroupExplicit() *diode.TunnelGroup {
	return &diode.TunnelGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
