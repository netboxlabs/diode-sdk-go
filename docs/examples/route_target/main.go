// Package main demonstrates ingesting RouteTarget entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "route_target-example"
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
	routeTarget := RouteTargetMinimal()
	// routeTarget := RouteTargetExtended()
	// routeTarget := RouteTargetExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{routeTarget})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("RouteTarget ingested successfully")
	}
}

// RouteTargetMinimal Creates a RouteTarget with only required fields.
func RouteTargetMinimal() *diode.RouteTarget {
	return &diode.RouteTarget{
		Name: diode.String("Example Name"),
	}
}

// RouteTargetExtended Creates a RouteTarget with common optional fields.
func RouteTargetExtended() *diode.RouteTarget {
	return &diode.RouteTarget{
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
	}
}

// RouteTargetExplicit Creates a RouteTarget with fully nested objects and all common fields.
func RouteTargetExplicit() *diode.RouteTarget {
	return &diode.RouteTarget{
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
