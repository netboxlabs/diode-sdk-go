# Diode SDK Go

Diode SDK Go is a Go library for interacting with the Diode ingestion service utilizing gRPC.

Diode is a new [NetBox](https://netboxlabs.com/oss/netbox/) ingestion service that greatly simplifies and enhances the
process to add and update network data
in NetBox, ensuring your network source of truth is always accurate and can be trusted to power your network automation
pipelines.

More information about Diode can be found
at [https://netboxlabs.com/blog/introducing-diode-streamlining-data-ingestion-in-netbox/](https://netboxlabs.com/blog/introducing-diode-streamlining-data-ingestion-in-netbox/).

## Installation

```bash
go get github.com/netboxlabs/diode-sdk-go
```

## Usage

### Environment variables

* `DIODE_API_KEY` - API key for the Diode service
* `DIODE_SDK_LOG_LEVEL` - Log level for the SDK (default: `INFO`)

### Example

* `target` should be the address of the Diode service, e.g. `grpc://localhost:8080/diode` for insecure connection
  or `grpcs://example.com` for secure connection.

```go
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

	entities := []diode.Entity{
		deviceEntity,
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
```

See all [examples](./examples/main.go) for reference.

## Supported entities (object types)

* Device
* Device Type
* IP Address
* Interface
* Manufacturer
* Platform
* Prefix
* Role
* Site
* Cluster Group
* Cluster Type
* Cluster
* Virtual Machine
* Virtual Machine Interface
* Virtual Disk

#### Linting

```shell
make lint
```

#### Testing

```shell
make test
```

## License

Distributed under the Apache 2.0 License. See [LICENSE.txt](./LICENSE.txt) for more information.
