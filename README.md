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
* `DIODE_MAX_AUTH_RETRIES` - Maximum attempts for OAuth2 token fetch and gRPC re-authentication on `Unauthenticated` (default: `3`). Token fetch retries with exponential backoff on `429`, `500`, `502`, and `503`, honouring `Retry-After` when present on `429`/`503`.
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

See all [entities examples](./docs/README.md).

### Using Metadata

Entities support attaching custom metadata as key-value pairs. Metadata can be used to store additional context, tracking information, or custom attributes that don't fit into the standard NetBox fields.

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

	// Create a device with metadata
	// Note: Both the device and its nested site can have its own metadata
	deviceEntity := &diode.Device{
		Name: diode.String("Device A"),
		Site: &diode.Site{
			Name: diode.String("Site ABC"),
			Metadata: diode.Metadata{
				"site_region":      "us-west",
				"site_cost_center": "CC-001",
			},
		},
		DeviceType: &diode.DeviceType{
			Model: diode.String("Device Type A"),
		},
		Role: &diode.DeviceRole{
			Name: diode.String("Role ABC"),
		},
		Metadata: diode.Metadata{
			"source":        "network_discovery",
			"discovered_at": "2024-01-15T10:30:00Z",
			"import_batch":  "batch-123",
			"priority":      1,
			"verified":      true,
		},
	}

	// Create an IP address with metadata
	ipEntity := &diode.IPAddress{
		Address: diode.String("192.168.1.10/24"),
		Status:  diode.String("active"),
		Metadata: diode.Metadata{
			"last_scan":     "2024-01-15T12:00:00Z",
			"scan_id":       "scan-456",
			"response_time": 23.5,
			"reachable":     true,
			"owner_team":    "network-ops",
		},
	}

	// Create a site with metadata
	siteEntity := &diode.Site{
		Name:   diode.String("Data Center 1"),
		Status: diode.String("active"),
		Metadata: diode.Metadata{
			"region":       "us-west",
			"cost_center":  "CC-001",
			"capacity":     500,
			"is_primary":   true,
			"contact_email": "dc1-ops@example.com",
		},
	}

	entities := []diode.Entity{
		deviceEntity,
		ipEntity,
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
```

#### Adding request-level metadata

In addition to entity-level metadata, you can attach metadata to the entire ingestion request using `WithIngestMetadata`. This is useful for tracking information about the ingestion batch itself, such as the data source, batch ID, or processing context.

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

	// Create entities
	entities := []diode.Entity{
		&diode.Device{
			Name: diode.String("Device A"),
			Site: &diode.Site{
				Name: diode.String("Site ABC"),
			},
		},
		&diode.Device{
			Name: diode.String("Device B"),
			Site: &diode.Site{
				Name: diode.String("Site XYZ"),
			},
		},
	}

	// Add request-level metadata to track the ingestion batch
	resp, err := client.Ingest(
		context.Background(),
		entities,
		diode.WithIngestMetadata(diode.Metadata{
			"batch_id":      "import-2024-01-15",
			"source_system": "network_scanner",
			"import_type":   "automated",
			"record_count":  2,
			"validated":     true,
		}),
	)
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

Request-level metadata is included in the `IngestRequest` and can be useful for:
- Tracking data sources and ingestion pipelines
- Correlating entities within a batch
- Debugging and auditing data imports
- Adding contextual information for downstream processing

### Message chunking

When ingesting large datasets, you may encounter gRPC message size limits (typically 4 MB). The SDK provides automatic message chunking to split large entity lists into size-appropriate chunks that stay within these limits.

#### Automatic chunking

Enable chunking by using `WithChunking()` when calling `Ingest`:

```go
package main

import (
	"context"
	"fmt"
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

	// Create a large number of entities
	entities := make([]diode.Entity, 0)
	for i := 0; i < 10000; i++ {
		entities = append(entities, &diode.Device{
			Name: diode.String(fmt.Sprintf("Device %d", i)),
			Site: &diode.Site{
				Name: diode.String("Site ABC"),
			},
			DeviceType: &diode.DeviceType{
				Model: diode.String("Device Type A"),
			},
			Role: &diode.DeviceRole{
				Name: diode.String("Role ABC"),
			},
		})
	}

	// Use chunking with default 3.0 MB chunk size
	resp, err := client.Ingest(
		context.Background(),
		entities,
		diode.WithChunking(0), // 0 = use default
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Success\n")
}
```

#### Custom chunk size

You can specify a custom chunk size in megabytes:

```go
// Use 3.5 MB chunks instead of the default 3.0 MB
resp, err := client.Ingest(
	context.Background(),
	entities,
	diode.WithChunking(3.5),
)
```

#### Manual chunking

For more control, you can manually chunk entities using the `CreateMessageChunks` function:

```go
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
	pb "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
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

	// Create entities
	entities := []diode.Entity{
		// ... many entities
	}

	// Convert to proto entities
	protoEntities := make([]*pb.Entity, 0)
	for _, entity := range entities {
		protoEntities = append(protoEntities, entity.ConvertToProtoEntity())
	}

	// Manually chunk with custom size (3.5 MB)
	chunks := diode.CreateMessageChunks(protoEntities, 3.5)

	log.Printf("Split %d entities into %d chunks\n", len(protoEntities), len(chunks))

	// Ingest each chunk
	for i, chunk := range chunks {
		log.Printf("Ingesting chunk %d of %d (%d entities)\n", i+1, len(chunks), len(chunk))
		resp, err := client.IngestProto(context.Background(), chunk)
		if err != nil {
			log.Fatalf("Failed to ingest chunk %d: %v\n", i+1, err)
		}
		if resp != nil && resp.Errors != nil {
			log.Printf("Chunk %d errors: %v\n", i+1, resp.Errors)
		}
	}
	log.Printf("Successfully ingested all chunks\n")
}
```

#### Estimating message size

You can estimate the size of your entities before chunking:

```go
import pb "github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"

protoEntities := make([]*pb.Entity, 0)
for _, entity := range entities {
	protoEntities = append(protoEntities, entity.ConvertToProtoEntity())
}

sizeBytes := diode.EstimateMessageSize(protoEntities)
sizeMB := float64(sizeBytes) / (1024 * 1024)
log.Printf("Estimated message size: %.2f MB\n", sizeMB)

if sizeMB > 3.0 {
	log.Printf("Message exceeds 3 MB, chunking recommended\n")
}
```

#### How chunking works

The chunking algorithm uses greedy bin-packing to efficiently group entities:

1. It accumulates entities until adding the next one would exceed the size limit
2. When the limit would be exceeded, it starts a new chunk
3. Each chunk includes the base overhead of an `IngestRequest` protobuf message
4. Entity order is preserved across chunks

The default chunk size of 3.0 MB provides a safe margin below the gRPC 4 MB message size limit, accounting for protobuf serialization overhead and network protocol overhead.

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

#### Adding request-level metadata to dry run output

You can include request-level metadata in the dry run output using `WithIngestMetadata`. This metadata will be included in the JSON output file as part of the `IngestRequest`:

```go
client, err := diode.NewDryRunClient("example-app", "/tmp")
if err != nil {
        log.Fatal(err)
}

// Add request-level metadata
_, _ = client.Ingest(
        context.Background(),
        []diode.Entity{
                &diode.Device{Name: diode.String("Device A")},
        },
        diode.WithIngestMetadata(diode.Metadata{
                "batch_id":   "import-2024-01",
                "source":     "csv_import",
                "validated":  true,
                "record_count": 150,
        }),
)
_ = client.Close()
```

The resulting JSON file will include the metadata in the `IngestRequest`, making it visible when reviewing the dry run output.

#### Loading and replaying dry run data

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

`OTLPClient` converts ingestion entities into OpenTelemetry log records and exports them to an OTLP collector over gRPC. This is useful when a collector receives log data and forwards it to Diode.

```go
client, err := diode.NewOTLPClient(
        "grpc://localhost:4317",
        "otlp-producer",
        "0.0.1",
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

Each entity is serialised with protobuf field names and emitted as a log record that includes SDK and producer metadata via resource attributes so downstream collectors can enrich and forward the payload. The client raises an `OTLPClientError` when the export fails. TLS behaviour honours the existing `DIODE_SKIP_TLS_VERIFY` and `DIODE_CERT_FILE` environment variables, and the export timeout can be customised via `diode.WithOTLPTimeout`.

#### Adding request-level metadata as OTLP resource attributes

You can add request-level metadata to OTLP exports using `WithIngestMetadata`. This metadata is automatically mapped to OTLP resource attributes with a `diode.metadata.` prefix:

```go
client, err := diode.NewOTLPClient(
        "grpc://localhost:4317",
        "otlp-producer",
        "0.0.1",
)
if err != nil {
        log.Fatal(err)
}
defer client.Close()

// Add request-level metadata
_, err = client.Ingest(
        context.Background(),
        []diode.Entity{
                &diode.Site{Name: diode.String("Site 1")},
        },
        diode.WithIngestMetadata(diode.Metadata{
                "environment": "production",
                "deployment":  "us-west-2",
                "version":     "1.2.3",
                "priority":    5,
        }),
)
if err != nil {
        log.Fatal(err)
}
```

The resulting OTLP log records will include resource attributes like:
- `diode.metadata.environment="production"`
- `diode.metadata.deployment="us-west-2"`
- `diode.metadata.version="1.2.3"`
- `diode.metadata.priority=5` (as integer)

These attributes are added alongside standard OTLP resource attributes (`service.name`, `service.version`, `diode.stream`, etc.), allowing downstream collectors and observability platforms to filter, route, and enrich the data based on this metadata.

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
