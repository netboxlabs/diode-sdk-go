// Package main demonstrates ingesting VirtualMachine entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_machine-example"
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
	virtualMachine := VirtualMachineMinimal()
	// virtualMachine := VirtualMachineExtended()
	// virtualMachine := VirtualMachineExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualMachine})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualMachine ingested successfully")
	}
}

// VirtualMachineMinimal Creates a VirtualMachine with only required fields.
func VirtualMachineMinimal() *diode.VirtualMachine {
	return &diode.VirtualMachine{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VirtualMachineExtended Creates a VirtualMachine with common optional fields.
func VirtualMachineExtended() *diode.VirtualMachine {
	return &diode.VirtualMachine{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
		Vcpus:       diode.Float64(1.0),
		Memory:      diode.Int64(1),
		Disk:        diode.Int64(1),
		Comments:    diode.String("Example comments"),
		StartOnBoot: diode.String("laststate"),
	}
}

// VirtualMachineExplicit Creates a VirtualMachine with fully nested objects and all common fields.
func VirtualMachineExplicit() *diode.VirtualMachine {
	return &diode.VirtualMachine{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Vcpus:       diode.Float64(1.0),
		Memory:      diode.Int64(1),
		Disk:        diode.Int64(1),
		StartOnBoot: diode.String("laststate"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
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
		Role: &diode.DeviceRole{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Color:    diode.String("0000ff"),
			Metadata: diode.Metadata{"source": "example"},
		},
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
		VirtualMachineType: &diode.VirtualMachineType{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
