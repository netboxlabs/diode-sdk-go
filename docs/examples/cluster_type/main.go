// Package main demonstrates ingesting ClusterType entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cluster_type-example"
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
	clusterType := ClusterTypeMinimal()
	// clusterType := ClusterTypeExtended()
	// clusterType := ClusterTypeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{clusterType})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ClusterType ingested successfully")
	}
}

// ClusterTypeMinimal Creates a ClusterType with only required fields.
func ClusterTypeMinimal() *diode.ClusterType {
	return &diode.ClusterType{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// ClusterTypeExtended Creates a ClusterType with common optional fields.
func ClusterTypeExtended() *diode.ClusterType {
	return &diode.ClusterType{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// ClusterTypeExplicit Creates a ClusterType with fully nested objects and all common fields.
func ClusterTypeExplicit() *diode.ClusterType {
	return &diode.ClusterType{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
