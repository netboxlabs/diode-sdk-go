// Package main demonstrates ingesting IPSecPolicy entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ip_sec_policy-example"
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
	ipSecPolicy := IPSecPolicyMinimal()
	// ipSecPolicy := IPSecPolicyExtended()
	// ipSecPolicy := IPSecPolicyExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ipSecPolicy})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IPSecPolicy ingested successfully")
	}
}

// IPSecPolicyMinimal Creates a IPSecPolicy with only required fields.
func IPSecPolicyMinimal() *diode.IPSecPolicy {
	return &diode.IPSecPolicy{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// IPSecPolicyExtended Creates a IPSecPolicy with common optional fields.
func IPSecPolicyExtended() *diode.IPSecPolicy {
	return &diode.IPSecPolicy{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		PfsGroup:    diode.Int64(1),
		Comments:    diode.String("Example comments"),
	}
}

// IPSecPolicyExplicit Creates a IPSecPolicy with fully nested objects and all common fields.
func IPSecPolicyExplicit() *diode.IPSecPolicy {
	return &diode.IPSecPolicy{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		PfsGroup:    diode.Int64(1),
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
