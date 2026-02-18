// Package main demonstrates ingesting IPRange entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ip_range-example"
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
	ipRange := IPRangeMinimal()
	// ipRange := IPRangeExtended()
	// ipRange := IPRangeExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ipRange})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IPRange ingested successfully")
	}
}

// IPRangeMinimal Creates a IPRange with only required fields.
func IPRangeMinimal() *diode.IPRange {
	return &diode.IPRange{
		StartAddress: diode.String("Example StartAddress"),
		EndAddress:   diode.String("Example EndAddress"),
		Metadata:     diode.Metadata{"source": "example"},
	}
}

// IPRangeExtended Creates a IPRange with common optional fields.
func IPRangeExtended() *diode.IPRange {
	return &diode.IPRange{
		StartAddress:  diode.String("Example StartAddress"),
		EndAddress:    diode.String("Example EndAddress"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Status:        diode.String("active"),
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		MarkUtilized:  diode.Bool(true),
		MarkPopulated: diode.Bool(true),
	}
}

// IPRangeExplicit Creates a IPRange with fully nested objects and all common fields.
func IPRangeExplicit() *diode.IPRange {
	return &diode.IPRange{
		StartAddress:  diode.String("Example StartAddress"),
		EndAddress:    diode.String("Example EndAddress"),
		Metadata:      diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Status:        diode.String("active"),
		Description:   diode.String("Example description"),
		Comments:      diode.String("Example comments"),
		MarkUtilized:  diode.Bool(true),
		MarkPopulated: diode.Bool(true),
		Vrf: &diode.VRF{
			Name:     diode.String("Example Name"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Tenant: &diode.Tenant{
			Name:     diode.String("Example Name"),
			Slug:     diode.String("example-slug"),
			Metadata: diode.Metadata{"source": "example"},
		},
		Role: &diode.Role{
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
