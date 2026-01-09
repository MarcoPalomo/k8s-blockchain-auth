# Quick Start - Workflow Seamless 🚀

## Configuration Automatique

Toutes les variables importantes sont **automatiquement sauvegardées** dans `.env` et chargées par le Makefile.

**Plus besoin de copier-coller!**

## Démarrage Rapide (3 étapes)

```bash
# 1. Setup initial
make setup

# 2. Démarrer blockchain + déployer contract
make dev-start
# → CONTRACT_ADDRESS sauvegardé automatiquement

# 3. (Nouveau terminal) Démarrer API server
make dev-run
# → Utilise automatiquement CONTRACT_ADDRESS de .env
```

## Test Complet (2 étapes)

```bash
# 1. Créer wallet
make wallet-create
# → WALLET_ADDRESS et PRIVATE_KEY sauvegardés automatiquement

# 2. Accorder permissions et tester
make wallet-grant
make wallet-test
# → Utilise automatiquement les valeurs de .env
```

## Voir la Configuration

```bash
make info
```

Output:
```
=== Configuration Actuelle ===

Blockchain:
  RPC URL:          http://localhost:8545
  Network:          localhost
  Contract Address: 0x0165878A594ca255338adfa4d48449f69242Eb8F

Wallet:
  Address:          0x353E10C5ba24fd84194E4c97b2ff6dba20FFFa52
  Private Key:      ***configuré***

✓ Fichier .env trouvé
```

## Commandes Principales

### Setup & Build
```bash
make setup              # Installation complète
make build              # Compiler binaires
make clean              # Nettoyer
```

### Development
```bash
make dev-start          # Blockchain + deploy contract (sauvegarde CONTRACT_ADDRESS)
make dev-stop           # Arrêter blockchain
make dev-run            # Lancer API server (utilise .env)
```

### Wallet & Permissions
```bash
make wallet-create      # Créer wallet (sauvegarde WALLET_ADDRESS + PRIVATE_KEY)
make wallet-grant       # Accorder permissions (utilise .env)
make wallet-token       # Générer token JWT (utilise .env)
make wallet-test        # Tester auth (utilise .env)
```

### Kubernetes
```bash
make k8s-deploy         # Déployer (utilise .env)
make k8s-status         # Statut
make k8s-logs           # Logs
make k8s-delete         # Supprimer
```

### Utilities
```bash
make info               # Voir configuration actuelle
make test               # Tests Go
make test-contracts     # Tests smart contracts
make help               # Toutes les commandes
```

## Terminologie

- **WALLET_ADDRESS**: Adresse publique (40 char hex)
  - Exemple: `0x353E10C5ba24fd84194E4c97b2ff6dba20FFFa52`
  
- **PRIVATE_KEY**: Clé privée (64 char hex) 
  - Exemple: `0xe4febc2ef9a03c93e527cba223391f3f92bf99c7c5b71b1e9bb8c6655df96db9`
  
- **CONTRACT_ADDRESS**: Adresse du smart contract
  - Exemple: `0x0165878A594ca255338adfa4d48449f69242Eb8F`

## Override Manuel (optionnel)

Vous pouvez toujours override si nécessaire:

```bash
# Utiliser un contract différent
make dev-run CONTRACT_ADDRESS=0x123...

# Tester avec une autre clé
make wallet-test PRIVATE_KEY=0xabc...
```

## Messages d'Erreur Utiles

Le système vous guide automatiquement:

```bash
$ make dev-run
CONTRACT_ADDRESS requis. Lancez d'abord: make dev-start

$ make wallet-grant
WALLET_ADDRESS requis
Créez un wallet: make wallet-create
```

## Fichier .env

Créé et géré automatiquement:

```bash
# Configuration Blockchain Auth

BLOCKCHAIN_RPC_URL=http://localhost:8545
CONTRACT_ADDRESS=0x0165878A594ca255338adfa4d48449f69242Eb8F
WALLET_ADDRESS=0x353E10C5ba24fd84194E4c97b2ff6dba20FFFa52
PRIVATE_KEY=0xe4febc2ef9a03c93e527cba223391f3f92bf99c7c5b71b1e9bb8c6655df96db9
NETWORK=localhost
```

**Sécurisé:** `.env` est dans `.gitignore`

## Workflow Kubernetes

```bash
# 1. Démarrer dev
make dev-start

# 2. Déployer sur K8s
make k8s-deploy
# → Utilise automatiquement CONTRACT_ADDRESS de .env

# 3. Vérifier
make k8s-status
make k8s-logs
```

## Troubleshooting

**Port 8080 occupé:**
```bash
lsof -ti:8080 | xargs kill -9
make dev-run
```

**Reset complet:**
```bash
make dev-stop
make clean
rm -f .env
make setup
make dev-start
```

## Avantages du Système Seamless

✅ Configuration automatique et persistante  
✅ Pas de copier-coller d'adresses  
✅ Messages d'erreur avec instructions  
✅ Override manuel toujours possible  
✅ Sécurisé (.env dans .gitignore)  
✅ Workflow ultra-simplifié  

---

**Pour plus de détails:** Voir [README-SEAMLESS.md](README-SEAMLESS.md)
