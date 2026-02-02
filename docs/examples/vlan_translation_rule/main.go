// Package main demonstrates ingesting VLANTranslationRule entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "vlan_translation_rule-example"
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
	vlanTranslationRule := VLANTranslationRuleMinimal()
	// vlanTranslationRule := VLANTranslationRuleExtended()
	// vlanTranslationRule := VLANTranslationRuleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{vlanTranslationRule})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("VLANTranslationRule ingested successfully")
	}
}

// VLANTranslationRuleMinimal Creates a VLANTranslationRule with only required fields.
func VLANTranslationRuleMinimal() *diode.VLANTranslationRule {
	return &diode.VLANTranslationRule{
		Policy: &diode.VLANTranslationPolicy{
			Name: diode.String("Example Name"),
		},
		LocalVid:  diode.Int64(1),
		RemoteVid: diode.Int64(1),
	}
}

// VLANTranslationRuleExtended Creates a VLANTranslationRule with common optional fields.
func VLANTranslationRuleExtended() *diode.VLANTranslationRule {
	return &diode.VLANTranslationRule{
		Policy: &diode.VLANTranslationPolicy{
			Name: diode.String("Example Name"),
		},
		LocalVid:    diode.Int64(1),
		RemoteVid:   diode.Int64(1),
		Description: diode.String("Example description"),
	}
}

// VLANTranslationRuleExplicit Creates a VLANTranslationRule with fully nested objects and all common fields.
func VLANTranslationRuleExplicit() *diode.VLANTranslationRule {
	return &diode.VLANTranslationRule{
		Policy: &diode.VLANTranslationPolicy{
			Name: diode.String("Example Name"),
		},
		LocalVid:    diode.Int64(1),
		RemoteVid:   diode.Int64(1),
		Description: diode.String("Example description"),
	}
}
