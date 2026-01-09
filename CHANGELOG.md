# Changelog

## [2.0.0] - Réorganisation complète du projet

### Architecture

#### Ajouté
- **pkg/blockchain/**
  - `contract.go` - Fonctions de gestion du smart contract
  - `types.go` - Types de données blockchain

- **pkg/authenticator/**
  - `cache.go` - Cache thread-safe pour les utilisateurs authentifiés
  - `token.go` - Extraction et validation des tokens

- **pkg/apiserver/** (nouveau package)
  - `config.go` - Configuration centralisée du serveur
  - `server.go` - Serveur HTTP avec graceful shutdown
  - `options.go` - Options CLI

- **pkg/util/** (nouveau package)
  - `crypto.go` - Utilitaires cryptographiques Ethereum
  - `validation.go` - Validation des entrées utilisateur

- **cmd/permission-manager/** (nouveau binaire)
  - `main.go` - CLI principal
  - `grant.go` - Accorder des permissions
  - `revoke.go` - Révoquer des permissions
  - `list.go` - Lister et vérifier les permissions

#### Modifié
- **pkg/blockchain/client.go**
  - Ajout du champ `contractAddress` dans la struct `Client`
  - Méthodes `GetChainID()` et `GetContractAddress()`

- **pkg/authenticator/blockchain_authenticator.go**
  - Refactorisé pour utiliser `cache.go` et `token.go`
  - Suppression du code dupliqué

- **cmd/apiserver/main.go**
  - Refactorisé pour utiliser `pkg/apiserver`
  - Support des options CLI

- **cmd/kubectl-wallet/**
  - Séparé en plusieurs fichiers: `main.go`, `auth.go`, `sign.go`
  - Ajout de commandes `verify` et `info`

- **contracts/hardhat.config.js**
  - Mise à jour du chemin sources: `.` au lieu de `./contracts`
  - Ajout de la configuration `migrations`

### Déploiements

#### Ajouté
- **deployments/kubernetes/rbac/** (nouveau)
  - `serviceaccount.yaml`
  - `clusterrole.yaml`
  - `clusterrolebinding.yaml`

- **deployments/kubernetes/examples/** (nouveau)
  - `admin-wallet.yaml` - Exemple de wallet admin
  - `developer-wallet.yaml` - Exemple de wallet développeur

- **deployments/helm/k8s-blockchain-auth/** (nouveau)
  - Chart Helm complet
  - `Chart.yaml`, `values.yaml`, `templates/`

- **deployments/docker/**
  - `kubectl-wallet.Dockerfile` - Dockerfile pour kubectl-wallet

#### Modifié
- **deployments/kubernetes/**
  - Déplacement des manifests dans `apiserver/`
  - Structure organisée par composant

- **deployments/docker/Dockerfile**
  - Renommé en `apiserver.Dockerfile`

### Scripts

#### Ajouté
- `scripts/deploy-contract.sh` - Déploiement du smart contract
- `scripts/build.sh` - Build tous les binaires
- `scripts/test-e2e.sh` - Tests end-to-end
- `scripts/setup-dev.sh` - Setup environnement de développement

#### Supprimé
- `scripts/mvp-setup.sh` - Remplacé par `setup-dev.sh`
- `scripts/deploy-minikube.sh` - Fonctionnalité intégrée dans d'autres scripts
- `scripts/test-auth.sh` - Remplacé par `test-e2e.sh`

#### Conservé
- `scripts/generate-bindings.sh` - Inchangé

### Documentation

#### Ajouté
- **docs/** (nouveau dossier)
  - `README.md` - Index de la documentation

- **examples/** (nouveau dossier)
  - `README.md` - Guide des exemples

- **test/** (nouveau dossier)
  - `README.md` - Guide des tests
  - Sous-dossiers: `integration/`, `e2e/`, `fixtures/`

- **hack/** (nouveau dossier)
  - `README.md` - Scripts de développement

- `MIGRATION.md` - Guide de migration détaillé
- `CHANGELOG.md` - Ce fichier

#### Modifié
- `README.md` - Mise à jour complète avec la nouvelle structure

### Fonctionnalités

#### Ajouté
- Cache thread-safe avec nettoyage automatique
- Validation complète des entrées utilisateur
- Support TLS pour le serveur API
- Graceful shutdown du serveur
- CLI dédié pour la gestion des permissions
- Support Helm pour le déploiement
- Commandes `verify` et `info` pour kubectl-wallet
- Configuration via environnement et CLI
- Health checks et readiness probes

#### Amélioré
- Séparation des préoccupations
- Testabilité du code
- Réutilisabilité des packages
- Documentation du code
- Gestion des erreurs
- Logging structuré

### Breaking Changes

#### API
- Les chemins des manifests Kubernetes ont changé
- Le Dockerfile principal a été renommé
- Les anciens scripts MVP ont été supprimés

#### Configuration
- Nouvelles variables d'environnement supportées
- Options CLI disponibles pour apiserver

### Déprécations

- `scripts/mvp-setup.sh` - Utiliser `scripts/setup-dev.sh`
- `scripts/deploy-minikube.sh` - Utiliser kubectl ou Helm
- `scripts/test-auth.sh` - Utiliser `scripts/test-e2e.sh`

### Migrations

Pour migrer depuis la version précédente:

1. Rebuilder tous les binaires: `./scripts/build.sh`
2. Mettre à jour les chemins des manifests Kubernetes
3. Utiliser les nouveaux scripts dans `scripts/`
4. Consulter `MIGRATION.md` pour plus de détails

### Compatibilité

- ✅ Smart contracts: Aucun changement
- ✅ Wallets: Compatibles
- ✅ Permissions on-chain: Compatibles
- ⚠️ Scripts de déploiement: Nouveaux chemins
- ⚠️ Manifests K8s: Nouvelle organisation

### Dépendances

Aucun changement dans les dépendances Go ou Node.js.

### Performance

- Cache optimisé avec cleanup automatique
- Validation plus rapide grâce à la séparation du code
- Meilleure gestion de la concurrence

### Sécurité

- Validation stricte des entrées
- Support TLS
- Meilleure gestion des clés privées
- Séparation des responsabilités

### Tests

- Structure de tests clarifiée
- Tests unitaires, d'intégration et E2E séparés
- Fixtures pour les tests

### Documentation

- Documentation complète dans `docs/`
- Exemples dans `examples/`
- Guide de migration détaillé
- README mis à jour

---

## [1.0.0] - Version MVP initiale

### Fonctionnalités initiales

- Smart contract K8sAccessControl
- Serveur d'authentification Go
- Plugin kubectl-wallet
- Déploiement Kubernetes basique
- Scripts MVP

---

Pour plus de détails sur la migration, consultez [MIGRATION.md](MIGRATION.md).
