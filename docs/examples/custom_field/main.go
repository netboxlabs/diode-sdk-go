// Package main demonstrates ingesting CustomField entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "custom_field-example"
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
	customField := CustomFieldMinimal()
	// customField := CustomFieldExtended()
	// customField := CustomFieldExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{customField})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CustomField ingested successfully")
	}
}

// CustomFieldMinimal Creates a CustomField with only required fields.
func CustomFieldMinimal() *diode.CustomField {
	return &diode.CustomField{
		Type:     diode.String("boolean"),
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CustomFieldExtended Creates a CustomField with common optional fields.
func CustomFieldExtended() *diode.CustomField {
	return &diode.CustomField{
		Type:        diode.String("boolean"),
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// CustomFieldExplicit Creates a CustomField with fully nested objects and all common fields.
func CustomFieldExplicit() *diode.CustomField {
	return &diode.CustomField{
		Type:        diode.String("boolean"),
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}
