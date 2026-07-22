// Package main demonstrates ingesting CircuitTermination entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "circuit_termination-example"
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
	circuitTermination := CircuitTerminationMinimal()
	// circuitTermination := CircuitTerminationExtended()
	// circuitTermination := CircuitTerminationExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{circuitTermination})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CircuitTermination ingested successfully")
	}
}

// CircuitTerminationMinimal Creates a CircuitTermination with only required fields.
func CircuitTerminationMinimal() *diode.CircuitTermination {
	return &diode.CircuitTermination{
		Circuit: &diode.Circuit{
			Cid: diode.String("CID-001"),
			Provider: &diode.Provider{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Type: &diode.CircuitType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		TermSide: diode.String("A"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CircuitTerminationExtended Creates a CircuitTermination with common optional fields.
func CircuitTerminationExtended() *diode.CircuitTermination {
	return &diode.CircuitTermination{
		Circuit: &diode.Circuit{
			Cid: diode.String("CID-001"),
			Provider: &diode.Provider{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Type: &diode.CircuitType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		TermSide:      diode.String("A"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		PortSpeed:     diode.Int64(1),
		UpstreamSpeed: diode.Int64(1),
		XconnectId:    diode.String("Example XconnectId"),
		PpInfo:        diode.String("Example PpInfo"),
		MarkConnected: diode.Bool(true),
	}
}

// CircuitTerminationExplicit Creates a CircuitTermination with fully nested objects and all common fields.
func CircuitTerminationExplicit() *diode.CircuitTermination {
	return &diode.CircuitTermination{
		Circuit: &diode.Circuit{
			Cid: diode.String("CID-001"),
			Provider: &diode.Provider{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Type: &diode.CircuitType{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Color:    diode.String("0000ff"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		TermSide:      diode.String("A"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		PortSpeed:     diode.Int64(1),
		UpstreamSpeed: diode.Int64(1),
		XconnectId:    diode.String("Example XconnectId"),
		PpInfo:        diode.String("Example PpInfo"),
		MarkConnected: diode.Bool(true),
		// Polymorphic 'termination' — choose ONE variant for Termination:
		Termination: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		// Termination: &diode.Location{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.ProviderNetwork{ Provider: &diode.Provider{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.Region{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		// Termination: &diode.SiteGroup{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
