// Package main demonstrates ingesting Contact entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "contact-example"
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
	contact := ContactMinimal()
	// contact := ContactExtended()
	// contact := ContactExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{contact})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Contact ingested successfully")
	}
}

// ContactMinimal Creates a Contact with only required fields.
func ContactMinimal() *diode.Contact {
	return &diode.Contact{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ContactExtended Creates a Contact with common optional fields.
func ContactExtended() *diode.Contact {
	return &diode.Contact{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Title:       diode.String("Example Title"),
		Phone:       diode.String("Example Phone"),
		Email:       diode.String("Example Email"),
		Address:     diode.String("192.0.2.1/32"),
		Link:        diode.String("Example Link"),
		Comments:    diode.String("Example comments"),
	}
}

// ContactExplicit Creates a Contact with fully nested objects and all common fields.
func ContactExplicit() *diode.Contact {
	return &diode.Contact{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Title:       diode.String("Example Title"),
		Phone:       diode.String("Example Phone"),
		Email:       diode.String("Example Email"),
		Address:     diode.String("192.0.2.1/32"),
		Link:        diode.String("Example Link"),
		Group: &diode.ContactGroup{
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
