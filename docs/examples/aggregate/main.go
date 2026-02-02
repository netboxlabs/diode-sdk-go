// Package main demonstrates ingesting Aggregate entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "aggregate-example"
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
	aggregate := AggregateMinimal()
	// aggregate := AggregateExtended()
	// aggregate := AggregateExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{aggregate})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Aggregate ingested successfully")
	}
}

// AggregateMinimal Creates a Aggregate with only required fields.
func AggregateMinimal() *diode.Aggregate {
	return &diode.Aggregate{
		Prefix: diode.String("192.0.2.0/24"),
		Rir: &diode.RIR{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
	}
}

// AggregateExtended Creates a Aggregate with common optional fields.
func AggregateExtended() *diode.Aggregate {
	return &diode.Aggregate{
		Prefix: diode.String("192.0.2.0/24"),
		Rir: &diode.RIR{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Description: diode.String("Example description"),
	}
}

// AggregateExplicit Creates a Aggregate with fully nested objects and all common fields.
func AggregateExplicit() *diode.Aggregate {
	return &diode.Aggregate{
		Prefix: diode.String("192.0.2.0/24"),
		Rir: &diode.RIR{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
