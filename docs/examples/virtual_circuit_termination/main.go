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
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Name: diode.String("Example Name"),
			},
			Type: &diode.VirtualCircuitType{
				Name: diode.String("Example Name"),
				Slug: diode.String("example-slug"),
			},
		},
		Interface: &diode.Interface{
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
			Type: diode.String("Example Type"),
		},
	}
}

// VirtualCircuitTerminationExtended Creates a VirtualCircuitTermination with common optional fields.
func VirtualCircuitTerminationExtended() *diode.VirtualCircuitTermination {
	return &diode.VirtualCircuitTermination{
		VirtualCircuit: &diode.VirtualCircuit{
			Cid: diode.String("CID-001"),
			ProviderNetwork: &diode.ProviderNetwork{
				Provider: &diode.Provider{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Name: diode.String("Example Name"),
			},
			Type: &diode.VirtualCircuitType{
				Name: diode.String("Example Name"),
				Slug: diode.String("example-slug"),
			},
		},
		Interface: &diode.Interface{
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
			Type: diode.String("Example Type"),
		},
		Description: diode.String("Example description"),
	}
}

// VirtualCircuitTerminationExplicit Creates a VirtualCircuitTermination with fully nested objects and all common fields.
func VirtualCircuitTerminationExplicit() *diode.VirtualCircuitTermination {
	return &diode.VirtualCircuitTermination{
		VirtualCircuit: &diode.VirtualCircuit{
			Cid: diode.String("CID-001"),
			ProviderNetwork: &diode.ProviderNetwork{
				Provider: &diode.Provider{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Name: diode.String("Example Name"),
			},
			Type: &diode.VirtualCircuitType{
				Name:  diode.String("Example Name"),
				Slug:  diode.String("example-slug"),
				Color: diode.String("0000ff"),
			},
			Status: diode.String("active"),
		},
		Interface: &diode.Interface{
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
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
		Description: diode.String("Example description"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
