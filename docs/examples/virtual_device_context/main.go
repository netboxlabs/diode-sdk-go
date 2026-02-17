// Package main demonstrates ingesting VirtualDeviceContext entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_device_context-example"
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
	virtualDeviceContext := VirtualDeviceContextMinimal()
	// virtualDeviceContext := VirtualDeviceContextExtended()
	// virtualDeviceContext := VirtualDeviceContextExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualDeviceContext})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualDeviceContext ingested successfully")
	}
}

// VirtualDeviceContextMinimal Creates a VirtualDeviceContext with only required fields.
func VirtualDeviceContextMinimal() *diode.VirtualDeviceContext {
	return &diode.VirtualDeviceContext{
		Name: diode.String("Example Name"),
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
				Metadata: diode.Metadata{"source": "example"},
			},
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Status:   diode.String("active"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VirtualDeviceContextExtended Creates a VirtualDeviceContext with common optional fields.
func VirtualDeviceContextExtended() *diode.VirtualDeviceContext {
	return &diode.VirtualDeviceContext{
		Name: diode.String("Example Name"),
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
				Metadata: diode.Metadata{"source": "example"},
			},
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Status:      diode.String("active"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Identifier:  diode.Int64(1),
		Comments:    diode.String("Example comments"),
	}
}

// VirtualDeviceContextExplicit Creates a VirtualDeviceContext with fully nested objects and all common fields.
func VirtualDeviceContextExplicit() *diode.VirtualDeviceContext {
	return &diode.VirtualDeviceContext{
		Name: diode.String("Example Name"),
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
		Status:      diode.String("active"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Identifier:  diode.Int64(1),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		PrimaryIp4: &diode.IPAddress{
			Address:  diode.String("192.0.2.1/32"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		PrimaryIp6: &diode.IPAddress{
			Address:  diode.String("192.0.2.1/32"),
			Status:   diode.String("active"),
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
