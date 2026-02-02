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
	}
}

// IKEProposalExtended Creates a IKEProposal with common optional fields.
func IKEProposalExtended() *diode.IKEProposal {
	return &diode.IKEProposal{
		Name:                 diode.String("Example Name"),
		AuthenticationMethod: diode.String("certificates"),
		EncryptionAlgorithm:  diode.String("3des-cbc"),
		Group:                diode.Int64(1),
		Description:          diode.String("Example description"),
	}
}

// IKEProposalExplicit Creates a IKEProposal with fully nested objects and all common fields.
func IKEProposalExplicit() *diode.IKEProposal {
	return &diode.IKEProposal{
		Name:                 diode.String("Example Name"),
		AuthenticationMethod: diode.String("certificates"),
		EncryptionAlgorithm:  diode.String("3des-cbc"),
		Group:                diode.Int64(1),
		Description:          diode.String("Example description"),
		Comments:             diode.String("Example comments"),
		Tags:                 []*diode.Tag{{Name: diode.String("production")}},
	}
}
