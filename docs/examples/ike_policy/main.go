// Package main demonstrates ingesting IKEPolicy entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ike_policy-example"
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
	ikePolicy := IKEPolicyMinimal()
	// ikePolicy := IKEPolicyExtended()
	// ikePolicy := IKEPolicyExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ikePolicy})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IKEPolicy ingested successfully")
	}
}

// IKEPolicyMinimal Creates a IKEPolicy with only required fields.
func IKEPolicyMinimal() *diode.IKEPolicy {
	return &diode.IKEPolicy{
		Name:     diode.String("Example Name"),
		Version:  diode.Int64(1),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// IKEPolicyExtended Creates a IKEPolicy with common optional fields.
func IKEPolicyExtended() *diode.IKEPolicy {
	return &diode.IKEPolicy{
		Name:         diode.String("Example Name"),
		Version:      diode.Int64(1),
		Metadata:     diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:  diode.String("Example description"),
		Mode:         diode.String("aggressive"),
		PresharedKey: diode.String("Example PresharedKey"),
		Comments:     diode.String("Example comments"),
	}
}

// IKEPolicyExplicit Creates a IKEPolicy with fully nested objects and all common fields.
func IKEPolicyExplicit() *diode.IKEPolicy {
	return &diode.IKEPolicy{
		Name:         diode.String("Example Name"),
		Version:      diode.Int64(1),
		Metadata:     diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:  diode.String("Example description"),
		Comments:     diode.String("Example comments"),
		Mode:         diode.String("aggressive"),
		PresharedKey: diode.String("Example PresharedKey"),
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
