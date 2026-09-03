// Package main demonstrates ingesting Service entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "service-example"
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
	service := ServiceMinimal()
	// service := ServiceExtended()
	// service := ServiceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{service})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Service ingested successfully")
	}
}

// ServiceMinimal Creates a Service with only required fields.
func ServiceMinimal() *diode.Service {
	return &diode.Service{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ServiceExtended Creates a Service with common optional fields.
func ServiceExtended() *diode.Service {
	return &diode.Service{
		Name:         diode.String("Example Name"),
		Metadata:     diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:  diode.String("Example description"),
		Comments:     diode.String("Example comments"),
		PortMappings: []string{"tcp/80", "udp/53"},
	}
}

// ServiceExplicit Creates a Service with fully nested objects and all common fields.
func ServiceExplicit() *diode.Service {
	return &diode.Service{
		Name:         diode.String("Example Name"),
		Metadata:     diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:  diode.String("Example description"),
		Comments:     diode.String("Example comments"),
		PortMappings: []string{"tcp/80", "udp/53"},
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
		VirtualMachine: &diode.VirtualMachine{
			Name:     diode.String("Example Name"),
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
		// Polymorphic 'parent_object' — choose ONE variant for ParentObject:
		ParentObject: &diode.Device{
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
		// ParentObject: &diode.FHRPGroup{ Protocol: diode.String("carp"), GroupId: diode.Int64(1), Metadata: diode.Metadata{"source": "example"}, },
		// ParentObject: &diode.VirtualMachine{ Name: diode.String("Example Name"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
