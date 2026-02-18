// Package main demonstrates ingesting VirtualCircuit entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "virtual_circuit-example"
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
	virtualCircuit := VirtualCircuitMinimal()
	// virtualCircuit := VirtualCircuitExtended()
	// virtualCircuit := VirtualCircuitExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{virtualCircuit})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VirtualCircuit ingested successfully")
	}
}

// VirtualCircuitMinimal Creates a VirtualCircuit with only required fields.
func VirtualCircuitMinimal() *diode.VirtualCircuit {
	return &diode.VirtualCircuit{
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
	}
}

// VirtualCircuitExtended Creates a VirtualCircuit with common optional fields.
func VirtualCircuitExtended() *diode.VirtualCircuit {
	return &diode.VirtualCircuit{
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
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}

// VirtualCircuitExplicit Creates a VirtualCircuit with fully nested objects and all common fields.
func VirtualCircuitExplicit() *diode.VirtualCircuit {
	return &diode.VirtualCircuit{
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
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		ProviderAccount: &diode.ProviderAccount{
			Provider: &diode.Provider{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Account:  diode.String("Example Account"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
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
