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
		Name:        diode.String("Example Name"),
		Version:     diode.Int64(1),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// IKEPolicyExplicit Creates a IKEPolicy with fully nested objects and all common fields.
func IKEPolicyExplicit() *diode.IKEPolicy {
	return &diode.IKEPolicy{
		Name:        diode.String("Example Name"),
		Version:     diode.Int64(1),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
