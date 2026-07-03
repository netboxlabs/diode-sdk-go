// Package main demonstrates ingesting Tunnel entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "tunnel-example"
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
	tunnel := TunnelMinimal()
	// tunnel := TunnelExtended()
	// tunnel := TunnelExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{tunnel})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("Tunnel ingested successfully")
	}
}

// TunnelMinimal Creates a Tunnel with only required fields.
func TunnelMinimal() *diode.Tunnel {
	return &diode.Tunnel{
		Name:          diode.String("Example Name"),
		Status:        diode.String("active"),
		Encapsulation: diode.String("gre"),
		Metadata:      diode.Metadata{"source": "example"},
	}
}

// TunnelExtended Creates a Tunnel with common optional fields.
func TunnelExtended() *diode.Tunnel {
	return &diode.Tunnel{
		Name:          diode.String("Example Name"),
		Status:        diode.String("active"),
		Encapsulation: diode.String("gre"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:   diode.String("Example description"),
		TunnelId:      diode.Int64(1),
		Comments:      diode.String("Example comments"),
	}
}

// TunnelExplicit Creates a Tunnel with fully nested objects and all common fields.
func TunnelExplicit() *diode.Tunnel {
	return &diode.Tunnel{
		Name:          diode.String("Example Name"),
		Status:        diode.String("active"),
		Encapsulation: diode.String("gre"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		TunnelId:      diode.Int64(1),
		Group: &diode.TunnelGroup{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		IpsecProfile: &diode.IPSecProfile{
			Name: diode.String("Example Name"),
			Mode: diode.String("ah"),
			IkePolicy: &diode.IKEPolicy{
				Name:     diode.String("Example Name"),
				Version:  diode.Int64(1),
				Metadata: diode.Metadata{"source": "example"},
			},
			IpsecPolicy: &diode.IPSecPolicy{
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
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
