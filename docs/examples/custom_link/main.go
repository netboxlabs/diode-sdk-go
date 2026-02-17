// Package main demonstrates ingesting CustomLink entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "custom_link-example"
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
	customLink := CustomLinkMinimal()
	// customLink := CustomLinkExtended()
	// customLink := CustomLinkExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{customLink})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("CustomLink ingested successfully")
	}
}

// CustomLinkMinimal Creates a CustomLink with only required fields.
func CustomLinkMinimal() *diode.CustomLink {
	return &diode.CustomLink{
		Name:     diode.String("Example Name"),
		LinkText: diode.String("Example LinkText"),
		LinkUrl:  diode.String("Example LinkUrl"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// CustomLinkExtended Creates a CustomLink with common optional fields.
func CustomLinkExtended() *diode.CustomLink {
	return &diode.CustomLink{
		Name:        diode.String("Example Name"),
		LinkText:    diode.String("Example LinkText"),
		LinkUrl:     diode.String("Example LinkUrl"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Enabled:     diode.Bool(true),
		Weight:      diode.Int64(1),
		GroupName:   diode.String("Example GroupName"),
		ButtonClass: diode.String("black"),
		NewWindow:   diode.Bool(true),
	}
}

// CustomLinkExplicit Creates a CustomLink with fully nested objects and all common fields.
func CustomLinkExplicit() *diode.CustomLink {
	return &diode.CustomLink{
		Name:        diode.String("Example Name"),
		LinkText:    diode.String("Example LinkText"),
		LinkUrl:     diode.String("Example LinkUrl"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Enabled:     diode.Bool(true),
		Weight:      diode.Int64(1),
		GroupName:   diode.String("Example GroupName"),
		ButtonClass: diode.String("black"),
		NewWindow:   diode.Bool(true),
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
