// Package main demonstrates ingesting PowerFeed entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "power_feed-example"
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
	powerFeed := PowerFeedMinimal()
	// powerFeed := PowerFeedExtended()
	// powerFeed := PowerFeedExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{powerFeed})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("PowerFeed ingested successfully")
	}
}

// PowerFeedMinimal Creates a PowerFeed with only required fields.
func PowerFeedMinimal() *diode.PowerFeed {
	return &diode.PowerFeed{
		PowerPanel: &diode.PowerPanel{
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// PowerFeedExtended Creates a PowerFeed with common optional fields.
func PowerFeedExtended() *diode.PowerFeed {
	return &diode.PowerFeed{
		PowerPanel: &diode.PowerPanel{
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Status:      diode.String("active"),
		Description: diode.String("Example description"),
	}
}

// PowerFeedExplicit Creates a PowerFeed with fully nested objects and all common fields.
func PowerFeedExplicit() *diode.PowerFeed {
	return &diode.PowerFeed{
		PowerPanel: &diode.PowerPanel{
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
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
