// Package main demonstrates ingesting Interface entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "interface-example"
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
	interfaceEntity := InterfaceMinimal()
	// interfaceEntity := InterfaceExtended()
	// interfaceEntity := InterfaceExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{interfaceEntity})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Interface ingested successfully")
	}
}

// InterfaceMinimal Creates a Interface with only required fields.
func InterfaceMinimal() *diode.Interface {
	return &diode.Interface{
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
		Type:     diode.String("1000base-t"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// InterfaceExtended Creates a Interface with common optional fields.
func InterfaceExtended() *diode.Interface {
	return &diode.Interface{
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
		Name:               diode.String("Example Name"),
		Type:               diode.String("1000base-t"),
		Metadata:           diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:        diode.String("Example description"),
		Label:              diode.String("Example Label"),
		Enabled:            diode.Bool(true),
		Mtu:                diode.Int64(1),
		Speed:              diode.Int64(1),
		Duplex:             diode.String("auto"),
		Wwn:                diode.String("Example Wwn"),
		MgmtOnly:           diode.Bool(true),
		Mode:               diode.String("access"),
		RfRole:             diode.String("ap"),
		RfChannel:          diode.String("2.4g-1-2412-22"),
		PoeMode:            diode.String("pse"),
		PoeType:            diode.String("type1-ieee802.3af"),
		RfChannelFrequency: diode.Float64(1.0),
		RfChannelWidth:     diode.Float64(1.0),
		TxPower:            diode.Int64(1),
		MarkConnected:      diode.Bool(true),
		Channels:           diode.Int64(1),
		MacAddress:         diode.String("00:11:22:33:44:55"),
	}
}

// InterfaceExplicit Creates a Interface with fully nested objects and all common fields.
func InterfaceExplicit() *diode.Interface {
	return &diode.Interface{
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
		Name:               diode.String("Example Name"),
		Type:               diode.String("1000base-t"),
		Metadata:           diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:        diode.String("Example description"),
		Label:              diode.String("Example Label"),
		Enabled:            diode.Bool(true),
		Mtu:                diode.Int64(1),
		Speed:              diode.Int64(1),
		Duplex:             diode.String("auto"),
		Wwn:                diode.String("Example Wwn"),
		MgmtOnly:           diode.Bool(true),
		Mode:               diode.String("access"),
		RfRole:             diode.String("ap"),
		RfChannel:          diode.String("2.4g-1-2412-22"),
		PoeMode:            diode.String("pse"),
		PoeType:            diode.String("type1-ieee802.3af"),
		RfChannelFrequency: diode.Float64(1.0),
		RfChannelWidth:     diode.Float64(1.0),
		TxPower:            diode.Int64(1),
		MarkConnected:      diode.Bool(true),
		Channels:           diode.Int64(1),
		MacAddress:         diode.String("00:11:22:33:44:55"),
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
		Parent: &diode.Interface{
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
		Bridge: &diode.Interface{
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
		Lag: &diode.Interface{
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
		PrimaryMacAddress: &diode.MACAddress{
			MacAddress: diode.String("00:11:22:33:44:55"),
			Metadata:   diode.Metadata{"source": "example"},
		},
		UntaggedVlan: &diode.VLAN{
			Vid:      diode.Int64(1),
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		QinqSvlan: &diode.VLAN{
			Vid:      diode.Int64(1),
			Name:     diode.String("Example Name"),
			Status:   diode.String("active"),
			Metadata: diode.Metadata{"source": "example"},
		},
		VlanTranslationPolicy: &diode.VLANTranslationPolicy{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Vrf: &diode.VRF{
			Name:     diode.String("Example Name"),
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
