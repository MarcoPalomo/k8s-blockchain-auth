#!/bin/bash
# Script d'onboarding interactif pour un utilisateur

set -e

echo "=== Onboarding Utilisateur ==="
echo ""

# 1. Demander le nom de l'utilisateur
read -p "Nom de l'utilisateur: " USERNAME

# 2. Demander le type de profil
echo ""
echo "Choisir un profil:"
echo "  1) Admin (tous les droits)"
echo "  2) Developer (accès dev + lecture prod)"
echo "  3) ReadOnly (lecture seule partout)"
echo "  4) Custom (spécifier manuellement)"
read -p "Profil [1-4]: " PROFILE

# 3. Créer le wallet
echo ""
echo "Création du wallet pour $USERNAME..."
WALLET_OUTPUT=$(./bin/kubectl-wallet create-wallet)
echo "$WALLET_OUTPUT"

WALLET_ADDRESS=$(echo "$WALLET_OUTPUT" | grep "Adresse:" | awk '{print $2}')
PRIVATE_KEY=$(echo "$WALLET_OUTPUT" | grep "Clé privée:" | awk '{print $3}')

# 4. Sauvegarder les infos dans un fichier sécurisé
ONBOARD_DIR="onboarding"
mkdir -p "$ONBOARD_DIR"
WALLET_FILE="$ONBOARD_DIR/${USERNAME}-wallet.txt"

cat > "$WALLET_FILE" << WALLET_EOF
Utilisateur: $USERNAME
Date: $(date)
Wallet Address: $WALLET_ADDRESS
Private Key: $PRIVATE_KEY

IMPORTANT: Envoyez ce fichier de manière sécurisée à l'utilisateur
et supprimez-le après confirmation de réception.
WALLET_EOF

echo ""
echo "✓ Wallet info sauvegardé dans: $WALLET_FILE"

# 5. Définir les permissions selon le profil
case $PROFILE in
  1)
    NAMESPACES="*"
    VERBS="*"
    RESOURCES="*"
    echo "Profil: Admin complet"
    ;;
  2)
    NAMESPACES="dev,staging"
    VERBS="*"
    RESOURCES="*"
    echo "Profil: Developer"
    ;;
  3)
    NAMESPACES="*"
    VERBS="get,list,watch"
    RESOURCES="*"
    echo "Profil: ReadOnly"
    ;;
  4)
    echo ""
    read -p "Namespaces (séparés par virgule, * pour tous): " NAMESPACES
    read -p "Verbs (get,list,create,update,delete,watch, * pour tous): " VERBS
    read -p "Resources (pods,services,deployments, * pour tous): " RESOURCES
    ;;
esac

# 6. Demander l'expiration
echo ""
read -p "Expiration (jours, 0 pour jamais) [0]: " EXPIRATION_DAYS
EXPIRATION_DAYS=${EXPIRATION_DAYS:-0}

if [ "$EXPIRATION_DAYS" -gt 0 ]; then
  EXPIRATION=$(($(date +%s) + ($EXPIRATION_DAYS * 86400)))
else
  EXPIRATION=0
fi

# 7. Accorder les permissions
echo ""
echo "Attribution des permissions..."
echo "  Wallet: $WALLET_ADDRESS"
echo "  Namespaces: $NAMESPACES"
echo "  Verbs: $VERBS"
echo "  Resources: $RESOURCES"
echo "  Expiration: $([ $EXPIRATION -eq 0 ] && echo 'Jamais' || date -d @$EXPIRATION)"
echo ""

# Vérifier que les variables d'environnement sont définies
if [ -z "$BLOCKCHAIN_RPC_URL" ] || [ -z "$CONTRACT_ADDRESS" ]; then
  # Charger depuis .env
  if [ -f .env ]; then
    source .env
  else
    echo "Erreur: .env non trouvé. Lancez 'make dev-start' d'abord."
    exit 1
  fi
fi

# Accorder les permissions
BLOCKCHAIN_RPC_URL=${BLOCKCHAIN_RPC_URL} \
CONTRACT_ADDRESS=${CONTRACT_ADDRESS} \
ADMIN_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
./bin/permission-manager grant \
  "$WALLET_ADDRESS" \
  "$NAMESPACES" \
  "$VERBS" \
  "$RESOURCES" \
  "$EXPIRATION"

echo ""
echo "✅ Onboarding terminé!"
echo ""
echo "Prochaines étapes:"
echo "  1. Envoyez $WALLET_FILE à $USERNAME de manière sécurisée"
echo "  2. L'utilisateur doit configurer kubectl (voir RBAC-GUIDE.md)"
echo "  3. Supprimez $WALLET_FILE après confirmation"
