package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "generate":
		generateToken()
	case "create-wallet":
		createWallet()
	case "verify":
		verifyToken()
	case "info":
		showTokenInfo()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("kubectl-wallet - Utilitaire pour l'authentification blockchain K8s")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  kubectl wallet create-wallet             Créer un nouveau wallet")
	fmt.Println("  kubectl wallet generate [private-key]    Générer un token d'authentification")
	fmt.Println("  kubectl wallet verify <token>            Vérifier un token")
	fmt.Println("  kubectl wallet info <token>              Afficher les infos d'un token")
	fmt.Println("  kubectl wallet help                      Afficher cette aide")
	fmt.Println()
	fmt.Println("Exemples:")
	fmt.Println("  # Créer un nouveau wallet")
	fmt.Println("  kubectl wallet create-wallet")
	fmt.Println()
	fmt.Println("  # Générer un token d'authentification (clé privée en argument)")
	fmt.Println("  kubectl wallet generate 0xYourPrivateKey")
	fmt.Println()
	fmt.Println("  # Générer un token (clé privée depuis l'environnement)")
	fmt.Println("  export WALLET_PRIVATE_KEY=0xYourPrivateKey")
	fmt.Println("  kubectl wallet generate")
	fmt.Println()
	fmt.Println("  # Utiliser le token avec kubectl")
	fmt.Println("  export K8S_AUTH_TOKEN=$(kubectl wallet generate)")
	fmt.Println("  kubectl --token=\"$K8S_AUTH_TOKEN\" get pods")
	fmt.Println()
	fmt.Println("  # Vérifier un token")
	fmt.Println("  kubectl wallet verify '{\"wallet\":\"0x...\",\"message\":\"...\",\"signature\":\"0x...\",\"timestamp\":1234567890}'")
	fmt.Println()
}
