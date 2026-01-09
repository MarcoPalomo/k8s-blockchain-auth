# Kubernetes Blockchain Authentication

Système d'authentification et d'autorisation (RBAC) pour Kubernetes basé sur la blockchain Ethereum.

## 🎯 Concept

Utilisez des **wallets Ethereum comme identité** pour accéder à Kubernetes, avec des permissions stockées dans un smart contract.

**Avantages:**
- ✅ Décentralisé et auditable
- ✅ Pas de base de données centralisée
- ✅ Permissions immuables et traçables
- ✅ Cryptographie forte (ECDSA)
- ✅ Support des permissions temporaires

## 🚀 Quick Start

```bash
# 1. Installation
make setup

# 2. Démarrer blockchain locale + déployer contract
make dev-start

# 3. (Nouveau terminal) Démarrer API server
make dev-run

# 4. Onboarder un utilisateur
make onboard-user
```

**C'est tout!** Le système gère automatiquement les configurations via `.env`.

## 📚 Documentation

- **[QUICKSTART.md](QUICKSTART.md)** - Démarrage rapide et workflow seamless
- **[RBAC-GUIDE.md](RBAC-GUIDE.md)** - Guide complet RBAC et onboarding
- **[README-SEAMLESS.md](README-SEAMLESS.md)** - Détails du système de configuration automatique

## 🎯 Use Cases

### Utilisateurs

```bash
make onboard-user
# → Crée wallet
# → Assigne permissions (Admin/Dev/ReadOnly/Custom)
# → Génère credentials sécurisées
```

### Applications CI/CD

```bash
make onboard-app
# → Crée wallet pour l'app
# → Permissions spécifiques (deploy, monitoring, etc.)
# → Génère config GitLab CI / GitHub Actions
```

## 🏗️ Architecture

```
User/App → JWT Token (signed with wallet) 
         → K8s API Server (Auth Webhook)
         → Blockchain Auth Service
         → Smart Contract (Ethereum)
         → Permissions Check
         → Allow/Deny
```

## 🔐 Sécurité

- **Wallets Ethereum** comme identité unique
- **Signatures ECDSA** pour l'authentification
- **Smart contract** pour les permissions (immuable, auditable)
- **Pas de mots de passe** à gérer
- **Permissions temporaires** avec expiration
- **Principe du moindre privilège** par défaut

## 📋 Permissions RBAC

Format flexible:

```bash
Namespaces: [dev, staging, prod, *]
Verbs: [get, list, create, update, delete, watch, *]
Resources: [pods, services, deployments, configmaps, secrets, *]
ExpiresAt: timestamp Unix ou 0 (jamais)
```

Exemples de profils:
- **Admin**: `*, *, *`
- **Developer**: `dev,staging`, `*`, `*`
- **CI/CD**: `prod`, `get,list,create,update`, `deployments,services`
- **Monitoring**: `*`, `get,list,watch`, `pods,nodes,services`
- **ReadOnly**: `*`, `get,list`, `*`

## 🛠️ Commandes Principales

```bash
# Setup & Dev
make setup              # Installation complète
make dev-start          # Blockchain + contract deployment
make dev-run            # API server
make info               # Voir configuration actuelle

# Onboarding
make onboard-user       # Onboarding utilisateur (interactif)
make onboard-app        # Onboarding application (interactif)

# Wallet & Permissions
make wallet-create      # Créer wallet (sauvegarde auto dans .env)
make wallet-grant       # Accorder permissions
make wallet-test        # Tester authentification

# Kubernetes
make k8s-deploy         # Déployer sur K8s
make k8s-status         # Statut
make k8s-logs           # Logs

# Utilities
make help               # Toutes les commandes
make test               # Tests
make clean              # Nettoyer
```

## 📁 Structure du Projet

```
├── cmd/
│   ├── apiserver/           # Serveur d'authentification K8s
│   ├── kubectl-wallet/      # Plugin kubectl pour wallets
│   └── permission-manager/  # CLI gestion permissions
├── contracts/
│   ├── contracts/           # Smart contracts Solidity
│   └── migrations/          # Scripts de déploiement
├── pkg/
│   ├── authenticator/       # Logique d'authentification
│   ├── blockchain/          # Client blockchain + bindings
│   └── apiserver/           # Serveur HTTP
├── deployments/
│   └── kubernetes/          # Manifests K8s
├── scripts/
│   ├── onboard-user.sh      # Onboarding utilisateur
│   ├── onboard-app.sh       # Onboarding application
│   └── test-auth.sh         # Test authentification
├── examples/
│   ├── kubectl-config-example.yaml
│   ├── gitlab-ci-example.yml
│   └── github-actions-example.yml
├── Makefile                 # Commandes principales
├── QUICKSTART.md            # Guide démarrage rapide
└── RBAC-GUIDE.md            # Guide RBAC complet
```

## 🎓 Exemples

### Exemple 1: Developer Onboarding

```bash
# Admin
make onboard-user
# Nom: alice
# Profil: 2 (Developer)
# → Génère: onboarding/alice-wallet.txt

# Alice reçoit ses credentials
# Configure kubectl avec sa clé privée
kubectl get pods -n dev      # ✓ Autorisé
kubectl get pods -n prod     # ✓ Lecture seule
kubectl delete pod -n prod   # ✗ Refusé
```

### Exemple 2: CI/CD Pipeline

```bash
# DevOps
make onboard-app
# App: gitlab-ci
# Type: 1 (CI/CD Pipeline)
# → Génère: onboarding/gitlab-ci-credentials.env

# Ajouter WALLET_PRIVATE_KEY dans GitLab CI secrets
# Utiliser examples/gitlab-ci-example.yml
# Pipeline peut maintenant déployer automatiquement
```

### Exemple 3: Monitoring App

```bash
# SRE
make onboard-app
# App: prometheus
# Type: 2 (Monitoring)
# → Permissions lecture seule sur tout

# Dans Prometheus
TOKEN=$(kubectl-wallet generate $WALLET_PRIVATE_KEY)
# Utiliser TOKEN pour scraper les métriques K8s
```

## 🌟 Features

- ✅ **Seamless Configuration** - Variables auto-sauvegardées dans `.env`
- ✅ **Onboarding Interactif** - Scripts guidés pour users et apps
- ✅ **Smart Contract RBAC** - Permissions sur blockchain
- ✅ **Permissions Granulaires** - Par namespace, verb, resource
- ✅ **Expiration Temporaire** - Permissions avec TTL
- ✅ **Audit Trail** - Événements blockchain + logs
- ✅ **Multi-tenant** - Support de plusieurs clusters
- ✅ **CI/CD Ready** - Exemples GitLab CI et GitHub Actions

## 🔧 Développement

```bash
# Tests
make test               # Tests Go
make test-contracts     # Tests smart contracts

# Format
make fmt                # Format code Go

# Build
make build              # Compiler binaires
make contracts-compile  # Compiler smart contracts
```

## 📊 Monitoring & Audit

```bash
# Voir les logs du serveur
make k8s-logs

# Statut du cluster
make k8s-status

# Événements blockchain
# Via web3 explorer ou logs du contract
```

## 🤝 Contribution

1. Fork le projet
2. Créer une branche (`git checkout -b feature/amazing`)
3. Commit (`git commit -m 'Add amazing feature'`)
4. Push (`git push origin feature/amazing`)
5. Ouvrir une Pull Request

## 📝 License

MIT License - voir [LICENSE](LICENSE)

## 🙋 Support

- **Documentation**: Voir [QUICKSTART.md](QUICKSTART.md) et [RBAC-GUIDE.md](RBAC-GUIDE.md)
- **Issues**: GitHub Issues
- **Questions**: GitHub Discussions

## 🎯 Roadmap

- [ ] Support multi-cluster
- [ ] Dashboard web pour gestion permissions
- [ ] Support d'autres blockchains (Polygon, BSC)
- [ ] Intégration avec secrets managers
- [ ] Support ABAC (Attribute-Based Access Control)
- [ ] Helm chart officiel

---

**Made with ❤️ using Ethereum & Kubernetes**
