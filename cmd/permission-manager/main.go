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
	case "grant":
		grantPermission()
	case "revoke":
		revokePermission()
	case "list":
		listPermissions()
	case "check":
		checkPermission()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("permission-manager - Gestionnaire de permissions blockchain K8s")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  permission-manager grant <wallet> <namespaces> <verbs> <resources> [expiration]")
	fmt.Println("  permission-manager revoke <wallet>")
	fmt.Println("  permission-manager list <wallet>")
	fmt.Println("  permission-manager check <wallet> <namespace> <verb> <resource>")
	fmt.Println("  permission-manager help")
	fmt.Println()
	fmt.Println("Exemples:")
	fmt.Println("  # Accorder des permissions")
	fmt.Println("  permission-manager grant 0x123... default,prod get,list,create pods,services")
	fmt.Println()
	fmt.Println("  # Accorder toutes les permissions (admin)")
	fmt.Println("  permission-manager grant 0x123... '*' '*' '*'")
	fmt.Println()
	fmt.Println("  # Révoquer des permissions")
	fmt.Println("  permission-manager revoke 0x123...")
	fmt.Println()
	fmt.Println("  # Lister les permissions")
	fmt.Println("  permission-manager list 0x123...")
	fmt.Println()
	fmt.Println("  # Vérifier une permission spécifique")
	fmt.Println("  permission-manager check 0x123... default get pods")
	fmt.Println()
}
