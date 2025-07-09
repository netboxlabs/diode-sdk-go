package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

type fileList []string

func (f *fileList) String() string { return strings.Join(*f, ",") }

func (f *fileList) Set(s string) error {
	*f = append(*f, s)
	return nil
}

func main() {
	var files fileList
	flag.Var(&files, "file", "Path to dry run JSON file (may be repeated)")
	target := flag.String("target", "", "Diode gRPC target")
	app := flag.String("app-name", "", "Producer application name")
	version := flag.String("app-version", "", "Producer application version")
	clientID := flag.String("client-id", "", "OAuth2 client ID")
	clientSecret := flag.String("client-secret", "", "OAuth2 client secret")
	flag.Parse()

	// Fall back to environment variables if flags are not provided
	if *clientID == "" {
		*clientID = os.Getenv("DIODE_CLIENT_ID")
	}
	if *clientSecret == "" {
		*clientSecret = os.Getenv("DIODE_CLIENT_SECRET")
	}

	if len(files) == 0 || *target == "" || *app == "" || *version == "" {
		flag.Usage()
		os.Exit(1)
	}

	client, err := diode.NewClient(
		*target,
		*app,
		*version,
		diode.WithClientID(*clientID),
		diode.WithClientSecret(*clientSecret),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Failed to close client: %v", err)
		}
	}()

	ctx := context.Background()
	for _, f := range files {
		entities, err := diode.LoadDryRunEntities(f)
		if err != nil {
			log.Printf("Failed to load %s: %v", f, err)
			continue
		}

		resp, err := client.IngestProto(ctx, entities)
		if err != nil {
			log.Printf("Failed to ingest %s: %v", f, err)
			continue
		}

		if resp.GetErrors() != nil {
			log.Printf("Errors while ingesting %s: %v", f, resp.GetErrors())
			continue
		}

		log.Printf("Ingested %d entities from %s successfully", len(entities), f)
	}
}
