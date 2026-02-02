// Package main demonstrates ingesting Interface entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "interface-example"
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
	interfaceEntity := InterfaceMinimal()
	// interfaceEntity := InterfaceExtended()
	// interfaceEntity := InterfaceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{interfaceEntity})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Interface ingested successfully")
	}
}

// InterfaceMinimal Creates a Interface with only required fields.
func InterfaceMinimal() *diode.Interface {
	return &diode.Interface{
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
		Type: diode.String("1000base-bx10-d"),
	}
}

// InterfaceExtended Creates a Interface with common optional fields.
func InterfaceExtended() *diode.Interface {
	return &diode.Interface{
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
		Type:        diode.String("1000base-bx10-d"),
		Description: diode.String("Example description"),
	}
}

// InterfaceExplicit Creates a Interface with fully nested objects and all common fields.
func InterfaceExplicit() *diode.Interface {
	return &diode.Interface{
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
		Type:        diode.String("1000base-bx10-d"),
		Description: diode.String("Example description"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
