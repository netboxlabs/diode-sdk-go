// Package main demonstrates ingesting IPSecProposal entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ip_sec_proposal-example"
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
	ipSecProposal := IPSecProposalMinimal()
	// ipSecProposal := IPSecProposalExtended()
	// ipSecProposal := IPSecProposalExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ipSecProposal})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IPSecProposal ingested successfully")
	}
}

// IPSecProposalMinimal Creates a IPSecProposal with only required fields.
func IPSecProposalMinimal() *diode.IPSecProposal {
	return &diode.IPSecProposal{
		Name:     diode.String("Example Name"),
		Metadata: diode.Metadata{"source": "example"},
	}
}

// IPSecProposalExtended Creates a IPSecProposal with common optional fields.
func IPSecProposalExtended() *diode.IPSecProposal {
	return &diode.IPSecProposal{
		Name:                    diode.String("Example Name"),
		Metadata:                diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:             diode.String("Example description"),
		EncryptionAlgorithm:     diode.String("3des-cbc"),
		AuthenticationAlgorithm: diode.String("hmac-md5"),
		SaLifetimeSeconds:       diode.Int64(1),
		SaLifetimeData:          diode.Int64(1),
		Comments:                diode.String("Example comments"),
	}
}

// IPSecProposalExplicit Creates a IPSecProposal with fully nested objects and all common fields.
func IPSecProposalExplicit() *diode.IPSecProposal {
	return &diode.IPSecProposal{
		Name:                    diode.String("Example Name"),
		Metadata:                diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:             diode.String("Example description"),
		Comments:                diode.String("Example comments"),
		EncryptionAlgorithm:     diode.String("3des-cbc"),
		AuthenticationAlgorithm: diode.String("hmac-md5"),
		SaLifetimeSeconds:       diode.Int64(1),
		SaLifetimeData:          diode.Int64(1),
		Owner: &diode.Owner{
			Name: diode.String("Example Name"),
			Group: &diode.OwnerGroup{
				Name:     diode.String("Example Name"),
				Metadata: diode.Metadata{"source": "example"},
			},
			Metadata: diode.Metadata{"source": "example"},
		},
		Tags: []*diode.Tag{{Name: diode.String("production")}},
	}
}
