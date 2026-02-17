// Package main demonstrates ingesting Contact entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "contact-example"
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
	contact := ContactMinimal()
	// contact := ContactExtended()
	// contact := ContactExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{contact})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Contact ingested successfully")
	}
}

// ContactMinimal Creates a Contact with only required fields.
func ContactMinimal() *diode.Contact {
	return &diode.Contact{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ContactExtended Creates a Contact with common optional fields.
func ContactExtended() *diode.Contact {
	return &diode.Contact{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// ContactExplicit Creates a Contact with fully nested objects and all common fields.
func ContactExplicit() *diode.Contact {
	return &diode.Contact{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
