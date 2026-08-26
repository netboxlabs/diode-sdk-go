// Package main demonstrates ingesting InventoryItem entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "inventory_item-example"
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
	inventoryItem := InventoryItemMinimal()
	// inventoryItem := InventoryItemExtended()
	// inventoryItem := InventoryItemExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{inventoryItem})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("InventoryItem ingested successfully")
	}
}

// InventoryItemMinimal Creates a InventoryItem with only required fields.
func InventoryItemMinimal() *diode.InventoryItem {
	return &diode.InventoryItem{
		Device: &diode.Device{
			DeviceType: &diode.DeviceType{
				Manufacturer: &diode.Manufacturer{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Model:    diode.String("Model X"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Role: &diode.DeviceRole{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// InventoryItemExtended Creates a InventoryItem with common optional fields.
func InventoryItemExtended() *diode.InventoryItem {
	return &diode.InventoryItem{
		Device: &diode.Device{
			DeviceType: &diode.DeviceType{
				Manufacturer: &diode.Manufacturer{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Model:    diode.String("Model X"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Role: &diode.DeviceRole{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
		Label:       diode.String("Example Label"),
		PartId:      diode.String("Example PartId"),
		AssetTag:    diode.String("ASSET-001"),
		Discovered:  diode.Bool(true),
	}
}

// InventoryItemExplicit Creates a InventoryItem with fully nested objects and all common fields.
func InventoryItemExplicit() *diode.InventoryItem {
	return &diode.InventoryItem{
		Device: &diode.Device{
			DeviceType: &diode.DeviceType{
				Manufacturer: &diode.Manufacturer{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Model:    diode.String("Model X"),
				Slug:     diode.String("example-slug"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Role: &diode.DeviceRole{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Color:    diode.String("0000ff"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Site: &diode.Site{
				Name:     diode.String("Example Name"),
				Slug:     diode.String("example-slug"),
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:      diode.String("active"),
		Serial:      diode.String("SN-001234"),
		Description: diode.String("Example description"),
		AssetTag:    diode.String("ASSET-001"),
		Label:       diode.String("Example Label"),
		PartId:      diode.String("Example PartId"),
		Discovered:  diode.Bool(true),
		Parent: &diode.InventoryItem{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name:     diode.String("Example Name"),
						Slug:     diode.String("example-slug"),
						Metadata: diode.Metadata{"source": "example"},
					},
					Model:    diode.String("Model X"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Role: &diode.DeviceRole{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Color:    diode.String("0000ff"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Site: &diode.Site{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Status:   diode.String("active"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Role: &diode.InventoryItemRole{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Color:    diode.String("0000ff"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Manufacturer: &diode.Manufacturer{
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
		// Polymorphic 'component' — choose ONE variant for Component:
		Component: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name:     diode.String("Example Name"),
						Slug:     diode.String("example-slug"),
						Metadata: diode.Metadata{"source": "example"},
					},
					Model:    diode.String("Model X"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Role: &diode.DeviceRole{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Color:    diode.String("0000ff"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Site: &diode.Site{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Status:   diode.String("active"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Status:   diode.String("active"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Name:     diode.String("Example Name"),
			Type:     diode.String("1000base-t"),
			Metadata: diode.Metadata{"source": "example"},
		},
		// Component: &diode.ConsolePort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.ConsoleServerPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.FrontPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Type: diode.String("110-punch"), Color: diode.String("0000ff"), RearPort: &diode.RearPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Type: diode.String("110-punch"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.PowerOutlet{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Color: diode.String("0000ff"), Status: diode.String("disabled"), Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.PowerPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.RearPort{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Type: diode.String("110-punch"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.CoolingIntake{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		// Component: &diode.CoolingOutflow{ Device: &diode.Device{ DeviceType: &diode.DeviceType{ Manufacturer: &diode.Manufacturer{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Model: diode.String("Model X"), Slug: diode.String("example-slug"), Metadata: diode.Metadata{"source": "example"}, }, Role: &diode.DeviceRole{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Color: diode.String("0000ff"), Metadata: diode.Metadata{"source": "example"}, }, Site: &diode.Site{ Name: diode.String("Example Name"), Slug: diode.String("example-slug"), Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Status: diode.String("active"), Metadata: diode.Metadata{"source": "example"}, }, Name: diode.String("Example Name"), Metadata: diode.Metadata{"source": "example"}, },
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
