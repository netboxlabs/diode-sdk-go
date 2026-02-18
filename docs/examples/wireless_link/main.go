// Package main demonstrates ingesting WirelessLink entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "wireless_link-example"
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
	wirelessLink := WirelessLinkMinimal()
	// wirelessLink := WirelessLinkExtended()
	// wirelessLink := WirelessLinkExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{wirelessLink})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("WirelessLink ingested successfully")
	}
}

// WirelessLinkMinimal Creates a WirelessLink with only required fields.
func WirelessLinkMinimal() *diode.WirelessLink {
	return &diode.WirelessLink{
		InterfaceA: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		InterfaceB: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata: diode.Metadata{"source": "example"},
	}
}

// WirelessLinkExtended Creates a WirelessLink with common optional fields.
func WirelessLinkExtended() *diode.WirelessLink {
	return &diode.WirelessLink{
		InterfaceA: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		InterfaceB: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:     diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:       diode.String("connected"),
		Description:  diode.String("Example description"),
		Ssid:         diode.String("ExampleSSID"),
		AuthType:     diode.String("open"),
		AuthCipher:   diode.String("aes"),
		AuthPsk:      diode.String("Example AuthPsk"),
		Distance:     diode.Float64(1.0),
		DistanceUnit: diode.String("ft"),
		Comments:     diode.String("Example comments"),
	}
}

// WirelessLinkExplicit Creates a WirelessLink with fully nested objects and all common fields.
func WirelessLinkExplicit() *diode.WirelessLink {
	return &diode.WirelessLink{
		InterfaceA: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		InterfaceB: &diode.Interface{
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
			Type:     diode.String("Example Type"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Metadata:     diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:       diode.String("connected"),
		Description:  diode.String("Example description"),
		Comments:     diode.String("Example comments"),
		Ssid:         diode.String("ExampleSSID"),
		AuthType:     diode.String("open"),
		AuthCipher:   diode.String("aes"),
		AuthPsk:      diode.String("Example AuthPsk"),
		Distance:     diode.Float64(1.0),
		DistanceUnit: diode.String("ft"),
		Tenant: &diode.Tenant{
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
