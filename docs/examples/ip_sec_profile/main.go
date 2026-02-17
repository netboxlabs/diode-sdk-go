// Package main demonstrates ingesting IPSecProfile entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ip_sec_profile-example"
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
	ipSecProfile := IPSecProfileMinimal()
	// ipSecProfile := IPSecProfileExtended()
	// ipSecProfile := IPSecProfileExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ipSecProfile})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IPSecProfile ingested successfully")
	}
}

// IPSecProfileMinimal Creates a IPSecProfile with only required fields.
func IPSecProfileMinimal() *diode.IPSecProfile {
	return &diode.IPSecProfile{
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
	}
}

// IPSecProfileExtended Creates a IPSecProfile with common optional fields.
func IPSecProfileExtended() *diode.IPSecProfile {
	return &diode.IPSecProfile{
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
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// IPSecProfileExplicit Creates a IPSecProfile with fully nested objects and all common fields.
func IPSecProfileExplicit() *diode.IPSecProfile {
	return &diode.IPSecProfile{
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
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
