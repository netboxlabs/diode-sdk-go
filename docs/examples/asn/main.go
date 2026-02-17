// Package main demonstrates ingesting ASN entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "asn-example"
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
	asn := ASNMinimal()
	// asn := ASNExtended()
	// asn := ASNExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{asn})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ASN ingested successfully")
	}
}

// ASNMinimal Creates a ASN with only required fields.
func ASNMinimal() *diode.ASN {
	return &diode.ASN{
		Asn:      diode.Int64(64512),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ASNExtended Creates a ASN with common optional fields.
func ASNExtended() *diode.ASN {
	return &diode.ASN{
		Asn:         diode.Int64(64512),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// ASNExplicit Creates a ASN with fully nested objects and all common fields.
func ASNExplicit() *diode.ASN {
	return &diode.ASN{
		Asn:         diode.Int64(64512),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
