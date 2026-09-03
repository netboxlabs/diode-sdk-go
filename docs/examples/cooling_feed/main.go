// Package main demonstrates ingesting CoolingFeed entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "cooling_feed-example"
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
	coolingFeed := CoolingFeedMinimal()
	// coolingFeed := CoolingFeedExtended()
	// coolingFeed := CoolingFeedExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{coolingFeed})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CoolingFeed ingested successfully")
	}
}

// CoolingFeedMinimal Creates a CoolingFeed with only required fields.
func CoolingFeedMinimal() *diode.CoolingFeed {
	return &diode.CoolingFeed{
		CoolingSource: &diode.CoolingSource{
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Type:     diode.String("chiller"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CoolingFeedExtended Creates a CoolingFeed with common optional fields.
func CoolingFeedExtended() *diode.CoolingFeed {
	return &diode.CoolingFeed{
		CoolingSource: &diode.CoolingSource{
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Type:     diode.String("chiller"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:            diode.String("Example Name"),
		Metadata:        diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:          diode.String("active"),
		Description:     diode.String("Example description"),
		CoolingCapacity: diode.Float64(1.0),
		MaxFlow:         diode.Float64(1.0),
		MaxFlowUnit:     diode.String("gpm"),
		Comments:        diode.String("Example comments"),
	}
}

// CoolingFeedExplicit Creates a CoolingFeed with fully nested objects and all common fields.
func CoolingFeedExplicit() *diode.CoolingFeed {
	return &diode.CoolingFeed{
		CoolingSource: &diode.CoolingSource{
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Type:     diode.String("chiller"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:            diode.String("Example Name"),
		Metadata:        diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:          diode.String("active"),
		Description:     diode.String("Example description"),
		Comments:        diode.String("Example comments"),
		CoolingCapacity: diode.Float64(1.0),
		MaxFlow:         diode.Float64(1.0),
		MaxFlowUnit:     diode.String("gpm"),
		Rack: &diode.Rack{
			Name: diode.String("Example Name"),
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
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
