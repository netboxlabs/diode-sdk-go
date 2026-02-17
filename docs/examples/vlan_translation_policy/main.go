// Package main demonstrates ingesting VLANTranslationPolicy entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vlan_translation_policy-example"
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
	vlanTranslationPolicy := VLANTranslationPolicyMinimal()
	// vlanTranslationPolicy := VLANTranslationPolicyExtended()
	// vlanTranslationPolicy := VLANTranslationPolicyExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vlanTranslationPolicy})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VLANTranslationPolicy ingested successfully")
	}
}

// VLANTranslationPolicyMinimal Creates a VLANTranslationPolicy with only required fields.
func VLANTranslationPolicyMinimal() *diode.VLANTranslationPolicy {
	return &diode.VLANTranslationPolicy{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// VLANTranslationPolicyExtended Creates a VLANTranslationPolicy with common optional fields.
func VLANTranslationPolicyExtended() *diode.VLANTranslationPolicy {
	return &diode.VLANTranslationPolicy{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// VLANTranslationPolicyExplicit Creates a VLANTranslationPolicy with fully nested objects and all common fields.
func VLANTranslationPolicyExplicit() *diode.VLANTranslationPolicy {
	return &diode.VLANTranslationPolicy{
		Name:        diode.String("Example Name"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
	}
}
