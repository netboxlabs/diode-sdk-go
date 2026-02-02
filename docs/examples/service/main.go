// Package main demonstrates ingesting Service entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "service-example"
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
	service := ServiceMinimal()
	// service := ServiceExtended()
	// service := ServiceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{service})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Service ingested successfully")
	}
}

// ServiceMinimal Creates a Service with only required fields.
func ServiceMinimal() *diode.Service {
	return &diode.Service{
		Name: diode.String("Example Name"),
	}
}

// ServiceExtended Creates a Service with common optional fields.
func ServiceExtended() *diode.Service {
	return &diode.Service{
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
	}
}

// ServiceExplicit Creates a Service with fully nested objects and all common fields.
func ServiceExplicit() *diode.Service {
	return &diode.Service{
		Name:        diode.String("Example Name"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
