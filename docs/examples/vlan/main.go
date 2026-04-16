// Package main demonstrates ingesting VLAN entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vlan-example"
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
	vlan := VLANMinimal()
	// vlan := VLANExtended()
	// vlan := VLANExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vlan})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VLAN ingested successfully")
	}
}

// VLANMinimal Creates a VLAN with only required fields.
func VLANMinimal() *diode.VLAN {
	return &diode.VLAN{
		Vid:      diode.Int64(1),
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VLANExtended Creates a VLAN with common optional fields.
func VLANExtended() *diode.VLAN {
	return &diode.VLAN{
		Vid:         diode.Int64(1),
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		QinqRole:    diode.String("cvlan"),
		Comments:    diode.String("Example comments"),
	}
}

// VLANExplicit Creates a VLAN with fully nested objects and all common fields.
func VLANExplicit() *diode.VLAN {
	return &diode.VLAN{
		Vid:         diode.Int64(1),
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		QinqRole:    diode.String("cvlan"),
		Site: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Group: &diode.VLANGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Role: &diode.Role{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		QinqSvlan: &diode.VLAN{
			Vid:      diode.Int64(1),
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
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
