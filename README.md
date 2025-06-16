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

* `DIODE_SDK_LOG_LEVEL` - Log level for the SDK (default: `INFO`)
* `DIODE_CLIENT_ID` - Client ID for OAuth2 authentication
* `DIODE_CLIENT_SECRET` - Client Secret for OAuth2 authentication
* `DIODE_DRY_RUN_OUTPUT_DIR` - Directory to write dry run output files when using `DryRunClient`

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
		diode.WithClientID("YOUR_CLIENT_ID"),
		diode.WithClientSecret("YOUR_CLIENT_SECRET"),
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

### Dry run client

Use a `DryRunClient` to inspect what would be sent to Diode without actually sending any data. When a directory is provided a new JSON file is created for each ingest call.

```go
// Write ingest payload to a timestamped file in /tmp
client, err := diode.NewDryRunClient("/tmp", "example-app")
if err != nil {
        log.Fatal(err)
}
_, _ = client.Ingest(context.Background(), []diode.Entity{
        &diode.Device{Name: diode.String("Device A")},
})
_ = client.Close()
```

Loaded entities can later be ingested using a real client:

```go
protoEntities, err := diode.LoadDryRunEntities("/tmp/example-app_1750106879725947344.json")
if err != nil {
        log.Fatal(err)
}

diodeEntities := make([]diode.Entity, len(protoEntities))
for i, e := range protoEntities {
        diodeEntities[i] = diode.ProtoEntity{PB: e}
}

realClient, err := diode.NewClient(
        "grpc://localhost:8080/diode",
        "example-app",
        "0.1.0",
        diode.WithClientID("YOUR_CLIENT_ID"),
        diode.WithClientSecret("YOUR_CLIENT_SECRET"),
)
if err != nil {
        log.Fatal(err)
}

_, err = realClient.Ingest(context.Background(), diodeEntities)
if err != nil {
        log.Fatal(err)
}
```

## Supported entities (object types)

* ASN
* ASN Range
* Aggregate
* Circuit
* Circuit Group
* Circuit Group Assignment
* Circuit Termination
* Circuit Type
* Cluster
* Cluster Group
* Cluster Type
* Console Port
* Console Server Port
* Contact
* Contact Assignment
* Contact Group
* Contact Role
* Device
* Device Bay
* Device Role
* Device Type
* FHRP Group
* FHRP Group Assignment
* Front Port
* IKE Policy
* IKE Proposal
* IP Address
* IP Range
* IP Sec Policy
* IP Sec Profile
* IP Sec Proposal
* Interface
* Inventory Item
* Inventory Item Role
* L2VPN
* L2VPN Termination
* Location
* MAC Address
* Manufacturer
* Module
* Module Bay
* Module Type
* Platform
* Power Feed
* Power Outlet
* Power Panel
* Power Port
* Prefix
* Provider
* Provider Account
* Provider Network
* RIR
* Rack
* Rack Role
* Rack Type
* Rear Port
* Region
* Role
* Route Target
* Service
* Site
* Site Group
* Tag
* Tenant
* Tenant Group
* Tunnel
* Tunnel Group
* Tunnel Termination
* VLAN
* VLAN Group
* VLAN Translation Policy
* VLAN Translation Rule
* VM Interface
* VRF
* Virtual Chassis
* Virtual Circuit
* Virtual Circuit Termination
* Virtual Circuit Type
* Virtual Device Context
* Virtual Disk
* Virtual Machine
* Wireless Lan
* Wireless Lan Group
* Wireless Link

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
