package authenticator

import (
	"fmt"
	"net/http"
	"time"

	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain"
)

type BlockchainAuthenticator struct {
	blockchain     *blockchain.Client
	cache          *Cache
	tokenExtractor *TokenExtractor
}

func NewBlockchainAuthenticator(bc *blockchain.Client) *BlockchainAuthenticator {
	return &BlockchainAuthenticator{
		blockchain:     bc,
		cache:          NewCache(5 * time.Minute),
		tokenExtractor: NewTokenExtractor(5 * time.Minute),
	}
}

func (b *BlockchainAuthenticator) AuthenticateRequest(
	req *http.Request,
) (*authenticator.Response, bool, error) {

	// Extraire et valider le token
	token, err := b.tokenExtractor.ExtractToken(req)
	if err != nil {
		return nil, false, err
	}

	// Vérifier la signature
	valid, err := b.blockchain.VerifySignature(
		token.WalletAddress,
		token.Message,
		token.Signature,
	)
	if err != nil || !valid {
		return nil, false, fmt.Errorf("invalid signature")
	}

	// Vérifier le cache
	if userInfo, ok := b.cache.Get(token.WalletAddress); ok {
		return &authenticator.Response{User: userInfo}, true, nil
	}

	// Récupérer les permissions depuis la blockchain
	perms, err := b.blockchain.GetPermissions(token.WalletAddress)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get permissions: %v", err)
	}

	// Construire les groups K8s depuis les permissions blockchain
	groups := b.buildGroups(perms)

	userInfo := &user.DefaultInfo{
		Name:   token.WalletAddress,
		UID:    token.WalletAddress,
		Groups: groups,
		Extra: map[string][]string{
			"blockchain.io/namespaces": perms.Namespaces,
			"blockchain.io/verbs":      perms.Verbs,
			"blockchain.io/resources":  perms.Resources,
		},
	}

	// Mettre en cache
	b.cache.Set(token.WalletAddress, userInfo)

	return &authenticator.Response{User: userInfo}, true, nil
}

func (b *BlockchainAuthenticator) buildGroups(perms *blockchain.Permission) []string {
	groups := []string{"system:authenticated"}

	// Créer des groups basés sur les capabilities
	if b.hasAdminPerms(perms) {
		groups = append(groups, "blockchain:admin")
	}

	if b.hasDeveloperPerms(perms) {
		groups = append(groups, "blockchain:developer")
	}

	// Ajouter un group par namespace
	for _, ns := range perms.Namespaces {
		if ns != "*" {
			groups = append(groups, fmt.Sprintf("blockchain:namespace:%s", ns))
		}
	}

	return groups
}

func (b *BlockchainAuthenticator) hasAdminPerms(perms *blockchain.Permission) bool {
	for _, verb := range perms.Verbs {
		if verb == "*" {
			return true
		}
	}
	return false
}

func (b *BlockchainAuthenticator) hasDeveloperPerms(perms *blockchain.Permission) bool {
	requiredVerbs := map[string]bool{"get": false, "list": false, "create": false}
	for _, verb := range perms.Verbs {
		if _, ok := requiredVerbs[verb]; ok {
			requiredVerbs[verb] = true
		}
	}

	for _, has := range requiredVerbs {
		if !has {
			return false
		}
	}
	return true
}
