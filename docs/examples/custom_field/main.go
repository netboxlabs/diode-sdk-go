// Package main demonstrates ingesting CustomField entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "custom_field-example"
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
	customField := CustomFieldMinimal()
	// customField := CustomFieldExtended()
	// customField := CustomFieldExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{customField})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CustomField ingested successfully")
	}
}

// CustomFieldMinimal Creates a CustomField with only required fields.
func CustomFieldMinimal() *diode.CustomField {
	return &diode.CustomField{
		Type:     diode.String("boolean"),
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CustomFieldExtended Creates a CustomField with common optional fields.
func CustomFieldExtended() *diode.CustomField {
	return &diode.CustomField{
		Type:                diode.String("boolean"),
		Name:                diode.String("Example Name"),
		Metadata:            diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:         diode.String("Example description"),
		RelatedObjectType:   diode.String("circuits.circuit"),
		Label:               diode.String("Example Label"),
		GroupName:           diode.String("Example GroupName"),
		Required:            diode.Bool(true),
		Unique:              diode.Bool(true),
		SearchWeight:        diode.Int64(1),
		FilterLogic:         diode.String("disabled"),
		UiVisible:           diode.String("always"),
		UiEditable:          diode.String("hidden"),
		IsCloneable:         diode.Bool(true),
		Default:             diode.String("Example Default"),
		RelatedObjectFilter: diode.String("Example RelatedObjectFilter"),
		Weight:              diode.Int64(1),
		ValidationMinimum:   diode.Float64(1.0),
		ValidationMaximum:   diode.Float64(1.0),
		ValidationRegex:     diode.String("Example ValidationRegex"),
		Comments:            diode.String("Example comments"),
	}
}

// CustomFieldExplicit Creates a CustomField with fully nested objects and all common fields.
func CustomFieldExplicit() *diode.CustomField {
	return &diode.CustomField{
		Type:                diode.String("boolean"),
		Name:                diode.String("Example Name"),
		Metadata:            diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:         diode.String("Example description"),
		Comments:            diode.String("Example comments"),
		RelatedObjectType:   diode.String("circuits.circuit"),
		Label:               diode.String("Example Label"),
		GroupName:           diode.String("Example GroupName"),
		Required:            diode.Bool(true),
		Unique:              diode.Bool(true),
		SearchWeight:        diode.Int64(1),
		FilterLogic:         diode.String("disabled"),
		UiVisible:           diode.String("always"),
		UiEditable:          diode.String("hidden"),
		IsCloneable:         diode.Bool(true),
		Default:             diode.String("Example Default"),
		RelatedObjectFilter: diode.String("Example RelatedObjectFilter"),
		Weight:              diode.Int64(1),
		ValidationMinimum:   diode.Float64(1.0),
		ValidationMaximum:   diode.Float64(1.0),
		ValidationRegex:     diode.String("Example ValidationRegex"),
		ChoiceSet: &diode.CustomFieldChoiceSet{
			Name:     diode.String("Example Name"),
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
	}
}
