package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/authenticator"
	"github.com/marcopalomo/k8s-blockchain-auth/pkg/util"
)

// generateToken génère un token d'authentification signé
func generateToken() {
	// Récupérer la clé privée
	privateKeyHex := getPrivateKeyFromEnvOrArgs()

	// Valider et charger la clé privée
	if err := util.ValidatePrivateKey(privateKeyHex); err != nil {
		log.Fatalf("Clé privée invalide: %v", err)
	}

	privateKey, err := util.PrivateKeyFromHex(privateKeyHex)
	if err != nil {
		log.Fatalf("Impossible de charger la clé privée: %v", err)
	}

	// Obtenir l'adresse du wallet
	walletAddress := util.PrivateKeyToAddress(privateKey)

	// Créer le message à signer
	timestamp := time.Now().Unix()
	message := fmt.Sprintf("K8s Auth Request - %s - %d", walletAddress.Hex(), timestamp)

	// Signer le message
	signature, err := util.SignMessage(privateKey, message)
	if err != nil {
		log.Fatalf("Erreur lors de la signature: %v", err)
	}

	// Créer le token
	token := authenticator.CreateToken(
		walletAddress.Hex(),
		message,
		signature,
	)

	// Encoder en JSON
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		log.Fatalf("Erreur lors de l'encodage JSON: %v", err)
	}

	// Afficher le token (juste le JSON, pas de Bearer)
	fmt.Print(string(tokenJSON))
}

// verifyToken vérifie la validité d'un token
func verifyToken() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur: Token JSON requis")
		fmt.Println("Usage: kubectl wallet verify '<token-json>'")
		os.Exit(1)
	}

	tokenJSON := os.Args[2]

	// Décoder le token
	var token authenticator.AuthToken
	if err := json.Unmarshal([]byte(tokenJSON), &token); err != nil {
		log.Fatalf("Token invalide: %v", err)
	}

	// Vérifier la signature
	valid, err := util.VerifySignature(token.WalletAddress, token.Message, token.Signature)
	if err != nil {
		log.Fatalf("Erreur lors de la vérification: %v", err)
	}

	if !valid {
		fmt.Println("Signature invalide")
		os.Exit(1)
	}

	// Vérifier l'expiration
	now := time.Now().Unix()
	age := now - token.Timestamp

	fmt.Println("=== Vérification du Token ===")
	fmt.Printf("Wallet:        %s\n", token.WalletAddress)
	fmt.Printf("Message:       %s\n", token.Message)
	fmt.Printf("Timestamp:     %d (%s)\n", token.Timestamp, time.Unix(token.Timestamp, 0))
	fmt.Printf("Age:           %d secondes\n", age)
	fmt.Printf("Signature:     Valide\n")

	if age > 300 {
		fmt.Printf("Statut:        Expiré (max 300 secondes)\n")
	} else {
		fmt.Printf("Statut:        Valide\n")
	}
}

// showTokenInfo affiche les informations d'un token sans vérifier la signature
func showTokenInfo() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur: Token JSON requis")
		fmt.Println("Usage: kubectl wallet info '<token-json>'")
		os.Exit(1)
	}

	tokenJSON := os.Args[2]

	// Décoder le token
	var token authenticator.AuthToken
	if err := json.Unmarshal([]byte(tokenJSON), &token); err != nil {
		log.Fatalf("Token invalide: %v", err)
	}

	// Afficher les informations
	fmt.Println("=== Informations du Token ===")
	fmt.Printf("Wallet:        %s\n", token.WalletAddress)
	fmt.Printf("Message:       %s\n", token.Message)
	fmt.Printf("Timestamp:     %d (%s)\n", token.Timestamp, time.Unix(token.Timestamp, 0))
	fmt.Printf("Signature:     %s\n", token.Signature[:20]+"...")

	// Calculer l'âge
	now := time.Now().Unix()
	age := now - token.Timestamp
	fmt.Printf("Age:           %d secondes\n", age)

	if age > 300 {
		fmt.Printf("Statut:        Expiré\n")
	} else {
		fmt.Printf("Statut:        Valide (reste %d secondes)\n", 300-age)
	}
}
