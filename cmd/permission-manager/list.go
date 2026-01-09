package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/util"
)

func listPermissions() {
	if len(os.Args) < 3 {
		fmt.Println("Erreur: Adresse wallet requise")
		fmt.Println("Usage: permission-manager list <wallet>")
		os.Exit(1)
	}

	walletAddress := os.Args[2]

	if err := util.ValidateWalletAddress(walletAddress); err != nil {
		log.Fatalf("Adresse wallet invalide: %v", err)
	}

	client, _ := connectBlockchain()

	fmt.Printf("Permissions pour le wallet %s:\n\n", walletAddress)

	perms, err := client.GetPermissions(walletAddress)
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des permissions: %v", err)
	}

	if len(perms.Namespaces) == 0 {
		fmt.Println("Aucune permission trouvée.")
		return
	}

	fmt.Printf("Namespaces:  %v\n", perms.Namespaces)
	fmt.Printf("Verbes:      %v\n", perms.Verbs)
	fmt.Printf("Ressources:  %v\n", perms.Resources)

	if perms.ExpiresAt == 0 {
		fmt.Printf("Expiration:  Jamais\n")
	} else {
		expTime := time.Unix(int64(perms.ExpiresAt), 0)
		fmt.Printf("Expiration:  %s\n", expTime)

		if time.Now().After(expTime) {
			fmt.Printf("Statut:      EXPIRÉ\n")
		} else {
			fmt.Printf("Statut:      ACTIF\n")
		}
	}
}

func checkPermission() {
	if len(os.Args) < 6 {
		fmt.Println("Erreur: Arguments manquants")
		fmt.Println("Usage: permission-manager check <wallet> <namespace> <verb> <resource>")
		os.Exit(1)
	}

	walletAddress := os.Args[2]
	namespace := os.Args[3]
	verb := os.Args[4]
	resource := os.Args[5]

	if err := util.ValidateWalletAddress(walletAddress); err != nil {
		log.Fatalf("Adresse wallet invalide: %v", err)
	}

	client, _ := connectBlockchain()

	hasPermission, err := client.HasPermission(walletAddress, namespace, verb, resource)
	if err != nil {
		log.Fatalf("Erreur lors de la vérification: %v", err)
	}

	fmt.Printf("Vérification de permission:\n")
	fmt.Printf("  Wallet:     %s\n", walletAddress)
	fmt.Printf("  Namespace:  %s\n", namespace)
	fmt.Printf("  Verbe:      %s\n", verb)
	fmt.Printf("  Ressource:  %s\n", resource)
	fmt.Println()

	if hasPermission {
		fmt.Println("Résultat: AUTORISÉ")
	} else {
		fmt.Println("Résultat: REFUSÉ")
		os.Exit(1)
	}
}
