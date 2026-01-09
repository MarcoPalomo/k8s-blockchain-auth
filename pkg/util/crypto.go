package util

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateWallet génère une nouvelle paire de clés Ethereum
func GenerateWallet() (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to generate key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, fmt.Errorf("failed to convert public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, address, nil
}

// PrivateKeyToAddress convertit une clé privée en adresse Ethereum
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

// PrivateKeyFromHex charge une clé privée depuis une chaîne hexadécimale
func PrivateKeyFromHex(hexKey string) (*ecdsa.PrivateKey, error) {
	hexKey = StripHexPrefix(hexKey)
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}
	return privateKey, nil
}

// PrivateKeyToHex convertit une clé privée en chaîne hexadécimale
func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	return fmt.Sprintf("0x%x", crypto.FromECDSA(privateKey))
}

// SignMessage signe un message avec une clé privée
func SignMessage(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	messageHash := crypto.Keccak256Hash([]byte(message))
	signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %w", err)
	}
	return fmt.Sprintf("0x%x", signature), nil
}

// VerifySignature vérifie qu'un message a été signé par une adresse spécifique
func VerifySignature(address, message, signature string) (bool, error) {
	// Calculer le hash du message
	messageHash := crypto.Keccak256Hash([]byte(message))

	// Convertir la signature hex en bytes
	sig := common.FromHex(signature)
	if len(sig) != 65 {
		return false, fmt.Errorf("invalid signature length: %d", len(sig))
	}

	// Normaliser le v (recovery ID)
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	// Récupérer la clé publique depuis la signature
	pubKey, err := crypto.SigToPub(messageHash.Bytes(), sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Vérifier que l'adresse correspond
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	expectedAddr := common.HexToAddress(address)

	return recoveredAddr == expectedAddr, nil
}

// RecoverAddress récupère l'adresse qui a signé un message
func RecoverAddress(message, signature string) (common.Address, error) {
	messageHash := crypto.Keccak256Hash([]byte(message))
	sig := common.FromHex(signature)

	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("invalid signature length: %d", len(sig))
	}

	if sig[64] >= 27 {
		sig[64] -= 27
	}

	pubKey, err := crypto.SigToPub(messageHash.Bytes(), sig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to recover public key: %w", err)
	}

	return crypto.PubkeyToAddress(*pubKey), nil
}

// HashMessage calcule le hash Keccak256 d'un message
func HashMessage(message string) string {
	hash := crypto.Keccak256Hash([]byte(message))
	return hash.Hex()
}

// StripHexPrefix retire le préfixe "0x" d'une chaîne hexadécimale
func StripHexPrefix(s string) string {
	if len(s) >= 2 && strings.ToLower(s[0:2]) == "0x" {
		return s[2:]
	}
	return s
}

// AddHexPrefix ajoute le préfixe "0x" à une chaîne si elle ne l'a pas
func AddHexPrefix(s string) string {
	if len(s) >= 2 && strings.ToLower(s[0:2]) == "0x" {
		return s
	}
	return "0x" + s
}

// IsValidAddress vérifie si une chaîne est une adresse Ethereum valide
func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

// NormalizeAddress normalise une adresse Ethereum (checksummed)
func NormalizeAddress(address string) string {
	return common.HexToAddress(address).Hex()
}
