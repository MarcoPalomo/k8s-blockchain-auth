package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client          *ethclient.Client
	contract        *K8sAccessControl // Généré via abigen
	contractAddress common.Address
	chainID         *big.Int
}

// GetEthClient returns the underlying ethclient for transaction operations
func (c *Client) GetEthClient() *ethclient.Client {
	return c.client
}

// GrantPermission grants permissions to a wallet
func (c *Client) GrantPermission(
	opts *bind.TransactOpts,
	wallet common.Address,
	namespaces []string,
	verbs []string,
	resources []string,
	expiresAt uint64,
) (*types.Transaction, error) {
	return c.contract.GrantPermission(opts, wallet, namespaces, verbs, resources, big.NewInt(int64(expiresAt)))
}

// RevokePermission revokes permissions from a wallet
func (c *Client) RevokePermission(
	opts *bind.TransactOpts,
	wallet common.Address,
) (*types.Transaction, error) {
	return c.contract.RevokePermission(opts, wallet)
}

func NewClient(rpcURL, contractAddress string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contract, err := NewK8sAccessControl(
		common.HexToAddress(contractAddress),
		client,
	)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Client{
		client:          client,
		contract:        contract,
		contractAddress: common.HexToAddress(contractAddress),
		chainID:         chainID,
	}, nil
}

func (c *Client) GetPermissions(walletAddress string) (*Permission, error) {
	addr := common.HexToAddress(walletAddress)

	result, err := c.contract.GetPermissions(
		&bind.CallOpts{},
		addr,
	)
	if err != nil {
		return nil, err
	}

	return &Permission{
		Namespaces: result.Namespaces,
		Verbs:      result.Verbs,
		Resources:  result.Resources,
		ExpiresAt:  result.ExpiresAt.Uint64(),
	}, nil
}

func (c *Client) HasPermission(
	walletAddress, namespace, verb, resource string,
) (bool, error) {
	addr := common.HexToAddress(walletAddress)

	return c.contract.HasPermission(
		&bind.CallOpts{},
		addr,
		namespace,
		verb,
		resource,
	)
}

// VerifySignature vérifie qu'un message a été signé par le wallet
func (c *Client) VerifySignature(
	walletAddress, message, signature string,
) (bool, error) {
	// Récupérer la clé publique depuis la signature
	hash := crypto.Keccak256Hash([]byte(message))
	sig := common.FromHex(signature)

	// Ethereum signe avec un v de 27 ou 28, on doit le normaliser
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	expectedAddr := common.HexToAddress(walletAddress)

	return recoveredAddr == expectedAddr, nil
}

// Close closes the blockchain client connection
func (c *Client) Close() {
	c.client.Close()
}

// GetChainID returns the chain ID of the blockchain
func (c *Client) GetChainID() *big.Int {
	return c.chainID
}

// GetContractAddress returns the smart contract address
func (c *Client) GetContractAddress() common.Address {
	return c.contractAddress
}
