// Package main demonstrates ingesting Device entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "device-example"
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
	device := DeviceMinimal()
	// device := DeviceExtended()
	// device := DeviceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{device})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Device ingested successfully")
	}
}

// DeviceMinimal Creates a Device with only required fields.
func DeviceMinimal() *diode.Device {
	return &diode.Device{
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
			Metadata: diode.Metadata{"source": "example"},
		},
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// DeviceExtended Creates a Device with common optional fields.
func DeviceExtended() *diode.Device {
	return &diode.Device{
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
			Metadata: diode.Metadata{"source": "example"},
		},
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example"},
		Serial:      diode.String("SN-001234"),
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
	}
}

// DeviceExplicit Creates a Device with fully nested objects and all common fields.
func DeviceExplicit() *diode.Device {
	return &diode.Device{
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
		Metadata:    diode.Metadata{"source": "example"},
		Serial:      diode.String("SN-001234"),
		AssetTag:    diode.String("ASSET-001"),
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Platform: &diode.Platform{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
