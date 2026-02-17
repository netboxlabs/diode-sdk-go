// Package main demonstrates ingesting IKEProposal entities using the Diode SDK.
// This example includes three patterns: Minimal, Extended, and Explicit.
package main

import (
	"context"
	"log"

	"github.com/netboxlabs/diode-sdk-go/diode"
)

const (
	target     = "grpc://localhost:8080/diode"
	appName    = "ike_proposal-example"
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
	ikeProposal := IKEProposalMinimal()
	// ikeProposal := IKEProposalExtended()
	// ikeProposal := IKEProposalExplicit()

	resp, err := client.Ingest(context.Background(), []diode.Entity{ikeProposal})
	if err != nil {
		log.Fatalf("Ingestion failed: %v", err)
	}
	if resp.Errors != nil {
		log.Printf("Errors: %v", resp.Errors)
	} else {
		log.Println("IKEProposal ingested successfully")
	}
}

// IKEProposalMinimal Creates a IKEProposal with only required fields.
func IKEProposalMinimal() *diode.IKEProposal {
	return &diode.IKEProposal{
		Name:                 diode.String("Example Name"),
		AuthenticationMethod: diode.String("certificates"),
		EncryptionAlgorithm:  diode.String("3des-cbc"),
		Group:                diode.Int64(1),
		Metadata:             diode.Metadata{"source": "example"},
	}
}

// IKEProposalExtended Creates a IKEProposal with common optional fields.
func IKEProposalExtended() *diode.IKEProposal {
	return &diode.IKEProposal{
		Name:                    diode.String("Example Name"),
		AuthenticationMethod:    diode.String("certificates"),
		EncryptionAlgorithm:     diode.String("3des-cbc"),
		Group:                   diode.Int64(1),
		Metadata:                diode.Metadata{"source": "example", "custom_key": "custom_value"},
		Description:             diode.String("Example description"),
		AuthenticationAlgorithm: diode.String("hmac-md5"),
		SaLifetime:              diode.Int64(1),
		Comments:                diode.String("Example comments"),
	}
}

// IKEProposalExplicit Creates a IKEProposal with fully nested objects and all common fields.
func IKEProposalExplicit() *diode.IKEProposal {
	return &diode.IKEProposal{
		Name:                    diode.String("Example Name"),
		AuthenticationMethod:    diode.String("certificates"),
		EncryptionAlgorithm:     diode.String("3des-cbc"),
		Group:                   diode.Int64(1),
		Metadata:                diode.Metadata{"source": "example", "custom_key": "custom_value", "collected_at": "2024-01-15T10:30:00Z"},
		Description:             diode.String("Example description"),
		Comments:                diode.String("Example comments"),
		AuthenticationAlgorithm: diode.String("hmac-md5"),
		SaLifetime:              diode.Int64(1),
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
