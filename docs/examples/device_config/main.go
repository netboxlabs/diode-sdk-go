// Package main demonstrates ingesting DeviceConfig entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "device_config-example"
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
	deviceConfig := DeviceConfigMinimal()
	// deviceConfig := DeviceConfigExtended()
	// deviceConfig := DeviceConfigExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{deviceConfig})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("DeviceConfig ingested successfully")
	}
}

// DeviceConfigMinimal Creates a DeviceConfig with only required fields.
func DeviceConfigMinimal() *diode.DeviceConfig {
	return &diode.DeviceConfig{
		Running:  []byte("hostname router-01\ninterface eth0\n ip address 192.0.2.1/24"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// DeviceConfigExtended Creates a DeviceConfig with common optional fields.
func DeviceConfigExtended() *diode.DeviceConfig {
	return &diode.DeviceConfig{
		Running:  []byte("hostname router-01\ninterface eth0\n ip address 192.0.2.1/24"),
		Startup:  []byte("hostname router-01\nboot system flash:image.bin"),
		Metadata: diode.Metadata{"source": "network-discovery", "device": "router-01"},
	}
}

// DeviceConfigExplicit Creates a DeviceConfig with fully nested objects and all common fields.
func DeviceConfigExplicit() *diode.DeviceConfig {
	return &diode.DeviceConfig{
		Running:   []byte("hostname router-01\ninterface eth0\n ip address 192.0.2.1/24"),
		Startup:   []byte("hostname router-01\nboot system flash:image.bin"),
		Candidate: []byte("interface eth0\n description WAN uplink\n ip address 192.0.2.1/24"),
		Metadata:  diode.Metadata{"source": "network-discovery", "device": "router-01", "collected_at": "2024-01-15T10:30:00Z"},
	}
}
