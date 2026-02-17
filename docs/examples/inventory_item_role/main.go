// Package main demonstrates ingesting InventoryItemRole entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "inventory_item_role-example"
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
	inventoryItemRole := InventoryItemRoleMinimal()
	// inventoryItemRole := InventoryItemRoleExtended()
	// inventoryItemRole := InventoryItemRoleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{inventoryItemRole})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("InventoryItemRole ingested successfully")
	}
}

// InventoryItemRoleMinimal Creates a InventoryItemRole with only required fields.
func InventoryItemRoleMinimal() *diode.InventoryItemRole {
	return &diode.InventoryItemRole{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// InventoryItemRoleExtended Creates a InventoryItemRole with common optional fields.
func InventoryItemRoleExtended() *diode.InventoryItemRole {
	return &diode.InventoryItemRole{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Color:       diode.String("0000ff"),
		Comments:    diode.String("Example comments"),
	}
}

// InventoryItemRoleExplicit Creates a InventoryItemRole with fully nested objects and all common fields.
func InventoryItemRoleExplicit() *diode.InventoryItemRole {
	return &diode.InventoryItemRole{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Color:       diode.String("0000ff"),
		Comments:    diode.String("Example comments"),
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
