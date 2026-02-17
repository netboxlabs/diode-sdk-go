// Package main demonstrates ingesting CustomFieldChoiceSet entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "custom_field_choice_set-example"
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
	customFieldChoiceSet := CustomFieldChoiceSetMinimal()
	// customFieldChoiceSet := CustomFieldChoiceSetExtended()
	// customFieldChoiceSet := CustomFieldChoiceSetExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{customFieldChoiceSet})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CustomFieldChoiceSet ingested successfully")
	}
}

// CustomFieldChoiceSetMinimal Creates a CustomFieldChoiceSet with only required fields.
func CustomFieldChoiceSetMinimal() *diode.CustomFieldChoiceSet {
	return &diode.CustomFieldChoiceSet{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CustomFieldChoiceSetExtended Creates a CustomFieldChoiceSet with common optional fields.
func CustomFieldChoiceSetExtended() *diode.CustomFieldChoiceSet {
	return &diode.CustomFieldChoiceSet{
		Name:                diode.String("Example Name"),
		Metadata:            diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:         diode.String("Example description"),
		BaseChoices:         diode.String("IATA"),
		OrderAlphabetically: diode.Bool(true),
	}
}

// CustomFieldChoiceSetExplicit Creates a CustomFieldChoiceSet with fully nested objects and all common fields.
func CustomFieldChoiceSetExplicit() *diode.CustomFieldChoiceSet {
	return &diode.CustomFieldChoiceSet{
		Name:                diode.String("Example Name"),
		Metadata:            diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:         diode.String("Example description"),
		BaseChoices:         diode.String("IATA"),
		OrderAlphabetically: diode.Bool(true),
		Owner: &diode.Owner{
			Name: diode.String("Example Name"),
			Group: &diode.OwnerGroup{
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
	}
}
