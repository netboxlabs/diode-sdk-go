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
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:        diode.String("active"),
		Serial:        diode.String("SN-001234"),
		Description:   diode.String("Example description"),
		Name:          diode.String("Example Name"),
		AssetTag:      diode.String("ASSET-001"),
		Position:      diode.Float64(1.0),
		Face:          diode.String("front"),
		Latitude:      diode.Float64(1.0),
		Longitude:     diode.Float64(1.0),
		Airflow:       diode.String("bottom-to-top"),
		VcPosition:    diode.Int64(1),
		VcPriority:    diode.Int64(1),
		Comments:      diode.String("Example comments"),
		CoolingMethod: diode.String("air"),
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
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:        diode.String("active"),
		Serial:        diode.String("SN-001234"),
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		AssetTag:      diode.String("ASSET-001"),
		Name:          diode.String("Example Name"),
		Position:      diode.Float64(1.0),
		Face:          diode.String("front"),
		Latitude:      diode.Float64(1.0),
		Longitude:     diode.Float64(1.0),
		Airflow:       diode.String("bottom-to-top"),
		VcPosition:    diode.Int64(1),
		VcPriority:    diode.Int64(1),
		CoolingMethod: diode.String("air"),
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
		Location: &diode.Location{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
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
		OobIp: &diode.IPAddress{
			Address:  diode.String("192.0.2.1/32"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Cluster: &diode.Cluster{
			Name: diode.String("Example Name"),
			Type: &diode.ClusterType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		VirtualChassis: &diode.VirtualChassis{
			Name:     diode.String("Example Name"),
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
		Config: &diode.DeviceConfig{
			Startup:   []byte("example data"),
			Running:   []byte("example data"),
			Candidate: []byte("example data"),
			Metadata:  diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
