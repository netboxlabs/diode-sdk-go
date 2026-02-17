// Package main demonstrates ingesting VirtualCircuitTermination entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_circuit_termination-example"
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
	virtualCircuitTermination := VirtualCircuitTerminationMinimal()
	// virtualCircuitTermination := VirtualCircuitTerminationExtended()
	// virtualCircuitTermination := VirtualCircuitTerminationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualCircuitTermination})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualCircuitTermination ingested successfully")
	}
}

// VirtualCircuitTerminationMinimal Creates a VirtualCircuitTermination with only required fields.
func VirtualCircuitTerminationMinimal() *diode.VirtualCircuitTermination {
	return &diode.VirtualCircuitTermination{
		VirtualCircuit: &diode.VirtualCircuit{
			Cid: diode.String("CID-001"),
			ProviderNetwork: &diode.ProviderNetwork{
				Provider: &diode.Provider{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Type: &diode.VirtualCircuitType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Interface: &diode.Interface{
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
			Name:     diode.String("Example Name"),
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VirtualCircuitTerminationExtended Creates a VirtualCircuitTermination with common optional fields.
func VirtualCircuitTerminationExtended() *diode.VirtualCircuitTermination {
	return &diode.VirtualCircuitTermination{
		VirtualCircuit: &diode.VirtualCircuit{
			Cid: diode.String("CID-001"),
			ProviderNetwork: &diode.ProviderNetwork{
				Provider: &diode.Provider{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Type: &diode.VirtualCircuitType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Interface: &diode.Interface{
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
			Name:     diode.String("Example Name"),
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Role:        diode.String("hub"),
	}
}

// VirtualCircuitTerminationExplicit Creates a VirtualCircuitTermination with fully nested objects and all common fields.
func VirtualCircuitTerminationExplicit() *diode.VirtualCircuitTermination {
	return &diode.VirtualCircuitTermination{
		VirtualCircuit: &diode.VirtualCircuit{
			Cid: diode.String("CID-001"),
			ProviderNetwork: &diode.ProviderNetwork{
				Provider: &diode.Provider{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Type: &diode.VirtualCircuitType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Color:    diode.String("0000ff"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Interface: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Role:        diode.String("hub"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
