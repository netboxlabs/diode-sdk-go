// Package main demonstrates ingesting FHRPGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "fhrp_group-example"
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
	fhrpGroup := FHRPGroupMinimal()
	// fhrpGroup := FHRPGroupExtended()
	// fhrpGroup := FHRPGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{fhrpGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("FHRPGroup ingested successfully")
	}
}

// FHRPGroupMinimal Creates a FHRPGroup with only required fields.
func FHRPGroupMinimal() *diode.FHRPGroup {
	return &diode.FHRPGroup{
		Protocol: diode.String("carp"),
		GroupId:  diode.Int64(1),
	}
}

// FHRPGroupExtended Creates a FHRPGroup with common optional fields.
func FHRPGroupExtended() *diode.FHRPGroup {
	return &diode.FHRPGroup{
		Protocol:    diode.String("carp"),
		GroupId:     diode.Int64(1),
		Description: diode.String("Example description"),
	}
}

// FHRPGroupExplicit Creates a FHRPGroup with fully nested objects and all common fields.
func FHRPGroupExplicit() *diode.FHRPGroup {
	return &diode.FHRPGroup{
		Protocol:    diode.String("carp"),
		GroupId:     diode.Int64(1),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
