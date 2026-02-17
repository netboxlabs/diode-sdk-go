// Package main demonstrates ingesting MACAddress entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "mac_address-example"
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
	macAddress := MACAddressMinimal()
	// macAddress := MACAddressExtended()
	// macAddress := MACAddressExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{macAddress})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("MACAddress ingested successfully")
	}
}

// MACAddressMinimal Creates a MACAddress with only required fields.
func MACAddressMinimal() *diode.MACAddress {
	return &diode.MACAddress{
		MacAddress: diode.String("00:11:22:33:44:55"),
		Metadata:   diode.Metadata{"source": "example"},
	}
}

// MACAddressExtended Creates a MACAddress with common optional fields.
func MACAddressExtended() *diode.MACAddress {
	return &diode.MACAddress{
		MacAddress:  diode.String("00:11:22:33:44:55"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
	}
}

// MACAddressExplicit Creates a MACAddress with fully nested objects and all common fields.
func MACAddressExplicit() *diode.MACAddress {
	return &diode.MACAddress{
		MacAddress:  diode.String("00:11:22:33:44:55"),
		Metadata:    diode.Metadata{"source": "example"},
		Description: diode.String("Example description"),
		Comments:    diode.String("Example comments"),
		Tags:        []*diode.Tag{{Name: diode.String("production")}},
	}
}
