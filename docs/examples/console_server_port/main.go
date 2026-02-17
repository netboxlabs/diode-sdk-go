// Package main demonstrates ingesting ConsoleServerPort entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "console_server_port-example"
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
	consoleServerPort := ConsoleServerPortMinimal()
	// consoleServerPort := ConsoleServerPortExtended()
	// consoleServerPort := ConsoleServerPortExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{consoleServerPort})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ConsoleServerPort ingested successfully")
	}
}

// ConsoleServerPortMinimal Creates a ConsoleServerPort with only required fields.
func ConsoleServerPortMinimal() *diode.ConsoleServerPort {
	return &diode.ConsoleServerPort{
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

// ConsoleServerPortExtended Creates a ConsoleServerPort with common optional fields.
func ConsoleServerPortExtended() *diode.ConsoleServerPort {
	return &diode.ConsoleServerPort{
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
		Name:          diode.String("Example Name"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		Label:         diode.String("Example Label"),
		Type:          diode.String("db-25"),
		Speed:         diode.Int64(1),
		MarkConnected: diode.Bool(true),
	}
}

// ConsoleServerPortExplicit Creates a ConsoleServerPort with fully nested objects and all common fields.
func ConsoleServerPortExplicit() *diode.ConsoleServerPort {
	return &diode.ConsoleServerPort{
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
		Name:          diode.String("Example Name"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		Label:         diode.String("Example Label"),
		Type:          diode.String("db-25"),
		Speed:         diode.Int64(1),
		MarkConnected: diode.Bool(true),
		Module: &diode.Module{
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
			ModuleBay: &diode.ModuleBay{
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
				Metadata: diode.Metadata{"source": "example"},
			},
			ModuleType: &diode.ModuleType{
				Manufacturer: &diode.Manufacturer{
					Name:     diode.String("Example Name"),
					Slug:     diode.String("example-slug"),
					Metadata: diode.Metadata{"source": "example"},
				},
				Model:    diode.String("Model X"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Status:   diode.String("active"),
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
