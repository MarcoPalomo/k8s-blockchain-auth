package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain"
	"github.com/marcopalomo/k8s-blockchain-auth/pkg/util"
)

func grantPermission() {
	if len(os.Args) < 6 {
		fmt.Println("Erreur: Arguments manquants")
		fmt.Println("Usage: permission-manager grant <wallet> <namespaces> <verbs> <resources> [expiration]")
		fmt.Println()
		fmt.Println("Arguments:")
		fmt.Println("  wallet:      Adresse Ethereum du wallet (ex: 0x123...)")
		fmt.Println("  namespaces:  Liste de namespaces séparés par des virgules (ex: default,prod ou '*')")
		fmt.Println("  verbs:       Liste de verbes séparés par des virgules (ex: get,list,create ou '*')")
		fmt.Println("  resources:   Liste de ressources séparées par des virgules (ex: pods,services ou '*')")
		fmt.Println("  expiration:  (optionnel) Timestamp d'expiration Unix ou 0 pour sans expiration")
		os.Exit(1)
	}

	walletAddress := os.Args[2]
	namespacesStr := os.Args[3]
	verbsStr := os.Args[4]
	resourcesStr := os.Args[5]

	// Expiration optionnelle
	var expiresAt uint64 = 0
	if len(os.Args) > 6 {
		fmt.Sscanf(os.Args[6], "%d", &expiresAt)
	}

	// Parser les listes
	namespaces := strings.Split(namespacesStr, ",")
	verbs := strings.Split(verbsStr, ",")
	resources := strings.Split(resourcesStr, ",")

	// Valider les entrées
	if err := util.ValidateWalletAddress(walletAddress); err != nil {
		log.Fatalf("Adresse wallet invalide: %v", err)
	}

	if err := util.ValidatePermissions(namespaces, verbs, resources); err != nil {
		log.Fatalf("Permissions invalides: %v", err)
	}

	if expiresAt > 0 {
		if err := util.ValidateExpiration(expiresAt); err != nil {
			log.Fatalf("Expiration invalide: %v", err)
		}
	}

	// Connexion blockchain
	client, opts := connectBlockchain()

	// Accorder la permission
	fmt.Printf("Accord de permissions à %s...\n", walletAddress)
	fmt.Printf("  Namespaces: %v\n", namespaces)
	fmt.Printf("  Verbes:     %v\n", verbs)
	fmt.Printf("  Ressources: %v\n", resources)
	if expiresAt == 0 {
		fmt.Printf("  Expiration: Jamais\n")
	} else {
		fmt.Printf("  Expiration: %s\n", time.Unix(int64(expiresAt), 0))
	}

	tx, err := client.GrantPermission(
		opts,
		common.HexToAddress(walletAddress),
		namespaces,
		verbs,
		resources,
		expiresAt,
	)
	if err != nil {
		log.Fatalf("Erreur lors de l'accord de permission: %v", err)
	}

	fmt.Printf("\nTransaction envoyée: %s\n", tx.Hash().Hex())
	fmt.Println("En attente de confirmation...")

	// Attendre la confirmation
	receipt, err := bind.WaitMined(context.Background(), client.GetEthClient(), tx)
	if err != nil {
		log.Fatalf("Erreur lors de l'attente de confirmation: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("Permissions accordées avec succès!")
	} else {
		fmt.Println("Transaction échouée!")
		os.Exit(1)
	}
}

func connectBlockchain() (*blockchain.Client, *bind.TransactOpts) {
	rpcURL := os.Getenv("BLOCKCHAIN_RPC_URL")
	if rpcURL == "" {
		log.Fatal("BLOCKCHAIN_RPC_URL environment variable is required")
	}

	contractAddr := os.Getenv("CONTRACT_ADDRESS")
	if contractAddr == "" {
		log.Fatal("CONTRACT_ADDRESS environment variable is required")
	}

	privateKey := os.Getenv("ADMIN_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("ADMIN_PRIVATE_KEY environment variable is required")
	}

	// Connexion au client
	client, err := blockchain.NewClient(rpcURL, contractAddr)
	if err != nil {
		log.Fatalf("Impossible de se connecter à la blockchain: %v", err)
	}

	// Charger la clé privée admin
	key, err := util.PrivateKeyFromHex(privateKey)
	if err != nil {
		log.Fatalf("Clé privée admin invalide: %v", err)
	}

	// Créer les options de transaction
	chainID := client.GetChainID()
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("Impossible de créer le transactor: %v", err)
	}

	opts.GasLimit = 300000
	opts.GasPrice = big.NewInt(20000000000) // 20 Gwei

	return client, opts
}
