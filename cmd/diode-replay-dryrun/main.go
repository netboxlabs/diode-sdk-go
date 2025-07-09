package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/netboxlabs/diode-sdk-go/diode"
	"github.com/netboxlabs/diode-sdk-go/diode/v1/diodepb"
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

	var allEntities []*diodepb.Entity
	for _, f := range files {
		ents, err := diode.LoadDryRunEntities(f)
		if err != nil {
			log.Fatal(err)
		}
		allEntities = append(allEntities, ents...)
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
			log.Printf("failed to close client: %v", err)
		}
	}()

	response, err := client.IngestProto(context.Background(), allEntities)
	if err != nil {
		log.Fatal(err)
	}

	if response.GetErrors() != nil {
		log.Fatalf("Ingestion errors: %v", response.GetErrors())
	}

	log.Printf("Ingested %d entities successfully", len(allEntities))

}
