# Diode SDK Go

Diode SDK Go is a Go library for interacting with the Diode ingestion service utilizing gRPC.

Diode is a new [NetBox](https://netboxlabs.com/oss/netbox/) ingestion service that greatly simplifies and enhances the
process to add and update network data
in NetBox, ensuring your network source of truth is always accurate and can be trusted to power your network automation
pipelines.

More information about Diode can be found
at [https://netboxlabs.com/blog/introducing-diode-streamlining-data-ingestion-in-netbox/](https://netboxlabs.com/blog/introducing-diode-streamlining-data-ingestion-in-netbox/).

## Prerequisites
- Go 1.24 or later installed

## Installation

```bash
go get github.com/netboxlabs/diode-sdk-go
```

## Usage

### Environment variables

* `DIODE_SDK_LOG_LEVEL` - Log level for the SDK (default: `INFO`)
* `DIODE_CLIENT_ID` - Client ID for OAuth2 authentication
* `DIODE_CLIENT_SECRET` - Client Secret for OAuth2 authentication
* `DIODE_CERT_FILE` - Path to custom certificate file for TLS connections
* `DIODE_SKIP_TLS_VERIFY` - Skip TLS verification (default: `false`)
* `DIODE_DRY_RUN_OUTPUT_DIR` - Directory to write dry run output files when using `DryRunClient`

### Example

* `target` should be the address of the Diode service.
  * Insecure connections: `grpc://localhost:8080/diode` or `http://localhost:8080/diode`
  * Secure connections: `grpcs://example.com` or `https://example.com`

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
		Role: &diode.DeviceRole{
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

### TLS verification and certificates

TLS verification is controlled by the target URL scheme:
- **Secure schemes** (`grpcs://`, `https://`): TLS verification enabled  
- **Insecure schemes** (`grpc://`, `http://`): TLS verification disabled

```go
// TLS verification enabled (uses system certificates)
client, err := diode.NewClient("grpcs://example.com", ...)

// TLS verification disabled  
client, err := diode.NewClient("grpc://example.com", ...)
```

#### Using custom certificates

```go
// Via constructor parameter
client, err := diode.NewClient(
	"grpcs://example.com", 
	"example-app", "0.1.0",
	diode.WithCertFile("/path/to/cert.pem"),
)

// Or via environment variable
export DIODE_CERT_FILE=/path/to/cert.pem
```

#### Disabling TLS verification

```bash
export DIODE_SKIP_TLS_VERIFY=true
```

#### For legacy certificates (CN-only, no SANs)

```go
client, err := diode.NewClient(
	"grpcs://example.com",
	"example-app", "0.1.0", 
	diode.WithCertFile("/path/to/cert.pem"),
	diode.WithSkipTLSVerify(),
)
```

### Dry run client

Use a `DryRunClient` to inspect what would be sent to Diode without actually sending any data. When a directory is provided a new JSON file is created for each ingest call.

```go
// Write ingest payload to a timestamped file in /tmp
client, err := diode.NewDryRunClient("example-app", "/tmp")
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

_, err = realClient.IngestProto(context.Background(), protoEntities)
if err != nil {
        log.Fatal(err)
}
```

### OTLP client

`DiodeOTLPClient` converts ingestion entities into OpenTelemetry log records and exports them to an OTLP collector over gRPC. This is useful when a collector receives log data and forwards it to Diode.

```go
client, err := diode.NewDiodeOTLPClient(
        "grpc://localhost:4317",
        "otlp-producer",
        "0.0.1",
        diode.WithOtlpMetadata(map[string]string{"authorization": "Bearer token"}),
)
if err != nil {
        log.Fatal(err)
}
defer client.Close()

_, err = client.Ingest(context.Background(), []diode.Entity{
        &diode.Site{Name: diode.String("Site 1")},
})
if err != nil {
        log.Fatal(err)
}
```

Each entity is serialised with protobuf field names and emitted as a log record that includes SDK and producer metadata via resource attributes so downstream collectors can enrich and forward the payload. The client raises an `OTLPClientError` when the export fails. TLS behaviour honours the existing `DIODE_SKIP_TLS_VERIFY` and `DIODE_CERT_FILE` environment variables, and the export timeout can be customised via `diode.WithOtlpTimeout`.

### CLI to replay dry-run files

A small helper binary is included to ingest JSON files created by the
`DryRunClient` and send them to a running Diode service.

Install the helper using `go install`:

```bash
go install github.com/netboxlabs/diode-sdk-go/cmd/diode-replay-dryrun@latest
```

This installs the command separately from the SDK library obtained with
`go get`.

Run it by providing one or more JSON files and connection details. Use `-file`
multiple times to ingest several dry-run files in a single request:

```bash
diode-replay-dryrun \
  -file /tmp/example-app_1750106879725947344.json \
  -file /tmp/other.json \
  -target grpc://localhost:8080/diode \
  -app-name example-app \
  -app-version 0.1.0 \
  -client-id YOUR_CLIENT_ID \
  -client-secret YOUR_CLIENT_SECRET
```

The flags `-file`, `-target`, `-app-name`, and `-app-version` are required. You may
repeat `-file` to specify multiple files. OAuth2
credentials can be supplied using `-client-id` and `-client-secret` or the
`DIODE_CLIENT_ID` and `DIODE_CLIENT_SECRET` environment variables.

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
