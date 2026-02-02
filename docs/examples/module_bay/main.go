// Package main demonstrates ingesting ModuleBay entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "module_bay-example"
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
	moduleBay := ModuleBayMinimal()
	// moduleBay := ModuleBayExtended()
	// moduleBay := ModuleBayExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{moduleBay})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ModuleBay ingested successfully")
	}
}

// ModuleBayMinimal Creates a ModuleBay with only required fields.
func ModuleBayMinimal() *diode.ModuleBay {
	return &diode.ModuleBay{
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

// ModuleBayExtended Creates a ModuleBay with common optional fields.
func ModuleBayExtended() *diode.ModuleBay {
	return &diode.ModuleBay{
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
		Description: diode.String("Example description"),
	}
}

// ModuleBayExplicit Creates a ModuleBay with fully nested objects and all common fields.
func ModuleBayExplicit() *diode.ModuleBay {
	return &diode.ModuleBay{
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
		Description: diode.String("Example description"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
