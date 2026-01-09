package blockchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Permission représente les permissions d'un wallet dans le smart contract
type Permission struct {
	Namespaces []string
	Verbs      []string
	Resources  []string
	ExpiresAt  uint64
}

// WalletInfo contient les informations d'un wallet Ethereum
type WalletInfo struct {
	Address    common.Address
	PrivateKey string
}

// AccessRequest représente une demande d'accès à vérifier
type AccessRequest struct {
	WalletAddress string
	Namespace     string
	Verb          string
	Resource      string
}

// SignatureVerification représente le résultat d'une vérification de signature
type SignatureVerification struct {
	Valid             bool
	RecoveredAddress  common.Address
	ExpectedAddress   common.Address
	Message           string
	Signature         []byte
}

// ContractConfig représente la configuration du smart contract
type ContractConfig struct {
	Address common.Address
	ChainID *big.Int
	RPCURL  string
}

// TransactionOpts contient les options pour une transaction blockchain
type TransactionOpts struct {
	GasLimit uint64
	GasPrice *big.Int
	Nonce    *big.Int
	Value    *big.Int
}
