// Package main demonstrates ingesting ConfigContext entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "config_context-example"
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
	configContext := ConfigContextMinimal()
	// configContext := ConfigContextExtended()
	// configContext := ConfigContextExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{configContext})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ConfigContext ingested successfully")
	}
}

// ConfigContextMinimal Creates a ConfigContext with only required fields.
func ConfigContextMinimal() *diode.ConfigContext {
	return &diode.ConfigContext{
		Name:     diode.String("Example Name"),
		Data:     diode.String("Example Data"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ConfigContextExtended Creates a ConfigContext with common optional fields.
func ConfigContextExtended() *diode.ConfigContext {
	return &diode.ConfigContext{
		Name:        diode.String("Example Name"),
		Data:        diode.String("Example Data"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Weight:      diode.Int64(1),
		IsActive:    diode.Bool(true),
	}
}

// ConfigContextExplicit Creates a ConfigContext with fully nested objects and all common fields.
func ConfigContextExplicit() *diode.ConfigContext {
	return &diode.ConfigContext{
		Name:        diode.String("Example Name"),
		Data:        diode.String("Example Data"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Weight:      diode.Int64(1),
		IsActive:    diode.Bool(true),
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
