// Package main demonstrates ingesting Cluster entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cluster-example"
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
	cluster := ClusterMinimal()
	// cluster := ClusterExtended()
	// cluster := ClusterExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{cluster})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Cluster ingested successfully")
	}
}

// ClusterMinimal Creates a Cluster with only required fields.
func ClusterMinimal() *diode.Cluster {
	return &diode.Cluster{
		Name: diode.String("Example Name"),
		Type: &diode.ClusterType{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ClusterExtended Creates a Cluster with common optional fields.
func ClusterExtended() *diode.Cluster {
	return &diode.Cluster{
		Name: diode.String("Example Name"),
		Type: &diode.ClusterType{
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

// ClusterExplicit Creates a Cluster with fully nested objects and all common fields.
func ClusterExplicit() *diode.Cluster {
	return &diode.Cluster{
		Name: diode.String("Example Name"),
		Type: &diode.ClusterType{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Group: &diode.ClusterGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
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
		// Polymorphic 'scope' — choose ONE variant for Scope:
		Scope: &diode.Site{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		// Scope: &diode.Location{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.Region{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		// Scope: &diode.SiteGroup{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, },
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
