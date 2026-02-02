// Package main demonstrates ingesting RearPort entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "rear_port-example"
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
	rearPort := RearPortMinimal()
	// rearPort := RearPortExtended()
	// rearPort := RearPortExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{rearPort})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("RearPort ingested successfully")
	}
}

// RearPortMinimal Creates a RearPort with only required fields.
func RearPortMinimal() *diode.RearPort {
	return &diode.RearPort{
		Device: &diode.Device{
			DeviceType: &diode.DeviceType{
				Manufacturer: &diode.Manufacturer{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Model: diode.String("Model X"),
				Slug:  diode.String("example-slug"),
			},
			Role: &diode.DeviceRole{
				Name: diode.String("Example Name"),
				Slug: diode.String("example-slug"),
			},
			Site: &diode.Site{
				Name: diode.String("Example Name"),
				Slug: diode.String("example-slug"),
			},
		},
		Name: diode.String("Example Name"),
		Type: diode.String("110-punch"),
	}
}

// RearPortExtended Creates a RearPort with common optional fields.
func RearPortExtended() *diode.RearPort {
	return &diode.RearPort{
		Device: &diode.Device{
			DeviceType: &diode.DeviceType{
				Manufacturer: &diode.Manufacturer{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Model: diode.String("Model X"),
				Slug:  diode.String("example-slug"),
			},
			Role: &diode.DeviceRole{
				Name: diode.String("Example Name"),
				Slug: diode.String("example-slug"),
			},
			Site: &diode.Site{
				Name: diode.String("Example Name"),
				Slug: diode.String("example-slug"),
			},
		},
		Name:        diode.String("Example Name"),
		Type:        diode.String("110-punch"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}

// RearPortExplicit Creates a RearPort with fully nested objects and all common fields.
func RearPortExplicit() *diode.RearPort {
	return &diode.RearPort{
		Device: &diode.Device{
			DeviceType: &diode.DeviceType{
				Manufacturer: &diode.Manufacturer{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Model: diode.String("Model X"),
				Slug:  diode.String("example-slug"),
			},
			Role: &diode.DeviceRole{
				Name:  diode.String("Example Name"),
				Slug:  diode.String("example-slug"),
				Color: diode.String("0000ff"),
			},
			Site: &diode.Site{
				Name:   diode.String("Example Name"),
				Slug:   diode.String("example-slug"),
				Status: diode.String("active"),
			},
			Status: diode.String("active"),
		},
		Name:        diode.String("Example Name"),
		Type:        diode.String("110-punch"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
