package apiserver

import (
	"flag"
	"time"
)

// Options représente les options en ligne de commande de l'API server
type Options struct {
	// Blockchain options
	BlockchainRPCURL string
	ContractAddress  string

	// Server options
	Port     string
	Host     string
	LogLevel string

	// TLS options
	TLSCertFile string
	TLSKeyFile  string

	// Cache options
	CacheTTL    time.Duration
	TokenMaxAge time.Duration
}

// NewOptions crée de nouvelles options avec les valeurs par défaut
func NewOptions() *Options {
	return &Options{
		Port:        "8080",
		Host:        "0.0.0.0",
		LogLevel:    "info",
		CacheTTL:    5 * time.Minute,
		TokenMaxAge: 5 * time.Minute,
	}
}

// AddFlags ajoute les flags au FlagSet
func (o *Options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.BlockchainRPCURL, "blockchain-rpc-url", o.BlockchainRPCURL,
		"URL du RPC blockchain (peut aussi être défini via BLOCKCHAIN_RPC_URL)")

	fs.StringVar(&o.ContractAddress, "contract-address", o.ContractAddress,
		"Adresse du smart contract (peut aussi être défini via CONTRACT_ADDRESS)")

	fs.StringVar(&o.Port, "port", o.Port,
		"Port d'écoute du serveur")

	fs.StringVar(&o.Host, "host", o.Host,
		"Adresse d'écoute du serveur")

	fs.StringVar(&o.LogLevel, "log-level", o.LogLevel,
		"Niveau de log (debug, info, warn, error)")

	fs.StringVar(&o.TLSCertFile, "tls-cert-file", o.TLSCertFile,
		"Fichier de certificat TLS")

	fs.StringVar(&o.TLSKeyFile, "tls-key-file", o.TLSKeyFile,
		"Fichier de clé TLS")

	fs.DurationVar(&o.CacheTTL, "cache-ttl", o.CacheTTL,
		"Durée de vie du cache (ex: 5m, 1h)")

	fs.DurationVar(&o.TokenMaxAge, "token-max-age", o.TokenMaxAge,
		"Age maximal d'un token (ex: 5m, 10m)")
}

// ApplyTo applique les options à la configuration
func (o *Options) ApplyTo(cfg *Config) error {
	// Override config with command-line flags if provided
	if o.BlockchainRPCURL != "" {
		cfg.BlockchainRPCURL = o.BlockchainRPCURL
	}

	if o.ContractAddress != "" {
		cfg.ContractAddress = o.ContractAddress
	}

	if o.Port != "" {
		cfg.Port = o.Port
	}

	if o.Host != "" {
		cfg.Host = o.Host
	}

	if o.LogLevel != "" {
		cfg.LogLevel = o.LogLevel
	}

	if o.TLSCertFile != "" {
		cfg.TLSCertFile = o.TLSCertFile
		cfg.EnableTLS = true
	}

	if o.TLSKeyFile != "" {
		cfg.TLSKeyFile = o.TLSKeyFile
		cfg.EnableTLS = true
	}

	if o.CacheTTL > 0 {
		cfg.CacheTTL = o.CacheTTL
	}

	if o.TokenMaxAge > 0 {
		cfg.TokenMaxAge = o.TokenMaxAge
	}

	return nil
}
