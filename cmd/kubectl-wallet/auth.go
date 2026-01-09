package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/util"
)

// createWallet crée un nouveau wallet Ethereum
func createWallet() {
	fmt.Println("Création d'un nouveau wallet...")

	// Générer une nouvelle paire de clés
	privateKey, address, err := util.GenerateWallet()
	if err != nil {
		log.Fatalf("Erreur lors de la génération du wallet: %v", err)
	}

	// Afficher les informations
	fmt.Println()
	fmt.Println("=== Nouveau Wallet Créé ===")
	fmt.Printf("Adresse:       %s\n", address.Hex())
	fmt.Printf("Clé privée:    %s\n", util.PrivateKeyToHex(privateKey))
	fmt.Println()
	fmt.Println("IMPORTANT: Sauvegardez votre clé privée en lieu sûr!")
	fmt.Println("   Cette clé ne peut pas être récupérée si elle est perdue.")
	fmt.Println()
	fmt.Println("Pour utiliser ce wallet:")
	fmt.Printf("  export WALLET_PRIVATE_KEY=%s\n", util.PrivateKeyToHex(privateKey))
	fmt.Println()
}

// getPrivateKeyFromArgs récupère la clé privée depuis les arguments
func getPrivateKeyFromArgs() string {
	if len(os.Args) < 3 {
		fmt.Println("Erreur: Clé privée requise")
		fmt.Println("Usage: kubectl wallet generate <private-key>")
		fmt.Println()
		fmt.Println("Vous pouvez aussi définir la variable d'environnement:")
		fmt.Println("  export WALLET_PRIVATE_KEY=0xYourPrivateKey")
		fmt.Println("  kubectl wallet generate")
		os.Exit(1)
	}
	return os.Args[2]
}

// getPrivateKeyFromEnvOrArgs récupère la clé privée depuis l'environnement ou les arguments
func getPrivateKeyFromEnvOrArgs() string {
	// Essayer d'abord l'environnement
	if privateKey := os.Getenv("WALLET_PRIVATE_KEY"); privateKey != "" {
		return privateKey
	}

	// Sinon, depuis les arguments
	if len(os.Args) >= 3 {
		return os.Args[2]
	}

	fmt.Println("Erreur: Clé privée requise")
	fmt.Println()
	fmt.Println("Vous pouvez la fournir de deux façons:")
	fmt.Println("  1. Comme argument: kubectl wallet generate <private-key>")
	fmt.Println("  2. Comme variable d'environnement: export WALLET_PRIVATE_KEY=0xYourPrivateKey")
	os.Exit(1)
	return ""
}
