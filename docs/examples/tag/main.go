// Package main demonstrates ingesting Tag entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tag-example"
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
	tag := TagMinimal()
	// tag := TagExtended()
	// tag := TagExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tag})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Tag ingested successfully")
	}
}

// TagMinimal Creates a Tag with only required fields.
func TagMinimal() *diode.Tag {
	return &diode.Tag{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// TagExtended Creates a Tag with common optional fields.
func TagExtended() *diode.Tag {
	return &diode.Tag{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}

// TagExplicit Creates a Tag with fully nested objects and all common fields.
func TagExplicit() *diode.Tag {
	return &diode.Tag{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Color:       diode.String("0000ff"),
		Description: diode.String("Example description"),
	}
}
