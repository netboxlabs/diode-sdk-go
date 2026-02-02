// Package main demonstrates ingesting FHRPGroupAssignment entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "fhrp_group_assignment-example"
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
	fhrpGroupAssignment := FHRPGroupAssignmentMinimal()
	// fhrpGroupAssignment := FHRPGroupAssignmentExtended()
	// fhrpGroupAssignment := FHRPGroupAssignmentExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{fhrpGroupAssignment})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("FHRPGroupAssignment ingested successfully")
	}
}

// FHRPGroupAssignmentMinimal Creates a FHRPGroupAssignment with only required fields.
func FHRPGroupAssignmentMinimal() *diode.FHRPGroupAssignment {
	return &diode.FHRPGroupAssignment{
		Group: &diode.FHRPGroup{
			Protocol: diode.String("Example Protocol"),
			GroupId:  diode.Int64(1),
		},
		Priority: diode.Int64(1),
	}
}

// FHRPGroupAssignmentExtended Creates a FHRPGroupAssignment with common optional fields.
func FHRPGroupAssignmentExtended() *diode.FHRPGroupAssignment {
	return &diode.FHRPGroupAssignment{
		Group: &diode.FHRPGroup{
			Protocol: diode.String("Example Protocol"),
			GroupId:  diode.Int64(1),
		},
		Priority: diode.Int64(1),
	}
}

// FHRPGroupAssignmentExplicit Creates a FHRPGroupAssignment with fully nested objects and all common fields.
func FHRPGroupAssignmentExplicit() *diode.FHRPGroupAssignment {
	return &diode.FHRPGroupAssignment{
		Group: &diode.FHRPGroup{
			Protocol: diode.String("Example Protocol"),
			GroupId:  diode.Int64(1),
		},
		Priority: diode.Int64(1),
	}
}
