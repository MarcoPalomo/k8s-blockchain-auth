package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcopalomo/k8s-blockchain-auth/pkg/authenticator"
	"github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain"
)

// Server représente le serveur d'authentification blockchain
type Server struct {
	config        *Config
	httpServer    *http.Server
	blockchain    *blockchain.Client
	authenticator *authenticator.BlockchainAuthenticator
	authorizer    *authenticator.BlockchainAuthorizer
}

// NewServer crée un nouveau serveur avec la configuration donnée
func NewServer(cfg *Config) (*Server, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Initialiser le client blockchain
	log.Printf("Connecting to blockchain at %s...", cfg.BlockchainRPCURL)
	bcClient, err := blockchain.NewClient(cfg.BlockchainRPCURL, cfg.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain: %w", err)
	}
	log.Println("Successfully connected to blockchain")

	// Créer l'authenticator et l'authorizer
	bcAuthenticator := authenticator.NewBlockchainAuthenticator(bcClient)
	bcAuthorizer := authenticator.NewBlockchainAuthorizer(bcClient)

	// Créer le serveur HTTP
	mux := http.NewServeMux()
	httpServer := &http.Server{
		Addr:         cfg.GetListenAddress(),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	server := &Server{
		config:        cfg,
		httpServer:    httpServer,
		blockchain:    bcClient,
		authenticator: bcAuthenticator,
		authorizer:    bcAuthorizer,
	}

	// Enregistrer les routes
	server.setupRoutes(mux)

	return server, nil
}

// setupRoutes configure les routes HTTP
func (s *Server) setupRoutes(mux *http.ServeMux) {
	// Authentication endpoint
	mux.HandleFunc("/authenticate", s.handleAuthenticate)

	// Authorization endpoint
	mux.HandleFunc("/authorize", s.handleAuthorize)

	// Health checks
	mux.HandleFunc(s.config.HealthCheckPath, s.handleHealthCheck)
	mux.HandleFunc(s.config.ReadinessCheckPath, s.handleReadinessCheck)

	// Info endpoint
	mux.HandleFunc("/", s.handleInfo)
}

// handleAuthenticate gère les requêtes d'authentification
func (s *Server) handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, ok, err := s.authenticator.AuthenticateRequest(r)
	if err != nil {
		log.Printf("Authentication error: %v", err)
		http.Error(w, fmt.Sprintf("Authentication error: %v", err), http.StatusUnauthorized)
		return
	}

	if !ok {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	// Préparer la réponse
	response := map[string]interface{}{
		"authenticated": true,
		"user": map[string]interface{}{
			"name":   resp.User.GetName(),
			"uid":    resp.User.GetUID(),
			"groups": resp.User.GetGroups(),
			"extra":  resp.User.GetExtra(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// handleAuthorize gère les requêtes d'autorisation
func (s *Server) handleAuthorize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Note: L'implémentation complète nécessiterait de parser les attributs d'autorisation
	// Pour l'instant, on retourne juste un placeholder
	response := map[string]interface{}{
		"authorized": false,
		"reason":     "Not implemented yet",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleHealthCheck gère les health checks
func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

// handleReadinessCheck gère les readiness checks
func (s *Server) handleReadinessCheck(w http.ResponseWriter, r *http.Request) {
	// Vérifier la connexion blockchain
	if s.blockchain == nil {
		http.Error(w, "blockchain not connected", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ready")
}

// handleInfo affiche les informations du serveur
func (s *Server) handleInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"name":    "K8s Blockchain Auth Server",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"authenticate": "/authenticate",
			"authorize":    "/authorize",
			"health":       s.config.HealthCheckPath,
			"readiness":    s.config.ReadinessCheckPath,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// Start démarre le serveur
func (s *Server) Start() error {
	log.Printf("Starting server on %s...", s.httpServer.Addr)

	// Canal pour les signaux d'arrêt
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Démarrer le serveur dans une goroutine
	errChan := make(chan error, 1)
	go func() {
		var err error
		if s.config.EnableTLS {
			log.Printf("Starting HTTPS server with TLS cert: %s", s.config.TLSCertFile)
			err = s.httpServer.ListenAndServeTLS(s.config.TLSCertFile, s.config.TLSKeyFile)
		} else {
			log.Println("Starting HTTP server (TLS disabled)")
			err = s.httpServer.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Attendre un signal d'arrêt ou une erreur
	select {
	case err := <-errChan:
		return fmt.Errorf("server error: %w", err)
	case sig := <-stopChan:
		log.Printf("Received signal %v, shutting down...", sig)
		return s.Shutdown()
	}
}

// Shutdown arrête gracieusement le serveur
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Shutting down server...")

	// Arrêter le serveur HTTP
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	// Fermer la connexion blockchain
	if s.blockchain != nil {
		s.blockchain.Close()
	}

	log.Println("Server shutdown complete")
	return nil
}
