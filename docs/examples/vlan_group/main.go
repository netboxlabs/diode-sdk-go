// Package main demonstrates ingesting VLANGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vlan_group-example"
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
	vlanGroup := VLANGroupMinimal()
	// vlanGroup := VLANGroupExtended()
	// vlanGroup := VLANGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vlanGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VLANGroup ingested successfully")
	}
}

// VLANGroupMinimal Creates a VLANGroup with only required fields.
func VLANGroupMinimal() *diode.VLANGroup {
	return &diode.VLANGroup{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VLANGroupExtended Creates a VLANGroup with common optional fields.
func VLANGroupExtended() *diode.VLANGroup {
	return &diode.VLANGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}

// VLANGroupExplicit Creates a VLANGroup with fully nested objects and all common fields.
func VLANGroupExplicit() *diode.VLANGroup {
	return &diode.VLANGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
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
		// Polymorphic 'scope' — choose ONE variant for Scope:
		Scope: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		// Scope: &diode.Cluster{ Name: diode.String("Example Name"), Type: &diode.ClusterType{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.ClusterGroup{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.Location{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.Rack{ Name: diode.String("Example Name"), Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.Region{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.SiteGroup{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.RackGroup{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
