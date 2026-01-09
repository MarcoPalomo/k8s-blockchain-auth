# Migration vers la nouvelle structure

Ce document décrit la réorganisation complète du projet k8s-blockchain-auth.

## Résumé des changements

Le projet a été complètement réorganisé pour suivre les meilleures pratiques de structuration Go et Kubernetes.

## Nouvelle structure

### Packages Go

#### pkg/blockchain/
- **client.go** - Client blockchain (existait déjà, amélioré)
- **contract.go** - Fonctions de gestion du smart contract (NOUVEAU)
- **types.go** - Types de données blockchain (NOUVEAU)

#### pkg/authenticator/
- **blockchain_authenticator.go** - Authentificateur (refactorisé)
- **blockchain_authorizer.go** - Autorisateur (existait déjà)
- **cache.go** - Cache thread-safe extrait (NOUVEAU)
- **token.go** - Gestion des tokens extraite (NOUVEAU)

#### pkg/apiserver/ (NOUVEAU)
- **config.go** - Configuration du serveur
- **server.go** - Serveur HTTP complet
- **options.go** - Options CLI

#### pkg/util/ (NOUVEAU)
- **crypto.go** - Utilitaires cryptographiques
- **validation.go** - Validation des entrées

### Binaires (cmd/)

#### cmd/apiserver/
- **main.go** - Refactorisé pour utiliser pkg/apiserver

#### cmd/kubectl-wallet/
- **main.go** - Point d'entrée (refactorisé)
- **auth.go** - Gestion des wallets (NOUVEAU)
- **sign.go** - Signature et tokens (NOUVEAU)

#### cmd/permission-manager/ (NOUVEAU)
- **main.go** - CLI principal
- **grant.go** - Accorder des permissions
- **revoke.go** - Révoquer des permissions
- **list.go** - Lister et vérifier les permissions

### Déploiements

#### deployments/kubernetes/
```
kubernetes/
├── apiserver/           # DÉPLACÉ de la racine
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── configmap.yaml
│   └── secret.yaml
├── rbac/               # NOUVEAU
│   ├── serviceaccount.yaml
│   ├── clusterrole.yaml
│   └── clusterrolebinding.yaml
└── examples/           # NOUVEAU
    ├── admin-wallet.yaml
    └── developer-wallet.yaml
```

#### deployments/docker/
- **apiserver.Dockerfile** - Renommé depuis Dockerfile
- **kubectl-wallet.Dockerfile** - NOUVEAU

#### deployments/helm/ (NOUVEAU)
- Structure Helm chart complète
- Chart.yaml, values.yaml, templates/

### Scripts

Scripts MVP supprimés:
- mvp-setup.sh
- deploy-minikube.sh
- test-auth.sh

Nouveaux scripts de production:
- **deploy-contract.sh** - Déploiement du smart contract
- **build.sh** - Build tous les binaires
- **test-e2e.sh** - Tests end-to-end
- **setup-dev.sh** - Setup environnement de dev (remplace mvp-setup.sh)
- **generate-bindings.sh** - Existait déjà, conservé

### Nouveaux dossiers

- **test/** - Structure de tests (integration/, e2e/, fixtures/)
- **docs/** - Documentation
- **examples/** - Exemples d'utilisation
- **hack/** - Scripts de développement

## Changements fonctionnels

### Extraction du cache
Le cache a été extrait de `blockchain_authenticator.go` vers `cache.go`:
- Cache thread-safe avec sync.RWMutex
- Nettoyage automatique des entrées expirées
- API simplifiée (Get, Set, Delete, Clear)

### Extraction des tokens
La gestion des tokens a été extraite vers `token.go`:
- TokenExtractor pour extraction depuis HTTP
- Validation complète des tokens
- Fonctions helper pour création/conversion

### Nouveau serveur API
Le serveur a été complètement refactorisé dans `pkg/apiserver/`:
- Configuration centralisée avec validation
- Support des options CLI
- Endpoints structurés (/authenticate, /authorize, /healthz, /readyz)
- Graceful shutdown
- Support TLS optionnel

### CLI permission-manager
Nouveau CLI complet pour la gestion des permissions:
- `grant` - Accorder des permissions via smart contract
- `revoke` - Révoquer des permissions
- `list` - Lister les permissions d'un wallet
- `check` - Vérifier une permission spécifique

### Utilitaires
Nouveaux utilitaires dans `pkg/util/`:
- Fonctions crypto (GenerateWallet, SignMessage, VerifySignature, etc.)
- Validation complète (wallets, namespaces, verbs, resources, etc.)
- Helpers pour manipulation hexadécimale

## Guide de migration

### Pour les utilisateurs

1. **Mettre à jour les imports Go**
   ```go
   // Ancien
   import "github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain"

   // Nouveau (inchangé mais nouvelles fonctions disponibles)
   import "github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain"
   import "github.com/marcopalomo/k8s-blockchain-auth/pkg/util"
   ```

2. **Rebuild les binaires**
   ```bash
   ./scripts/build.sh
   ```

3. **Mettre à jour les scripts de déploiement**
   - Utiliser `./scripts/setup-dev.sh` au lieu de `mvp-setup.sh`
   - Utiliser `./scripts/deploy-contract.sh` pour le déploiement

4. **Utiliser le nouveau permission-manager**
   ```bash
   # Ancien (via hardhat console)
   npx hardhat console

   # Nouveau (CLI dédié)
   ./bin/permission-manager grant <wallet> <namespaces> <verbs> <resources>
   ```

### Pour les développeurs

1. **Structure de code**
   - Mettre les types dans `pkg/*/types.go`
   - Extraire les utilitaires dans `pkg/util/`
   - Séparer la configuration du code métier

2. **Tests**
   - Tests unitaires dans `test/`
   - Tests d'intégration dans `test/integration/`
   - Tests E2E dans `test/e2e/`

3. **Documentation**
   - Documenter dans `docs/`
   - Créer des exemples dans `examples/`

## Compatibilité

### Compatible
- Tous les smart contracts existants fonctionnent sans modification
- Les wallets existants fonctionnent sans modification
- Les permissions on-chain restent valides

### Changements breaking
- L'ancien script `mvp-setup.sh` n'existe plus
- Les chemins des manifests Kubernetes ont changé
- Le Dockerfile a été renommé en apiserver.Dockerfile

## Bénéfices

1. **Organisation claire** - Structure standard Go
2. **Séparation des préoccupations** - Chaque package a un rôle précis
3. **Réutilisabilité** - Les packages sont indépendants
4. **Testabilité** - Structure de tests claire
5. **Déploiement** - Support Helm complet
6. **CLI professionnel** - permission-manager pour la gestion
7. **Documentation** - Structure docs/ dédiée
8. **Exemples** - Dossier examples/ pour les cas d'usage

## Prochaines étapes

1. Exécuter `./scripts/setup-dev.sh` pour configurer l'environnement
2. Lire la nouvelle documentation dans `docs/`
3. Essayer les exemples dans `examples/`
4. Utiliser les nouveaux scripts de production

## Support

Pour toute question sur la migration, consulter:
- README.md principal
- docs/README.md
- examples/README.md
