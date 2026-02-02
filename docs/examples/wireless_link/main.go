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
						Name: diode.String("Example Name"),
						Slug: diode.String("example-slug"),
					},
					Model: diode.String("Model X"),
					Slug:  diode.String("example-slug"),
				},
				Role: &diode.DeviceRole{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Site: &diode.Site{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
			},
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
		InterfaceB: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Example Name"),
						Slug: diode.String("example-slug"),
					},
					Model: diode.String("Model X"),
					Slug:  diode.String("example-slug"),
				},
				Role: &diode.DeviceRole{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Site: &diode.Site{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
			},
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
	}
}

// WirelessLinkExtended Creates a WirelessLink with common optional fields.
func WirelessLinkExtended() *diode.WirelessLink {
	return &diode.WirelessLink{
		InterfaceA: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Example Name"),
						Slug: diode.String("example-slug"),
					},
					Model: diode.String("Model X"),
					Slug:  diode.String("example-slug"),
				},
				Role: &diode.DeviceRole{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Site: &diode.Site{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
			},
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
		InterfaceB: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Example Name"),
						Slug: diode.String("example-slug"),
					},
					Model: diode.String("Model X"),
					Slug:  diode.String("example-slug"),
				},
				Role: &diode.DeviceRole{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
				Site: &diode.Site{
					Name: diode.String("Example Name"),
					Slug: diode.String("example-slug"),
				},
			},
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
		Status:      diode.String("connected"),
		Description: diode.String("Example description"),
	}
}

// WirelessLinkExplicit Creates a WirelessLink with fully nested objects and all common fields.
func WirelessLinkExplicit() *diode.WirelessLink {
	return &diode.WirelessLink{
		InterfaceA: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Example Name"),
						Slug: diode.String("example-slug"),
					},
					Model: diode.String("Model X"),
					Slug:  diode.String("example-slug"),
				},
				Role: &diode.DeviceRole{
					Name:  diode.String("Example Name"),
					Slug:  diode.String("example-slug"),
					Color: diode.String("0000ff"),
				},
				Site: &diode.Site{
					Name:   diode.String("Example Name"),
					Slug:   diode.String("example-slug"),
					Status: diode.String("active"),
				},
				Status: diode.String("active"),
			},
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
		InterfaceB: &diode.Interface{
			Device: &diode.Device{
				DeviceType: &diode.DeviceType{
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Example Name"),
						Slug: diode.String("example-slug"),
					},
					Model: diode.String("Model X"),
					Slug:  diode.String("example-slug"),
				},
				Role: &diode.DeviceRole{
					Name:  diode.String("Example Name"),
					Slug:  diode.String("example-slug"),
					Color: diode.String("0000ff"),
				},
				Site: &diode.Site{
					Name:   diode.String("Example Name"),
					Slug:   diode.String("example-slug"),
					Status: diode.String("active"),
				},
				Status: diode.String("active"),
			},
			Name: diode.String("Example Name"),
			Type: diode.String("Example Type"),
		},
		Status:      diode.String("connected"),
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tenant: &diode.Tenant{
			Name: diode.String("Example Name"),
			Slug: diode.String("example-slug"),
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
