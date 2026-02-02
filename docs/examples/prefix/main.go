// Package main demonstrates ingesting Prefix entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "prefix-example"
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
	prefix := PrefixMinimal()
	// prefix := PrefixExtended()
	// prefix := PrefixExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{prefix})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Prefix ingested successfully")
	}
}

// PrefixMinimal Creates a Prefix with only required fields.
func PrefixMinimal() *diode.Prefix {
	return &diode.Prefix{
		Prefix: diode.String("192.0.2.0/24"),
	}
}

// PrefixExtended Creates a Prefix with common optional fields.
func PrefixExtended() *diode.Prefix {
	return &diode.Prefix{
		Prefix:      diode.String("192.0.2.0/24"),
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
	}
}

// PrefixExplicit Creates a Prefix with fully nested objects and all common fields.
func PrefixExplicit() *diode.Prefix {
	return &diode.Prefix{
		Prefix:      diode.String("192.0.2.0/24"),
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
