package main

import (
	"context"
	"flag"
	"fmt"
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
	flag.Var(&files, "file", "Dry-run JSON file to ingest (may be repeated)")
	target := flag.String("target", "", "gRPC target of the Diode server, e.g. grpc://localhost:8080/diode")
	app := flag.String("app-name", "", "Application name used when ingesting the dry-run messages")
	version := flag.String("app-version", "", "Application version used when ingesting the dry-run messages")
	clientID := flag.String("client-id", "", "OAuth2 client ID. Defaults to the DIODE_CLIENT_ID environment variable if not provided")
	clientSecret := flag.String("client-secret", "", "OAuth2 client secret. Defaults to the DIODE_CLIENT_SECRET environment variable if not provided")
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
		fmt.Fprintf(os.Stderr, "Error: the following arguments are required: -target, -app-name, -app-version, -file\n")
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
		fmt.Fprintf(os.Stderr, "Failed to create Diode client: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close Diode client: %v\n", err)
			os.Exit(1)
		}
	}()

	ctx := context.Background()
	for _, f := range files {
		entities, err := diode.LoadDryRunEntities(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load entities from %s: %v\n", f, err)
			continue
		}

		resp, err := client.IngestProto(ctx, entities)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to ingest %s: %v\n", f, err)
			continue
		}

		if resp.GetErrors() != nil {
			fmt.Fprintf(os.Stderr, "Errors while ingesting %s: %v\n", f, resp.GetErrors())
			continue
		}

		fmt.Fprintf(os.Stdout, "Successfully ingested %d entities from %s\n", len(entities), f)
	}
}
