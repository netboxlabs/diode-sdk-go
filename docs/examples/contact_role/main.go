// Package main demonstrates ingesting ContactRole entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "contact_role-example"
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
	contactRole := ContactRoleMinimal()
	// contactRole := ContactRoleExtended()
	// contactRole := ContactRoleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{contactRole})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ContactRole ingested successfully")
	}
}

// ContactRoleMinimal Creates a ContactRole with only required fields.
func ContactRoleMinimal() *diode.ContactRole {
	return &diode.ContactRole{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// ContactRoleExtended Creates a ContactRole with common optional fields.
func ContactRoleExtended() *diode.ContactRole {
	return &diode.ContactRole{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// ContactRoleExplicit Creates a ContactRole with fully nested objects and all common fields.
func ContactRoleExplicit() *diode.ContactRole {
	return &diode.ContactRole{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
