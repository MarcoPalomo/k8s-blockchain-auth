package util

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ValidateWalletAddress valide une adresse de wallet Ethereum
func ValidateWalletAddress(address string) error {
	if address == "" {
		return fmt.Errorf("wallet address cannot be empty")
	}

	if !IsValidAddress(address) {
		return fmt.Errorf("invalid Ethereum address format: %s", address)
	}

	return nil
}

// ValidateNamespace valide un nom de namespace Kubernetes
func ValidateNamespace(namespace string) error {
	if namespace == "" {
		return fmt.Errorf("namespace cannot be empty")
	}

	// Un namespace peut être "*" pour tous les namespaces
	if namespace == "*" {
		return nil
	}

	// Règles K8s: lowercase alphanumeric, hyphens, max 253 chars
	if len(namespace) > 253 {
		return fmt.Errorf("namespace too long (max 253 characters)")
	}

	validNamespace := regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	if !validNamespace.MatchString(namespace) {
		return fmt.Errorf("invalid namespace format: %s (must be lowercase alphanumeric with hyphens)", namespace)
	}

	return nil
}

// ValidateVerb valide un verbe Kubernetes
func ValidateVerb(verb string) error {
	if verb == "" {
		return fmt.Errorf("verb cannot be empty")
	}

	// Un verbe peut être "*" pour tous les verbes
	if verb == "*" {
		return nil
	}

	// Verbes K8s standards
	validVerbs := map[string]bool{
		"get":             true,
		"list":            true,
		"watch":           true,
		"create":          true,
		"update":          true,
		"patch":           true,
		"delete":          true,
		"deletecollection": true,
		"*":               true,
	}

	if !validVerbs[verb] {
		return fmt.Errorf("invalid verb: %s", verb)
	}

	return nil
}

// ValidateResource valide un nom de ressource Kubernetes
func ValidateResource(resource string) error {
	if resource == "" {
		return fmt.Errorf("resource cannot be empty")
	}

	// Une ressource peut être "*" pour toutes les ressources
	if resource == "*" {
		return nil
	}

	// Règles K8s: lowercase alphanumeric, hyphens, max 63 chars
	if len(resource) > 63 {
		return fmt.Errorf("resource too long (max 63 characters)")
	}

	validResource := regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	if !validResource.MatchString(resource) {
		return fmt.Errorf("invalid resource format: %s", resource)
	}

	return nil
}

// ValidatePermissions valide un ensemble de permissions
func ValidatePermissions(namespaces, verbs, resources []string) error {
	if len(namespaces) == 0 {
		return fmt.Errorf("at least one namespace is required")
	}

	if len(verbs) == 0 {
		return fmt.Errorf("at least one verb is required")
	}

	if len(resources) == 0 {
		return fmt.Errorf("at least one resource is required")
	}

	// Valider chaque namespace
	for _, ns := range namespaces {
		if err := ValidateNamespace(ns); err != nil {
			return fmt.Errorf("invalid namespace: %w", err)
		}
	}

	// Valider chaque verbe
	for _, verb := range verbs {
		if err := ValidateVerb(verb); err != nil {
			return fmt.Errorf("invalid verb: %w", err)
		}
	}

	// Valider chaque ressource
	for _, resource := range resources {
		if err := ValidateResource(resource); err != nil {
			return fmt.Errorf("invalid resource: %w", err)
		}
	}

	return nil
}

// ValidateExpiration valide un timestamp d'expiration
func ValidateExpiration(expiresAt uint64) error {
	if expiresAt == 0 {
		// 0 signifie pas d'expiration
		return nil
	}

	now := uint64(time.Now().Unix())
	if expiresAt <= now {
		return fmt.Errorf("expiration time must be in the future")
	}

	// Vérifier que l'expiration n'est pas trop loin dans le futur (ex: max 10 ans)
	maxExpiration := now + (10 * 365 * 24 * 60 * 60) // 10 ans
	if expiresAt > maxExpiration {
		return fmt.Errorf("expiration time too far in the future (max 10 years)")
	}

	return nil
}

// ValidateSignature valide le format d'une signature
func ValidateSignature(signature string) error {
	if signature == "" {
		return fmt.Errorf("signature cannot be empty")
	}

	// Retirer le préfixe 0x si présent
	sig := StripHexPrefix(signature)

	// Une signature ECDSA Ethereum fait 65 bytes = 130 caractères hex
	if len(sig) != 130 {
		return fmt.Errorf("invalid signature length: expected 130 hex characters, got %d", len(sig))
	}

	// Vérifier que c'est bien de l'hexadécimal
	validHex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	if !validHex.MatchString(sig) {
		return fmt.Errorf("signature must be hexadecimal")
	}

	return nil
}

// ValidatePrivateKey valide le format d'une clé privée
func ValidatePrivateKey(privateKey string) error {
	if privateKey == "" {
		return fmt.Errorf("private key cannot be empty")
	}

	// Retirer le préfixe 0x si présent
	key := StripHexPrefix(privateKey)

	// Une clé privée Ethereum fait 32 bytes = 64 caractères hex
	if len(key) != 64 {
		return fmt.Errorf("invalid private key length: expected 64 hex characters, got %d", len(key))
	}

	// Vérifier que c'est bien de l'hexadécimal
	validHex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	if !validHex.MatchString(key) {
		return fmt.Errorf("private key must be hexadecimal")
	}

	return nil
}

// SanitizeInput nettoie une chaîne d'entrée utilisateur
func SanitizeInput(input string) string {
	// Retirer les espaces en début et fin
	input = strings.TrimSpace(input)

	// Retirer les caractères de contrôle
	input = strings.Map(func(r rune) rune {
		if r < 32 && r != '\n' && r != '\t' {
			return -1
		}
		return r
	}, input)

	return input
}

// ValidateRPCURL valide une URL RPC blockchain
func ValidateRPCURL(url string) error {
	if url == "" {
		return fmt.Errorf("RPC URL cannot be empty")
	}

	// Vérifier que l'URL commence par http:// ou https:// ou ws:// ou wss://
	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") &&
		!strings.HasPrefix(url, "ws://") &&
		!strings.HasPrefix(url, "wss://") {
		return fmt.Errorf("RPC URL must start with http://, https://, ws://, or wss://")
	}

	return nil
}
