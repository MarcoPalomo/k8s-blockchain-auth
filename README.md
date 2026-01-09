# Kubernetes Blockchain Authentication

Ethereum blockchain-based authentication and authorization (RBAC) system for Kubernetes.

## 🎯 Concept

Use **Ethereum wallets as identity** to access Kubernetes, with permissions stored in a smart contract.

**Benefits:**
- ✅ Decentralized and auditable
- ✅ No centralized database
- ✅ Immutable and traceable permissions
- ✅ Strong cryptography (ECDSA)
- ✅ Support for temporary permissions

## 🚀 Quick Start

```bash
# 1. Installation
make setup

# 2. Start local blockchain + deploy contract
make dev-start

# 3. (New terminal) Start API server
make dev-run

# 4. Onboard a user
make onboard-user
```

**That's it!** The system automatically manages configuration via `.env`.

## 📚 Documentation

- **[QUICKSTART.md](QUICKSTART.md)** - Quick start and seamless workflow
- **[RBAC-GUIDE.md](RBAC-GUIDE.md)** - Complete RBAC and onboarding guide
- **[README-SEAMLESS.md](README-SEAMLESS.md)** - Automatic configuration system details

## 🎯 Use Cases

### Users

```bash
make onboard-user
# → Creates wallet
# → Assigns permissions (Admin/Dev/ReadOnly/Custom)
# → Generates secure credentials
```

### CI/CD Applications

```bash
make onboard-app
# → Creates wallet for the app
# → Specific permissions (deploy, monitoring, etc.)
# → Generates GitLab CI / GitHub Actions config
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

## 🔐 Security

- **Ethereum wallets** as unique identity
- **ECDSA signatures** for authentication
- **Smart contract** for permissions (immutable, auditable)
- **No passwords** to manage
- **Temporary permissions** with expiration
- **Least privilege principle** by default

## 📋 RBAC Permissions

Flexible format:

```bash
Namespaces: [dev, staging, prod, *]
Verbs: [get, list, create, update, delete, watch, *]
Resources: [pods, services, deployments, configmaps, secrets, *]
ExpiresAt: Unix timestamp or 0 (never)
```

Example profiles:
- **Admin**: `*, *, *`
- **Developer**: `dev,staging`, `*`, `*`
- **CI/CD**: `prod`, `get,list,create,update`, `deployments,services`
- **Monitoring**: `*`, `get,list,watch`, `pods,nodes,services`
- **ReadOnly**: `*`, `get,list`, `*`

## 🛠️ Main Commands

```bash
# Setup & Dev
make setup              # Complete installation
make dev-start          # Blockchain + contract deployment
make dev-run            # API server
make info               # View current configuration

# Onboarding
make onboard-user       # User onboarding (interactive)
make onboard-app        # Application onboarding (interactive)

# Wallet & Permissions
make wallet-create      # Create wallet (auto-saved to .env)
make wallet-grant       # Grant permissions
make wallet-test        # Test authentication

# Kubernetes
make k8s-deploy         # Deploy to K8s
make k8s-status         # Status
make k8s-logs           # Logs

# Utilities
make help               # All commands
make test               # Tests
make clean              # Clean
```

## 📁 Project Structure

```
├── cmd/
│   ├── apiserver/           # K8s authentication server
│   ├── kubectl-wallet/      # kubectl plugin for wallets
│   └── permission-manager/  # Permissions management CLI
├── contracts/
│   ├── contracts/           # Solidity smart contracts
│   └── migrations/          # Deployment scripts
├── pkg/
│   ├── authenticator/       # Authentication logic
│   ├── blockchain/          # Blockchain client + bindings
│   └── apiserver/           # HTTP server
├── deployments/
│   └── kubernetes/          # K8s manifests
├── scripts/
│   ├── onboard-user.sh      # User onboarding
│   ├── onboard-app.sh       # Application onboarding
│   └── test-auth.sh         # Authentication test
├── examples/
│   ├── kubectl-config-example.yaml
│   ├── gitlab-ci-example.yml
│   └── github-actions-example.yml
├── Makefile                 # Main commands
├── QUICKSTART.md            # Quick start guide
└── RBAC-GUIDE.md            # Complete RBAC guide
```

## 🎓 Examples

### Example 1: Developer Onboarding

```bash
# Admin
make onboard-user
# Name: alice
# Profile: 2 (Developer)
# → Generates: onboarding/alice-wallet.txt

# Alice receives credentials
# Configures kubectl with her private key
kubectl get pods -n dev      # ✓ Allowed
kubectl get pods -n prod     # ✓ Read-only
kubectl delete pod -n prod   # ✗ Denied
```

### Example 2: CI/CD Pipeline

```bash
# DevOps
make onboard-app
# App: gitlab-ci
# Type: 1 (CI/CD Pipeline)
# → Generates: onboarding/gitlab-ci-credentials.env

# Add WALLET_PRIVATE_KEY to GitLab CI secrets
# Use examples/gitlab-ci-example.yml
# Pipeline can now deploy automatically
```

### Example 3: Monitoring App

```bash
# SRE
make onboard-app
# App: prometheus
# Type: 2 (Monitoring)
# → Read-only permissions everywhere

# In Prometheus
TOKEN=$(kubectl-wallet generate $WALLET_PRIVATE_KEY)
# Use TOKEN to scrape K8s metrics
```

## 🌟 Features

- ✅ **Seamless Configuration** - Variables auto-saved to `.env`
- ✅ **Interactive Onboarding** - Guided scripts for users and apps
- ✅ **Smart Contract RBAC** - Permissions on blockchain
- ✅ **Granular Permissions** - By namespace, verb, resource
- ✅ **Temporary Expiration** - Permissions with TTL
- ✅ **Audit Trail** - Blockchain events + logs
- ✅ **Multi-tenant** - Multiple cluster support
- ✅ **CI/CD Ready** - GitLab CI and GitHub Actions examples

## 🔧 Development

```bash
# Tests
make test               # Go tests
make test-contracts     # Smart contract tests

# Format
make fmt                # Format Go code

# Build
make build              # Compile binaries
make contracts-compile  # Compile smart contracts
```

## 📊 Monitoring & Audit

```bash
# View server logs
make k8s-logs

# Cluster status
make k8s-status

# Blockchain events
# Via web3 explorer or contract logs
```

## 🤝 Contributing

1. Fork the project
2. Create a branch (`git checkout -b feature/amazing`)
3. Commit (`git commit -m 'Add amazing feature'`)
4. Push (`git push origin feature/amazing`)
5. Open a Pull Request

## 📝 License

MIT License - see [LICENSE](LICENSE)

## 🙋 Support

- **Documentation**: See [QUICKSTART.md](QUICKSTART.md) and [RBAC-GUIDE.md](RBAC-GUIDE.md)
- **Issues**: GitHub Issues
- **Questions**: GitHub Discussions

## 🎯 Roadmap

- [ ] Multi-cluster support
- [ ] Web dashboard for permission management
- [ ] Support for other blockchains (Polygon, BSC)
- [ ] Integration with secrets managers
- [ ] ABAC support (Attribute-Based Access Control)
- [ ] Official Helm chart

---

**Made with ❤️ using Ethereum & Kubernetes**
