// Package main demonstrates ingesting CircuitGroupAssignment entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "circuit_group_assignment-example"
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
	circuitGroupAssignment := CircuitGroupAssignmentMinimal()
	// circuitGroupAssignment := CircuitGroupAssignmentExtended()
	// circuitGroupAssignment := CircuitGroupAssignmentExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{circuitGroupAssignment})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CircuitGroupAssignment ingested successfully")
	}
}

// CircuitGroupAssignmentMinimal Creates a CircuitGroupAssignment with only required fields.
func CircuitGroupAssignmentMinimal() *diode.CircuitGroupAssignment {
	return &diode.CircuitGroupAssignment{
		Group: &diode.CircuitGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CircuitGroupAssignmentExtended Creates a CircuitGroupAssignment with common optional fields.
func CircuitGroupAssignmentExtended() *diode.CircuitGroupAssignment {
	return &diode.CircuitGroupAssignment{
		Group: &diode.CircuitGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Priority: diode.String("inactive"),
	}
}

// CircuitGroupAssignmentExplicit Creates a CircuitGroupAssignment with fully nested objects and all common fields.
func CircuitGroupAssignmentExplicit() *diode.CircuitGroupAssignment {
	return &diode.CircuitGroupAssignment{
		Group: &diode.CircuitGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Priority: diode.String("inactive"),
		Tags:     []*diode.Tag{{Name: diode.String("production")}},
	}
}
