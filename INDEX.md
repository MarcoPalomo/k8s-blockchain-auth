# Index - k8s-blockchain-auth

Guide de navigation rapide pour le projet k8s-blockchain-auth.

## Documents principaux

- [README.md](README.md) - Guide principal du projet
- [REORGANIZATION_SUMMARY.md](REORGANIZATION_SUMMARY.md) - Résumé complet de la réorganisation
- [MIGRATION.md](MIGRATION.md) - Guide de migration depuis la version MVP
- [CHANGELOG.md](CHANGELOG.md) - Journal détaillé des modifications

## Démarrage rapide

1. [README.md](README.md#installation-rapide) - Installation et setup
2. [scripts/setup-dev.sh](scripts/setup-dev.sh) - Script de configuration automatique

## Code source

### Packages Go (pkg/)

#### pkg/blockchain/ - Client blockchain
- [client.go](pkg/blockchain/client.go) - Client Ethereum
- [contract.go](pkg/blockchain/contract.go) - Gestion du smart contract
- [types.go](pkg/blockchain/types.go) - Types de données

#### pkg/authenticator/ - Authentification K8s
- [blockchain_authenticator.go](pkg/authenticator/blockchain_authenticator.go) - Authentificateur
- [blockchain_authorizer.go](pkg/authenticator/blockchain_authorizer.go) - Autorisateur
- [cache.go](pkg/authenticator/cache.go) - Cache thread-safe
- [token.go](pkg/authenticator/token.go) - Gestion des tokens

#### pkg/apiserver/ - Serveur API
- [config.go](pkg/apiserver/config.go) - Configuration
- [server.go](pkg/apiserver/server.go) - Serveur HTTP
- [options.go](pkg/apiserver/options.go) - Options CLI

#### pkg/util/ - Utilitaires
- [crypto.go](pkg/util/crypto.go) - Cryptographie Ethereum
- [validation.go](pkg/util/validation.go) - Validation des entrées

### Binaires (cmd/)

#### [cmd/apiserver/](cmd/apiserver/)
Serveur d'authentification principal
- Usage: `./bin/apiserver [options]`

#### [cmd/kubectl-wallet/](cmd/kubectl-wallet/)
Plugin kubectl pour la gestion des wallets
- [main.go](cmd/kubectl-wallet/main.go) - Point d'entrée
- [auth.go](cmd/kubectl-wallet/auth.go) - Création de wallets
- [sign.go](cmd/kubectl-wallet/sign.go) - Signature et tokens

#### [cmd/permission-manager/](cmd/permission-manager/)
CLI de gestion des permissions
- [main.go](cmd/permission-manager/main.go) - Point d'entrée
- [grant.go](cmd/permission-manager/grant.go) - Accorder permissions
- [revoke.go](cmd/permission-manager/revoke.go) - Révoquer permissions
- [list.go](cmd/permission-manager/list.go) - Lister permissions

## Smart Contracts

- [contracts/K8sAccessControl.sol](contracts/K8sAccessControl.sol) - Contrat principal
- [contracts/test/](contracts/test/) - Tests des contrats
- [contracts/migrations/](contracts/migrations/) - Scripts de déploiement

## Déploiement

### Kubernetes
- [deployments/kubernetes/apiserver/](deployments/kubernetes/apiserver/) - Manifests serveur
- [deployments/kubernetes/rbac/](deployments/kubernetes/rbac/) - RBAC
- [deployments/kubernetes/examples/](deployments/kubernetes/examples/) - Exemples

### Helm
- [deployments/helm/k8s-blockchain-auth/](deployments/helm/k8s-blockchain-auth/) - Chart Helm

### Docker
- [deployments/docker/apiserver.Dockerfile](deployments/docker/apiserver.Dockerfile)
- [deployments/docker/kubectl-wallet.Dockerfile](deployments/docker/kubectl-wallet.Dockerfile)

## Scripts

- [scripts/setup-dev.sh](scripts/setup-dev.sh) - Setup environnement
- [scripts/build.sh](scripts/build.sh) - Build binaires
- [scripts/deploy-contract.sh](scripts/deploy-contract.sh) - Déployer contrat
- [scripts/test-e2e.sh](scripts/test-e2e.sh) - Tests E2E
- [scripts/generate-bindings.sh](scripts/generate-bindings.sh) - Générer bindings Go

## Tests

- [test/README.md](test/README.md) - Guide des tests
- [test/integration/](test/integration/) - Tests d'intégration
- [test/e2e/](test/e2e/) - Tests end-to-end
- [test/fixtures/](test/fixtures/) - Données de test

## Documentation

- [docs/README.md](docs/README.md) - Index documentation
- [examples/README.md](examples/README.md) - Guide exemples
- [hack/README.md](hack/README.md) - Scripts développement

## Liens utiles par cas d'usage

### Je veux démarrer rapidement
1. [README.md - Installation rapide](README.md#installation-rapide)
2. [scripts/setup-dev.sh](scripts/setup-dev.sh)

### Je veux comprendre l'architecture
1. [REORGANIZATION_SUMMARY.md](REORGANIZATION_SUMMARY.md)
2. [README.md - Architecture](README.md#architecture)

### Je migre depuis la version MVP
1. [MIGRATION.md](MIGRATION.md)
2. [CHANGELOG.md](CHANGELOG.md)

### Je veux développer
1. [README.md - Développement](README.md#développement)
2. [docs/README.md](docs/README.md)
3. [test/README.md](test/README.md)

### Je veux déployer en production
1. [README.md - Déploiement](README.md#déploiement-sur-kubernetes)
2. [deployments/helm/k8s-blockchain-auth/](deployments/helm/k8s-blockchain-auth/)
3. [scripts/deploy-contract.sh](scripts/deploy-contract.sh)

### Je veux gérer les permissions
1. [cmd/permission-manager/](cmd/permission-manager/)
2. [README.md - Gérer permissions](README.md#gérer-les-permissions)

### Je veux créer un wallet
1. [cmd/kubectl-wallet/](cmd/kubectl-wallet/)
2. [README.md - Créer wallet](README.md#créer-un-wallet)

## Navigation rapide par technologie

### Go
- [pkg/](pkg/) - Tous les packages
- [cmd/](cmd/) - Tous les binaires

### Solidity
- [contracts/K8sAccessControl.sol](contracts/K8sAccessControl.sol)
- [contracts/test/](contracts/test/)

### Kubernetes
- [deployments/kubernetes/](deployments/kubernetes/)
- [deployments/helm/](deployments/helm/)

### Shell
- [scripts/](scripts/)

## Aide et support

- Problèmes courants: [README.md - Dépannage](README.md#dépannage)
- Documentation complète: [docs/](docs/)
- Exemples: [examples/](examples/)

---

Dernière mise à jour: Janvier 2026
