// Package main demonstrates ingesting CableTermination entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cable_termination-example"
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
	cableTermination := CableTerminationMinimal()
	// cableTermination := CableTerminationExtended()
	// cableTermination := CableTerminationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{cableTermination})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CableTermination ingested successfully")
	}
}

// CableTerminationMinimal Creates a CableTermination with only required fields.
func CableTerminationMinimal() *diode.CableTermination {
	return &diode.CableTermination{
		Cable: &diode.Cable{
			Metadata: diode.Metadata{"source": "example"},
		},
		CableEnd: diode.String("A"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CableTerminationExtended Creates a CableTermination with common optional fields.
func CableTerminationExtended() *diode.CableTermination {
	return &diode.CableTermination{
		Cable: &diode.Cable{
			Metadata: diode.Metadata{"source": "example"},
		},
		CableEnd: diode.String("A"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CableTerminationExplicit Creates a CableTermination with fully nested objects and all common fields.
func CableTerminationExplicit() *diode.CableTermination {
	return &diode.CableTermination{
		Cable: &diode.Cable{
			Status:   diode.String("active"),
			Color:    diode.String("0000ff"),
			Metadata: diode.Metadata{"source": "example"},
		},
		CableEnd: diode.String("A"),
		Metadata: diode.Metadata{"source": "example"},
	}
}
