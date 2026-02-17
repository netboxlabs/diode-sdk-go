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
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Comments:    diode.String("Example comments"),
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
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Owner: &diode.Owner{
			Name: diode.String("Example Name"),
			Group: &diode.OwnerGroup{
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
