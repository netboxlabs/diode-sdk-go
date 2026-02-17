// Package main demonstrates ingesting DeviceRole entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "device_role-example"
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
	deviceRole := DeviceRoleMinimal()
	// deviceRole := DeviceRoleExtended()
	// deviceRole := DeviceRoleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{deviceRole})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("DeviceRole ingested successfully")
	}
}

// DeviceRoleMinimal Creates a DeviceRole with only required fields.
func DeviceRoleMinimal() *diode.DeviceRole {
	return &diode.DeviceRole{
		Name:     diode.String("Example Name"),
		Slug:     diode.String("example-slug"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// DeviceRoleExtended Creates a DeviceRole with common optional fields.
func DeviceRoleExtended() *diode.DeviceRole {
	return &diode.DeviceRole{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Color:       diode.String("0000ff"),
		VmRole:      diode.Bool(true),
		Comments:    diode.String("Example comments"),
	}
}

// DeviceRoleExplicit Creates a DeviceRole with fully nested objects and all common fields.
func DeviceRoleExplicit() *diode.DeviceRole {
	return &diode.DeviceRole{
		Name:        diode.String("Example Name"),
		Slug:        diode.String("example-slug"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Color:       diode.String("0000ff"),
		Comments:    diode.String("Example comments"),
		VmRole:      diode.Bool(true),
		Parent: &diode.DeviceRole{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Color:    diode.String("0000ff"),
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
