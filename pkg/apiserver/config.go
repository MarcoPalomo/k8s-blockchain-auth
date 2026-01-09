package apiserver

import (
	"fmt"
	"os"
	"time"
)

// Config représente la configuration de l'API server
type Config struct {
	// Blockchain configuration
	BlockchainRPCURL    string
	ContractAddress     string

	// Server configuration
	Port                string
	Host                string

	// TLS configuration
	TLSCertFile         string
	TLSKeyFile          string
	EnableTLS           bool

	// Authentication configuration
	CacheTTL            time.Duration
	TokenMaxAge         time.Duration

	// Logging configuration
	LogLevel            string
	LogFormat           string

	// Health check configuration
	HealthCheckPath     string
	ReadinessCheckPath  string
}

// NewConfig crée une nouvelle configuration avec les valeurs par défaut
func NewConfig() *Config {
	return &Config{
		Port:                "8080",
		Host:                "0.0.0.0",
		CacheTTL:            5 * time.Minute,
		TokenMaxAge:         5 * time.Minute,
		LogLevel:            "info",
		LogFormat:           "json",
		HealthCheckPath:     "/healthz",
		ReadinessCheckPath:  "/readyz",
		EnableTLS:           false,
	}
}

// LoadFromEnv charge la configuration depuis les variables d'environnement
func (c *Config) LoadFromEnv() error {
	// Required configuration
	c.BlockchainRPCURL = os.Getenv("BLOCKCHAIN_RPC_URL")
	if c.BlockchainRPCURL == "" {
		return fmt.Errorf("BLOCKCHAIN_RPC_URL is required")
	}

	c.ContractAddress = os.Getenv("CONTRACT_ADDRESS")
	if c.ContractAddress == "" {
		return fmt.Errorf("CONTRACT_ADDRESS is required")
	}

	// Optional configuration
	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}

	if host := os.Getenv("HOST"); host != "" {
		c.Host = host
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		c.LogLevel = logLevel
	}

	if logFormat := os.Getenv("LOG_FORMAT"); logFormat != "" {
		c.LogFormat = logFormat
	}

	// TLS configuration
	c.TLSCertFile = os.Getenv("TLS_CERT_FILE")
	c.TLSKeyFile = os.Getenv("TLS_KEY_FILE")

	if c.TLSCertFile != "" && c.TLSKeyFile != "" {
		c.EnableTLS = true
	}

	// Parse duration configurations
	if cacheTTL := os.Getenv("CACHE_TTL"); cacheTTL != "" {
		duration, err := time.ParseDuration(cacheTTL)
		if err != nil {
			return fmt.Errorf("invalid CACHE_TTL: %w", err)
		}
		c.CacheTTL = duration
	}

	if tokenMaxAge := os.Getenv("TOKEN_MAX_AGE"); tokenMaxAge != "" {
		duration, err := time.ParseDuration(tokenMaxAge)
		if err != nil {
			return fmt.Errorf("invalid TOKEN_MAX_AGE: %w", err)
		}
		c.TokenMaxAge = duration
	}

	return nil
}

// Validate vérifie que la configuration est valide
func (c *Config) Validate() error {
	if c.BlockchainRPCURL == "" {
		return fmt.Errorf("blockchain RPC URL is required")
	}

	if c.ContractAddress == "" {
		return fmt.Errorf("contract address is required")
	}

	if c.Port == "" {
		return fmt.Errorf("port is required")
	}

	if c.EnableTLS {
		if c.TLSCertFile == "" {
			return fmt.Errorf("TLS cert file is required when TLS is enabled")
		}
		if c.TLSKeyFile == "" {
			return fmt.Errorf("TLS key file is required when TLS is enabled")
		}
	}

	if c.CacheTTL <= 0 {
		return fmt.Errorf("cache TTL must be positive")
	}

	if c.TokenMaxAge <= 0 {
		return fmt.Errorf("token max age must be positive")
	}

	return nil
}

// GetListenAddress retourne l'adresse d'écoute du serveur
func (c *Config) GetListenAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
