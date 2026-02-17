// Package main demonstrates ingesting SiteGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "site_group-example"
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
	siteGroup := SiteGroupMinimal()
	// siteGroup := SiteGroupExtended()
	// siteGroup := SiteGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{siteGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("SiteGroup ingested successfully")
	}
}

// SiteGroupMinimal Creates a SiteGroup with only required fields.
func SiteGroupMinimal() *diode.SiteGroup {
	return &diode.SiteGroup{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// SiteGroupExtended Creates a SiteGroup with common optional fields.
func SiteGroupExtended() *diode.SiteGroup {
	return &diode.SiteGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// SiteGroupExplicit Creates a SiteGroup with fully nested objects and all common fields.
func SiteGroupExplicit() *diode.SiteGroup {
	return &diode.SiteGroup{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
