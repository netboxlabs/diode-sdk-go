// Package main demonstrates ingesting FHRPGroup entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "fhrp_group-example"
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
	fhrpGroup := FHRPGroupMinimal()
	// fhrpGroup := FHRPGroupExtended()
	// fhrpGroup := FHRPGroupExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{fhrpGroup})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("FHRPGroup ingested successfully")
	}
}

// FHRPGroupMinimal Creates a FHRPGroup with only required fields.
func FHRPGroupMinimal() *diode.FHRPGroup {
	return &diode.FHRPGroup{
		Protocol: diode.String("carp"),
		GroupId:  diode.Int64(1),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// FHRPGroupExtended Creates a FHRPGroup with common optional fields.
func FHRPGroupExtended() *diode.FHRPGroup {
	return &diode.FHRPGroup{
		Protocol:    diode.String("carp"),
		GroupId:     diode.Int64(1),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description: diode.String("Example description"),
		Name:        diode.String("Example Name"),
		AuthType:    diode.String("md5"),
		AuthKey:     diode.String("Example AuthKey"),
		Comments:    diode.String("Example comments"),
	}
}

// FHRPGroupExplicit Creates a FHRPGroup with fully nested objects and all common fields.
func FHRPGroupExplicit() *diode.FHRPGroup {
	return &diode.FHRPGroup{
		Protocol:    diode.String("carp"),
		GroupId:     diode.Int64(1),
		Metadata:    diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Name:        diode.String("Example Name"),
		AuthType:    diode.String("md5"),
		AuthKey:     diode.String("Example AuthKey"),
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
