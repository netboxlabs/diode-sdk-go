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
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value"},
	}
}

// CableTerminationExplicit Creates a CableTermination with fully nested objects and all common fields.
func CableTerminationExplicit() *diode.CableTermination {
	return &diode.CableTermination{
		Cable: &diode.Cable{
			Status:   diode.String("planned"),
			Color:    diode.String("0000ff"),
			Metadata: diode.Metadata{"source": "example"},
		},
		CableEnd: diode.String("A"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		// Polymorphic 'termination' — choose ONE variant for Termination:
		Termination: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name:     diode.String("Example Name"),
						Slug:     diode.String("example-slug"),
						Metadata: diode.Metadata{"source": "example"},
					},
					Model:    diode.String("Model X"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Role: &diode.DeviceRole{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Color:    diode.String("0000ff"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Site: &diode.Site{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Status:   diode.String("active"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Type:     diode.String("1000base-t"),
			Metadata: diode.Metadata{"source": "example"},
		},
		// Termination: &diode.CircuitTermination{ Circuit: &diode.Circuit{ Cid: diode.String("CID-001"), Provider: &diode.Provider{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Type: &diode.CircuitType{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, TermSide: diode.String("A"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.ConsolePort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.ConsoleServerPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.FrontPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Type: diode.String("110-punch"), Color: diode.String("0000ff"), RearPort: &diode.RearPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Type: diode.String("110-punch"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.PowerFeed{ PowerPanel: &diode.PowerPanel{ Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.PowerOutlet{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Color: diode.String("0000ff"), Status: diode.String("disabled"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.PowerPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.RearPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Type: diode.String("110-punch"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, },
	}
}
