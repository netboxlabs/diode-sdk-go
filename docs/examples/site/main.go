// Package main demonstrates ingesting Site entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "site-example"
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
	site := SiteMinimal()
	// site := SiteExtended()
	// site := SiteExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{site})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Site ingested successfully")
	}
}

// SiteMinimal Creates a Site with only required fields.
func SiteMinimal() *diode.Site {
	return &diode.Site{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// SiteExtended Creates a Site with common optional fields.
func SiteExtended() *diode.Site {
	return &diode.Site{
		Name:            diode.String("Example Name"),
		Slug:            diode.String("example-slug"),
		Metadata:        diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:          diode.String("active"),
		Description:     diode.String("Example description"),
		Facility:        diode.String("Example Facility"),
		TimeZone:        diode.String("Example TimeZone"),
		PhysicalAddress: diode.String("Example PhysicalAddress"),
		ShippingAddress: diode.String("Example ShippingAddress"),
		Latitude:        diode.Float64(1.0),
		Longitude:       diode.Float64(1.0),
		Comments:        diode.String("Example comments"),
	}
}

// SiteExplicit Creates a Site with fully nested objects and all common fields.
func SiteExplicit() *diode.Site {
	return &diode.Site{
		Name:            diode.String("Example Name"),
		Slug:            diode.String("example-slug"),
		Metadata:        diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:          diode.String("active"),
		Description:     diode.String("Example description"),
		Comments:        diode.String("Example comments"),
		Facility:        diode.String("Example Facility"),
		TimeZone:        diode.String("Example TimeZone"),
		PhysicalAddress: diode.String("Example PhysicalAddress"),
		ShippingAddress: diode.String("Example ShippingAddress"),
		Latitude:        diode.Float64(1.0),
		Longitude:       diode.Float64(1.0),
		Region: &diode.Region{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Group: &diode.SiteGroup{
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
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
