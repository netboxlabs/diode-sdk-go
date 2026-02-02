// Package main demonstrates ingesting ContactGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "contact_group-example"
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
	contactGroup := ContactGroupMinimal()
	// contactGroup := ContactGroupExtended()
	// contactGroup := ContactGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{contactGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ContactGroup ingested successfully")
	}
}

// ContactGroupMinimal Creates a ContactGroup with only required fields.
func ContactGroupMinimal() *diode.ContactGroup {
	return &diode.ContactGroup{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// ContactGroupExtended Creates a ContactGroup with common optional fields.
func ContactGroupExtended() *diode.ContactGroup {
	return &diode.ContactGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// ContactGroupExplicit Creates a ContactGroup with fully nested objects and all common fields.
func ContactGroupExplicit() *diode.ContactGroup {
	return &diode.ContactGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
