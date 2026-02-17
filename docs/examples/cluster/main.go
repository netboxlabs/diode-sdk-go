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
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
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
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
