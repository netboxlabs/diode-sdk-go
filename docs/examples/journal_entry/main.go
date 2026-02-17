// Package main demonstrates ingesting JournalEntry entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "journal_entry-example"
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
	journalEntry := JournalEntryMinimal()
	// journalEntry := JournalEntryExtended()
	// journalEntry := JournalEntryExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{journalEntry})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("JournalEntry ingested successfully")
	}
}

// JournalEntryMinimal Creates a JournalEntry with only required fields.
func JournalEntryMinimal() *diode.JournalEntry {
	return &diode.JournalEntry{
		Comments: diode.String("Example comments"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// JournalEntryExtended Creates a JournalEntry with common optional fields.
func JournalEntryExtended() *diode.JournalEntry {
	return &diode.JournalEntry{
		Comments: diode.String("Example comments"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Kind:     diode.String("danger"),
	}
}

// JournalEntryExplicit Creates a JournalEntry with fully nested objects and all common fields.
func JournalEntryExplicit() *diode.JournalEntry {
	return &diode.JournalEntry{
		Comments: diode.String("Example comments"),
		Metadata: diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Kind:     diode.String("danger"),
		Tags:     []*diode.Tag{{Name: diode.String("production")}},
	}
}
