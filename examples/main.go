package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

func main() {
	client, err := diode.NewClient(
		"grpc://localhost:8080/diode",
		"example-app",
		"0.1.0",
		diode.WithAPIKey("YOUR_API_KEY"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create a device
	deviceEntity := &diode.Device{
		Name: diode.String("Device A"),
		DeviceType: &diode.DeviceType{
			Model: diode.String("Device Type A"),
			Manufacturer: &diode.Manufacturer{
				Name: diode.String("Manufacturer A"),
			},
		},
		Platform: &diode.Platform{
			Name: diode.String("Platform A"),
			Manufacturer: &diode.Manufacturer{
				Name: diode.String("Manufacturer A"),
			},
		},
		Site: &diode.Site{
			Name: diode.String("Site ABC"),
		},
		Role: &diode.Role{
			Name: diode.String("Role ABC"),
			Tags: []*diode.Tag{
				{
					Name: diode.String("tag 1"),
				},
				{
					Name: diode.String("tag 2"),
				},
			},
		},
		Serial:   diode.String("123456"),
		AssetTag: diode.String("123456"),
		Status:   diode.String("active"),
		Comments: diode.String("Lorem ipsum dolor sit amet"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 3"),
			},
		},
	}

	// Create a device type
	deviceTypeEntity := &diode.DeviceType{
		Model: diode.String("Device Type A"),
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Manufacturer A"),
		},
		PartNumber:  diode.String("123456"),
		Description: diode.String("Device Type A description"),
		Comments:    diode.String("Lorem ipsum dolor sit amet"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create an IP address
	ipAddressEntity := &diode.IPAddress{
		Address: diode.String("192.168.0.1/24"),
		AssignedObject: &diode.Interface{
			Name: diode.String("Interface ABC"),
			Device: &diode.Device{
				Name: diode.String("Device ABC"),
				DeviceType: &diode.DeviceType{
					Model: diode.String("Device Type ABC"),
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Manufacturer ABC"),
					},
				},
				Platform: &diode.Platform{
					Name: diode.String("Platform ABC"),
					Manufacturer: &diode.Manufacturer{
						Name: diode.String("Manufacturer ABC"),
					},
				},
				Site: &diode.Site{
					Name: diode.String("Site ABC"),
				},
				Role: &diode.Role{
					Name: diode.String("Role ABC"),
					Tags: []*diode.Tag{
						{
							Name: diode.String("tag 1"),
						},
						{
							Name: diode.String("tag 2"),
						},
					},
				},
			},
		},
		Status:      diode.String("active"),
		Role:        diode.String("anycast"),
		Description: diode.String("IP Address description"),
		Comments:    diode.String("Lorem ipsum dolor sit amet"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create an interface
	interfaceEntity := &diode.Interface{
		Name: diode.String("Interface A"),
		Device: &diode.Device{
			Name: diode.String("Device A"),
			DeviceType: &diode.DeviceType{
				Model: diode.String("Device Type A"),
				Manufacturer: &diode.Manufacturer{
					Name: diode.String("Manufacturer A"),
				},
			},
			Platform: &diode.Platform{
				Name: diode.String("Platform A"),
				Manufacturer: &diode.Manufacturer{
					Name: diode.String("Manufacturer A"),
				},
			},
			Site: &diode.Site{
				Name: diode.String("Site ABC"),
			},
			Role: &diode.Role{
				Name: diode.String("Role ABC"),
				Tags: []*diode.Tag{
					{
						Name: diode.String("tag 1"),
					},
					{
						Name: diode.String("tag 2"),
					},
				},
			},
		},
		Type:        diode.String("virtual"),
		Enabled:     diode.Bool(true),
		Mtu:         diode.Int32(1500),
		MacAddress:  diode.String("00:00:00:00:00:00"),
		Description: diode.String("Interface A description"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create a manufacturer
	manufacturerEntity := &diode.Manufacturer{
		Name:        diode.String("Manufacturer A"),
		Description: diode.String("Manufacturer A description"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create a platform
	platformEntity := &diode.Platform{
		Name: diode.String("Platform A"),
		Manufacturer: &diode.Manufacturer{
			Name: diode.String("Manufacturer A"),
			Tags: []*diode.Tag{
				{
					Name: diode.String("tag 1"),
				},
				{
					Name: diode.String("tag 2"),
				},
			},
		},
		Description: diode.String("Platform A description"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create a prefix
	prefixEntity := &diode.Prefix{
		Prefix: diode.String("192.168.0.0/32"),
		Site: &diode.Site{
			Name: diode.String("Site ABC"),
		},
		Status:       diode.String("active"),
		IsPool:       diode.Bool(true),
		MarkUtilized: diode.Bool(true),
		Description:  diode.String("Prefix description"),
		Comments:     diode.String("Lorem ipsum dolor sit amet"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create a role
	roleEntity := &diode.Role{
		Name:        diode.String("Role A"),
		Color:       diode.String("ffffff"),
		Description: diode.String("Role A description"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 1"),
			},
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	// Create a site
	siteEntity := &diode.Site{
		Name:     diode.String("site A"),
		Comments: diode.String("aaa"),
		Tags: []*diode.Tag{
			{
				Name: diode.String("tag 2"),
			},
		},
	}

	entities := []diode.Entity{
		deviceEntity,
		deviceTypeEntity,
		ipAddressEntity,
		interfaceEntity,
		manufacturerEntity,
		platformEntity,
		prefixEntity,
		roleEntity,
		siteEntity,
	}

	resp, err := client.Ingest(context.Background(), entities)
	if err != nil {
		log.Fatal(err)
	}
	if resp != nil && resp.Errors != nil {
		log.Printf("Errors: %v\n", resp.Errors)
	} else {
		log.Printf("Success\n")
	}

}
