.PHONY: all help setup build clean dev-start dev-stop dev-run

# Load .env if it exists
-include .env
export

# Variables
BINARY_NAME=k8s-blockchain-auth
NAMESPACE=blockchain-auth
BLOCKCHAIN_RPC_URL ?= http://localhost:8545
CONTRACT_ADDRESS ?=
NETWORK ?= localhost
WALLET_ADDRESS ?=
PRIVATE_KEY ?=

# Colors
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m

all: help

help: ## Afficher l'aide
	@echo "$(GREEN)K8s Blockchain Auth - Makefile$(NC)"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"; printf "\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2 } /^##@/ { printf "\n$(YELLOW)%s$(NC)\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Setup

setup: ## Setup complet (dépendances + build)
	@echo "$(GREEN)=== Setup ===$(NC)"
	@command -v go >/dev/null 2>&1 || { echo "$(RED)Go requis$(NC)"; exit 1; }
	@command -v node >/dev/null 2>&1 || { echo "$(RED)Node.js requis$(NC)"; exit 1; }
	@go mod download && go mod tidy
	@cd contracts && npm install
	@$(MAKE) contracts-compile
	@$(MAKE) contracts-bindings
	@$(MAKE) build
	@echo "$(GREEN)✓ Setup terminé$(NC)"

##@ Build

build: ## Compiler tous les binaires
	@mkdir -p bin
	@go build -o bin/apiserver ./cmd/apiserver
	@go build -o bin/kubectl-wallet ./cmd/kubectl-wallet
	@go build -o bin/permission-manager ./cmd/permission-manager
	@echo "$(GREEN)✓ Binaires compilés$(NC)"
	@ls -lh bin/

##@ Smart Contracts

contracts-compile: ## Compiler les smart contracts
	@cd contracts && npx hardhat compile

contracts-bindings: ## Générer les bindings Go
	@./scripts/generate-bindings.sh

contracts-deploy: ## Déployer le contract (NETWORK=localhost|sepolia)
	@cd contracts && npx hardhat run migrations/1_deploy_contracts.js --network $(NETWORK)

##@ Development

dev-start: ## Démarrer blockchain + déployer contract
	@echo "$(GREEN)=== Dev Start ===$(NC)"
	@cd contracts && npx hardhat node > /tmp/hardhat-node.log 2>&1 & echo $$! > /tmp/hardhat-node.pid
	@sleep 3
	@echo "$(GREEN)✓ Hardhat node démarré$(NC)"
	@$(MAKE) contracts-deploy NETWORK=localhost
	@echo ""
	@DEPLOYED_ADDRESS=$$(cat deployments/localhost.json | jq -r '.contractAddress'); \
	echo "$(GREEN)Contract déployé:$(NC) $$DEPLOYED_ADDRESS"; \
	./scripts/save-env.sh CONTRACT_ADDRESS $$DEPLOYED_ADDRESS
	@echo ""
	@echo "$(YELLOW)Contract address sauvegardé dans .env$(NC)"
	@echo "$(YELLOW)Next:$(NC) make dev-run"

dev-stop: ## Arrêter la blockchain
	@if [ -f /tmp/hardhat-node.pid ]; then \
		kill $$(cat /tmp/hardhat-node.pid) 2>/dev/null || true; \
		rm /tmp/hardhat-node.pid; \
		echo "$(GREEN)✓ Hardhat arrêté$(NC)"; \
	fi

dev-run: build ## Démarrer l'apiserver
	@if [ -z "$(CONTRACT_ADDRESS)" ]; then \
		echo "$(RED)CONTRACT_ADDRESS requis. Lancez d'abord: make dev-start$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN)Démarrage API server avec CONTRACT_ADDRESS=$(CONTRACT_ADDRESS)$(NC)"
	@BLOCKCHAIN_RPC_URL=$(BLOCKCHAIN_RPC_URL) CONTRACT_ADDRESS=$(CONTRACT_ADDRESS) ./bin/apiserver

##@ Wallet & Permissions

wallet-create: build ## Créer un nouveau wallet (auto-sauvegardé dans .env)
	@./scripts/wallet-create-wrapper.sh

onboard-user: build ## Onboarding interactif d'un utilisateur
	@./scripts/onboard-user.sh

onboard-app: build ## Onboarding interactif d'une application/service
	@./scripts/onboard-app.sh

wallet-grant: build ## Accorder permissions (utilise .env ou WALLET_ADDRESS)
	@if [ -z "$(WALLET_ADDRESS)" ]; then \
		echo "$(RED)WALLET_ADDRESS requis$(NC)"; \
		echo "Créez un wallet: make wallet-create"; \
		exit 1; \
	fi
	@if [ -z "$(CONTRACT_ADDRESS)" ]; then \
		echo "$(RED)CONTRACT_ADDRESS requis$(NC)"; \
		echo "Démarrez l'env dev: make dev-start"; \
		exit 1; \
	fi
	@echo "$(GREEN)Accordant permissions à $(WALLET_ADDRESS)$(NC)"
	@BLOCKCHAIN_RPC_URL=$(BLOCKCHAIN_RPC_URL) \
		CONTRACT_ADDRESS=$(CONTRACT_ADDRESS) \
		ADMIN_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
		./bin/permission-manager grant $(WALLET_ADDRESS) '*' '*' '*'

wallet-token: build ## Générer un token (utilise .env ou PRIVATE_KEY)
	@if [ -z "$(PRIVATE_KEY)" ]; then \
		echo "$(RED)PRIVATE_KEY requis$(NC)"; \
		echo "Créez un wallet: make wallet-create"; \
		exit 1; \
	fi
	@./bin/kubectl-wallet generate $(PRIVATE_KEY)

wallet-test: build ## Tester l'authentification (utilise .env ou PRIVATE_KEY)
	@if [ -z "$(PRIVATE_KEY)" ]; then \
		echo "$(RED)PRIVATE_KEY requis$(NC)"; \
		echo "Créez un wallet: make wallet-create"; \
		exit 1; \
	fi
	@./scripts/test-auth.sh $(PRIVATE_KEY)

##@ Kubernetes

k8s-deploy: ## Déployer sur K8s (utilise .env ou CONTRACT_ADDRESS)
	@if [ -z "$(CONTRACT_ADDRESS)" ]; then \
		echo "$(RED)CONTRACT_ADDRESS requis$(NC)"; \
		echo "Démarrez l'env dev: make dev-start"; \
		exit 1; \
	fi
	@echo "$(GREEN)Déploiement sur K8s avec CONTRACT_ADDRESS=$(CONTRACT_ADDRESS)$(NC)"
	@kubectl create namespace $(NAMESPACE) --dry-run=client -o yaml | kubectl apply -f -
	@kubectl create secret generic blockchain-auth-secret \
		-n $(NAMESPACE) \
		--from-literal=CONTRACT_ADDRESS=$(CONTRACT_ADDRESS) \
		--dry-run=client -o yaml | kubectl apply -f -
	@kubectl apply -f deployments/kubernetes/rbac/ -n $(NAMESPACE)
	@kubectl apply -f deployments/kubernetes/apiserver/ -n $(NAMESPACE)
	@echo "$(GREEN)✓ Déployé sur K8s$(NC)"

k8s-status: ## Statut K8s
	@kubectl get all -n $(NAMESPACE)

k8s-logs: ## Logs du pod
	@kubectl logs -f -n $(NAMESPACE) -l app=blockchain-auth-server

k8s-delete: ## Supprimer du K8s
	@kubectl delete namespace $(NAMESPACE)

##@ Tests

test: ## Tests Go
	@go test ./pkg/... -v

test-contracts: ## Tests smart contracts
	@cd contracts && npm test

##@ Utilities

clean: ## Nettoyer
	@rm -rf bin/ contracts/artifacts/ contracts/cache/
	@rm -f /tmp/hardhat-node.log /tmp/hardhat-node.pid
	@echo "$(GREEN)✓ Nettoyé$(NC)"

fmt: ## Formatter le code
	@go fmt ./...

info: ## Infos du projet et configuration
	@echo "$(GREEN)=== Configuration Actuelle ===$(NC)"
	@echo ""
	@echo "$(YELLOW)Blockchain:$(NC)"
	@echo "  RPC URL:          $(BLOCKCHAIN_RPC_URL)"
	@echo "  Network:          $(NETWORK)"
	@echo "  Contract Address: $(if $(CONTRACT_ADDRESS),$(CONTRACT_ADDRESS),$(RED)non configuré$(NC))"
	@echo ""
	@echo "$(YELLOW)Wallet:$(NC)"
	@echo "  Address:          $(if $(WALLET_ADDRESS),$(WALLET_ADDRESS),$(RED)non créé$(NC))"
	@echo "  Private Key:      $(if $(PRIVATE_KEY),***configuré***,$(RED)non configuré$(NC))"
	@echo ""
	@echo "$(YELLOW)Kubernetes:$(NC)"
	@echo "  Namespace:        $(NAMESPACE)"
	@echo ""
	@echo "$(YELLOW)Binaires:$(NC)"
	@ls -lh bin/ 2>/dev/null || echo "  $(RED)non compilés$(NC)"
	@echo ""
	@if [ -f .env ]; then \
		echo "$(GREEN)✓ Fichier .env trouvé$(NC)"; \
	else \
		echo "$(YELLOW)⚠ Fichier .env non trouvé (sera créé automatiquement)$(NC)"; \
	fi
