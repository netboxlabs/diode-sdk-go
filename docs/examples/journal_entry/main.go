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
	}
}

// JournalEntryExtended Creates a JournalEntry with common optional fields.
func JournalEntryExtended() *diode.JournalEntry {
	return &diode.JournalEntry{
		Comments: diode.String("Example comments"),
	}
}

// JournalEntryExplicit Creates a JournalEntry with fully nested objects and all common fields.
func JournalEntryExplicit() *diode.JournalEntry {
	return &diode.JournalEntry{
		Comments: diode.String("Example comments"),
		Tags:     []*diode.Tag{{Name: diode.String("production")}},
	}
}
