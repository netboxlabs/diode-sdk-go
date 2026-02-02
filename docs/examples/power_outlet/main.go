// Package main demonstrates ingesting PowerOutlet entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "power_outlet-example"
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
	powerOutlet := PowerOutletMinimal()
	// powerOutlet := PowerOutletExtended()
	// powerOutlet := PowerOutletExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{powerOutlet})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("PowerOutlet ingested successfully")
	}
}

// PowerOutletMinimal Creates a PowerOutlet with only required fields.
func PowerOutletMinimal() *diode.PowerOutlet {
	return &diode.PowerOutlet{
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
	}
}

// PowerOutletExtended Creates a PowerOutlet with common optional fields.
func PowerOutletExtended() *diode.PowerOutlet {
	return &diode.PowerOutlet{
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
		Status:      diode.String("disabled"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}

// PowerOutletExplicit Creates a PowerOutlet with fully nested objects and all common fields.
func PowerOutletExplicit() *diode.PowerOutlet {
	return &diode.PowerOutlet{
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
		Status:      diode.String("disabled"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
