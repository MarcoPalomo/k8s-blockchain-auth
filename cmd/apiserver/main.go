package main

import (
	"flag"
	"log"
	"os"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/apiserver"
)

func main() {
	log.Println("Starting K8s Blockchain Authentication Server...")

	// Parse command-line flags
	opts := apiserver.NewOptions()
	fs := flag.NewFlagSet("apiserver", flag.ExitOnError)
	opts.AddFlags(fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatalf("Failed to parse flags: %v", err)
	}

	// Create configuration
	cfg := apiserver.NewConfig()

	// Load from environment variables
	if err := cfg.LoadFromEnv(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Apply command-line options
	if err := opts.ApplyTo(cfg); err != nil {
		log.Fatalf("Failed to apply options: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Create and start server
	server, err := apiserver.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start server (blocks until shutdown)
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
