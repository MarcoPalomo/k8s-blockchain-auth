package authenticator

import (
	"context"
	"fmt"

	"k8s.io/apiserver/pkg/authorization/authorizer"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain"
)

type BlockchainAuthorizer struct {
	blockchain *blockchain.Client
}

func NewBlockchainAuthorizer(bc *blockchain.Client) *BlockchainAuthorizer {
	return &BlockchainAuthorizer{blockchain: bc}
}

func (b *BlockchainAuthorizer) Authorize(
	ctx context.Context,
	attrs authorizer.Attributes,
) (authorized authorizer.Decision, reason string, err error) {

	// Extraire le wallet address du user
	walletAddress := attrs.GetUser().GetName()

	// Vérifier dans le smart contract
	hasPermission, err := b.blockchain.HasPermission(
		walletAddress,
		attrs.GetNamespace(),
		attrs.GetVerb(),
		attrs.GetResource(),
	)

	if err != nil {
		return authorizer.DecisionNoOpinion,
			fmt.Sprintf("blockchain query failed: %v", err),
			err
	}

	if hasPermission {
		return authorizer.DecisionAllow, "allowed by blockchain", nil
	}

	return authorizer.DecisionDeny,
		"no blockchain permission",
		nil
}
