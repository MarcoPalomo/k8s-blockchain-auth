#!/bin/bash
# Script d'onboarding pour une application/service

set -e

echo "=== Onboarding Application/Service ==="
echo ""

# 1. Nom de l'application
read -p "Nom de l'application/service: " APPNAME

# 2. Type d'application
echo ""
echo "Type d'application:"
echo "  1) CI/CD Pipeline (deploy, update)"
echo "  2) Monitoring/Observability (lecture seule)"
echo "  3) Service Backend (accès config/secrets)"
echo "  4) Custom"
read -p "Type [1-4]: " APPTYPE

# 3. Créer le wallet
echo ""
echo "Création du wallet pour $APPNAME..."
WALLET_OUTPUT=$(./bin/kubectl-wallet create-wallet)
echo "$WALLET_OUTPUT"

WALLET_ADDRESS=$(echo "$WALLET_OUTPUT" | grep "Adresse:" | awk '{print $2}')
PRIVATE_KEY=$(echo "$WALLET_OUTPUT" | grep "Clé privée:" | awk '{print $3}')

# 4. Sauvegarder pour le CI/CD ou l'app
ONBOARD_DIR="onboarding"
mkdir -p "$ONBOARD_DIR"
ENV_FILE="$ONBOARD_DIR/${APPNAME}-credentials.env"

cat > "$ENV_FILE" << APP_EOF
# Credentials pour: $APPNAME
# Date: $(date)

# Variables d'environnement à configurer
export WALLET_ADDRESS="$WALLET_ADDRESS"
export WALLET_PRIVATE_KEY="$PRIVATE_KEY"

# Pour CI/CD (GitLab, GitHub Actions, etc.)
# Ajoutez ces secrets dans votre pipeline:
# - WALLET_ADDRESS: $WALLET_ADDRESS
# - WALLET_PRIVATE_KEY: $PRIVATE_KEY (MASKED/SECRET)
APP_EOF

echo ""
echo "✓ Credentials sauvegardées dans: $ENV_FILE"

# 5. Définir les permissions selon le type
case $APPTYPE in
  1)
    NAMESPACES="prod,staging"
    VERBS="get,list,create,update,patch"
    RESOURCES="deployments,services,configmaps"
    echo "Type: CI/CD Pipeline"
    ;;
  2)
    NAMESPACES="*"
    VERBS="get,list,watch"
    RESOURCES="pods,nodes,services,deployments"
    echo "Type: Monitoring"
    ;;
  3)
    read -p "Namespaces (séparés par virgule): " NAMESPACES
    VERBS="get,list"
    RESOURCES="configmaps,secrets"
    echo "Type: Service Backend"
    ;;
  4)
    echo ""
    read -p "Namespaces: " NAMESPACES
    read -p "Verbs: " VERBS
    read -p "Resources: " RESOURCES
    ;;
esac

# 6. Expiration (recommandé pour apps)
echo ""
read -p "Expiration (jours, 0 pour jamais) [365]: " EXPIRATION_DAYS
EXPIRATION_DAYS=${EXPIRATION_DAYS:-365}

if [ "$EXPIRATION_DAYS" -gt 0 ]; then
  EXPIRATION=$(($(date +%s) + ($EXPIRATION_DAYS * 86400)))
else
  EXPIRATION=0
fi

# 7. Accorder les permissions
echo ""
echo "Attribution des permissions..."
echo "  App: $APPNAME"
echo "  Wallet: $WALLET_ADDRESS"
echo "  Namespaces: $NAMESPACES"
echo "  Verbs: $VERBS"
echo "  Resources: $RESOURCES"
echo ""

# Charger .env si nécessaire
if [ -z "$BLOCKCHAIN_RPC_URL" ] || [ -z "$CONTRACT_ADDRESS" ]; then
  if [ -f .env ]; then
    source .env
  else
    echo "Erreur: .env non trouvé. Lancez 'make dev-start' d'abord."
    exit 1
  fi
fi

BLOCKCHAIN_RPC_URL=${BLOCKCHAIN_RPC_URL} \
CONTRACT_ADDRESS=${CONTRACT_ADDRESS} \
ADMIN_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
./bin/permission-manager grant \
  "$WALLET_ADDRESS" \
  "$NAMESPACES" \
  "$VERBS" \
  "$RESOURCES" \
  "$EXPIRATION"

# 8. Créer un exemple d'usage
USAGE_FILE="$ONBOARD_DIR/${APPNAME}-usage-example.sh"
cat > "$USAGE_FILE" << USAGE_EOF
#!/bin/bash
# Exemple d'utilisation pour $APPNAME

# Charger les credentials
source $(basename $ENV_FILE)

# Générer le token
TOKEN=\$(./bin/kubectl-wallet generate \$WALLET_PRIVATE_KEY)

# Exemple d'appel API
curl -H "Authorization: Bearer \$TOKEN" \\
  https://k8s-api.example.com/api/v1/namespaces/prod/pods

# Exemple kubectl (si configuré)
kubectl get pods -n prod
USAGE_EOF
chmod +x "$USAGE_FILE"

echo ""
echo "✅ Onboarding terminé!"
echo ""
echo "Fichiers créés:"
echo "  - Credentials: $ENV_FILE"
echo "  - Exemple: $USAGE_FILE"
echo ""
echo "Prochaines étapes:"
echo "  1. Ajoutez les variables dans votre CI/CD ou app"
echo "  2. Testez avec: source $ENV_FILE && ./bin/kubectl-wallet generate \$WALLET_PRIVATE_KEY"
echo "  3. Sécurisez et ne committez JAMAIS ces fichiers"
