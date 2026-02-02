// Package main demonstrates ingesting Role entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "role-example"
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
	role := RoleMinimal()
	// role := RoleExtended()
	// role := RoleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{role})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Role ingested successfully")
	}
}

// RoleMinimal Creates a Role with only required fields.
func RoleMinimal() *diode.Role {
	return &diode.Role{
		Name: diode.String("Example Name"),
		Slug: diode.String("example-slug"),
	}
}

// RoleExtended Creates a Role with common optional fields.
func RoleExtended() *diode.Role {
	return &diode.Role{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
	}
}

// RoleExplicit Creates a Role with fully nested objects and all common fields.
func RoleExplicit() *diode.Role {
	return &diode.Role{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
