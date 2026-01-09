package authenticator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// AuthToken représente le token d'authentification blockchain
type AuthToken struct {
	WalletAddress string `json:"wallet"`
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	Timestamp     int64  `json:"timestamp"`
}

// TokenExtractor extrait et valide un token d'authentification depuis une requête HTTP
type TokenExtractor struct {
	maxTokenAge time.Duration
}

// NewTokenExtractor crée un nouveau extracteur de token
func NewTokenExtractor(maxTokenAge time.Duration) *TokenExtractor {
	return &TokenExtractor{
		maxTokenAge: maxTokenAge,
	}
}

// ExtractToken extrait et valide le token depuis le header Authorization
func (te *TokenExtractor) ExtractToken(req *http.Request) (*AuthToken, error) {
	// Extraire le header Authorization
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing authorization header")
	}

	// Format: "Bearer <json_token>"
	// Use SplitN to only split on the first space
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)
	if tokenString == "" {
		return nil, fmt.Errorf("empty token")
	}

	// Décoder le token JSON
	var token AuthToken
	if err := json.Unmarshal([]byte(tokenString), &token); err != nil {
		return nil, fmt.Errorf("invalid token format: %w", err)
	}

	// Valider le token
	if err := te.ValidateToken(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

// ValidateToken vérifie la validité d'un token
func (te *TokenExtractor) ValidateToken(token *AuthToken) error {
	// Vérifier que les champs requis sont présents
	if token.WalletAddress == "" {
		return fmt.Errorf("missing wallet address")
	}

	if token.Message == "" {
		return fmt.Errorf("missing message")
	}

	if token.Signature == "" {
		return fmt.Errorf("missing signature")
	}

	if token.Timestamp == 0 {
		return fmt.Errorf("missing timestamp")
	}

	// Vérifier que le timestamp est récent (anti-replay)
	now := time.Now().Unix()
	tokenAge := now - token.Timestamp

	if tokenAge > int64(te.maxTokenAge.Seconds()) {
		return fmt.Errorf("token expired (age: %d seconds, max: %d seconds)",
			tokenAge, int64(te.maxTokenAge.Seconds()))
	}

	// Vérifier que le timestamp n'est pas dans le futur
	if tokenAge < 0 {
		return fmt.Errorf("token timestamp is in the future")
	}

	return nil
}

// CreateToken crée un nouveau token d'authentification (pour testing/client)
func CreateToken(walletAddress, message, signature string) *AuthToken {
	return &AuthToken{
		WalletAddress: walletAddress,
		Message:       message,
		Signature:     signature,
		Timestamp:     time.Now().Unix(),
	}
}

// ToJSON convertit le token en JSON
func (t *AuthToken) ToJSON() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", fmt.Errorf("failed to marshal token: %w", err)
	}
	return string(data), nil
}

// ToBearerToken convertit le token en format Bearer pour Authorization header
func (t *AuthToken) ToBearerToken() (string, error) {
	jsonToken, err := t.ToJSON()
	if err != nil {
		return "", err
	}
	return "Bearer " + jsonToken, nil
}
