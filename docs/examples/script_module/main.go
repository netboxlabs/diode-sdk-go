// Package main demonstrates ingesting ScriptModule entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "script_module-example"
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
	scriptModule := ScriptModuleMinimal()
	// scriptModule := ScriptModuleExtended()
	// scriptModule := ScriptModuleExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{scriptModule})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("ScriptModule ingested successfully")
	}
}

// ScriptModuleMinimal Creates a ScriptModule with only required fields.
func ScriptModuleMinimal() *diode.ScriptModule {
	return &diode.ScriptModule{
		File:     diode.String("Example File"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// ScriptModuleExtended Creates a ScriptModule with common optional fields.
func ScriptModuleExtended() *diode.ScriptModule {
	return &diode.ScriptModule{
		File:     diode.String("Example File"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value"},
	}
}

// ScriptModuleExplicit Creates a ScriptModule with fully nested objects and all common fields.
func ScriptModuleExplicit() *diode.ScriptModule {
	return &diode.ScriptModule{
		File:     diode.String("Example File"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
	}
}
