package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/marcopalomo/k8s-blockchain-auth/pkg/util"
)

func revokePermission() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur: Adresse wallet requise")
		fmt.Println("Usage: permission-manager revoke <wallet>")
		os.Exit(1)
	}

	walletAddress := os.Args[2]

	if err := util.ValidateWalletAddress(walletAddress); err != nil {
		log.Fatalf("Adresse wallet invalide: %v", err)
	}

	client, opts := connectBlockchain()

	fmt.Printf("Révocation des permissions de %s...\n", walletAddress)

	tx, err := client.RevokePermission(opts, common.HexToAddress(walletAddress))
	if err != nil {
		log.Fatalf("Erreur lors de la révocation: %v", err)
	}

	fmt.Printf("Transaction envoyée: %s\n", tx.Hash().Hex())
	fmt.Println("En attente de confirmation...")

	receipt, err := bind.WaitMined(context.Background(), client.GetEthClient(), tx)
	if err != nil {
		log.Fatalf("Erreur lors de l'attente: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("Permissions révoquées avec succès!")
	} else {
		fmt.Println("Transaction échouée!")
		os.Exit(1)
	}
}
