// Package main demonstrates ingesting RackReservation entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "rack_reservation-example"
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
	rackReservation := RackReservationMinimal()
	// rackReservation := RackReservationExtended()
	// rackReservation := RackReservationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{rackReservation})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("RackReservation ingested successfully")
	}
}

// RackReservationMinimal Creates a RackReservation with only required fields.
func RackReservationMinimal() *diode.RackReservation {
	return &diode.RackReservation{
		Rack: &diode.Rack{
			Name: diode.String("Example Name"),
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Description: diode.String("Example description"),
		Metadata:    diode.Metadata{"source": "example"},
	}
}

// RackReservationExtended Creates a RackReservation with common optional fields.
func RackReservationExtended() *diode.RackReservation {
	return &diode.RackReservation{
		Rack: &diode.Rack{
			Name: diode.String("Example Name"),
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Description: diode.String("Example description"),
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
	}
}

// RackReservationExplicit Creates a RackReservation with fully nested objects and all common fields.
func RackReservationExplicit() *diode.RackReservation {
	return &diode.RackReservation{
		Rack: &diode.Rack{
			Name: diode.String("Example Name"),
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Description: diode.String("Example description"),
		Metadata:    diode.Metadata{"source": "example"},
		Comments:    diode.String("Example comments"),
		Status:      diode.String("active"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
