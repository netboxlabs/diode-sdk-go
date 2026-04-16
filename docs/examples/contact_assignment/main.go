// Package main demonstrates ingesting ContactAssignment entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "contact_assignment-example"
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
	contactAssignment := ContactAssignmentMinimal()
	// contactAssignment := ContactAssignmentExtended()
	// contactAssignment := ContactAssignmentExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{contactAssignment})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ContactAssignment ingested successfully")
	}
}

// ContactAssignmentMinimal Creates a ContactAssignment with only required fields.
func ContactAssignmentMinimal() *diode.ContactAssignment {
	return &diode.ContactAssignment{
		Contact: &diode.Contact{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ContactAssignmentExtended Creates a ContactAssignment with common optional fields.
func ContactAssignmentExtended() *diode.ContactAssignment {
	return &diode.ContactAssignment{
		Contact: &diode.Contact{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Priority: diode.String("inactive"),
	}
}

// ContactAssignmentExplicit Creates a ContactAssignment with fully nested objects and all common fields.
func ContactAssignmentExplicit() *diode.ContactAssignment {
	return &diode.ContactAssignment{
		Contact: &diode.Contact{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Priority: diode.String("inactive"),
		Role: &diode.ContactRole{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
