// Package main demonstrates ingesting PowerPanel entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "power_panel-example"
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
	powerPanel := PowerPanelMinimal()
	// powerPanel := PowerPanelExtended()
	// powerPanel := PowerPanelExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{powerPanel})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("PowerPanel ingested successfully")
	}
}

// PowerPanelMinimal Creates a PowerPanel with only required fields.
func PowerPanelMinimal() *diode.PowerPanel {
	return &diode.PowerPanel{
		Site: &diode.Site{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Name: diode.String("Example Name"),
	}
}

// PowerPanelExtended Creates a PowerPanel with common optional fields.
func PowerPanelExtended() *diode.PowerPanel {
	return &diode.PowerPanel{
		Site: &diode.Site{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
	}
}

// PowerPanelExplicit Creates a PowerPanel with fully nested objects and all common fields.
func PowerPanelExplicit() *diode.PowerPanel {
	return &diode.PowerPanel{
		Site: &diode.Site{
			Name:   diode.String("Example Name"),
			Slug:   diode.String("example-slug"),
			Status: diode.String("active"),
		},
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
