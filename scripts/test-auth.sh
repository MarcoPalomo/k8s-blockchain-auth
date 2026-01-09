#!/bin/bash

# Script pour tester l'authentification

if [ -z "$1" ]; then
    echo "Usage: $0 <wallet-private-key>"
    exit 1
fi

WALLET_KEY="$1"
API_URL="${API_URL:-http://localhost:8080/authenticate}"

echo "=== Test d'authentification ==="
echo ""

# Générer le token
echo "1. Génération du token..."
TOKEN=$(./bin/kubectl-wallet generate "$WALLET_KEY")

if [ -z "$TOKEN" ]; then
    echo "❌ Erreur: impossible de générer le token"
    exit 1
fi

echo "✓ Token généré"
echo "$TOKEN" | jq .
echo ""

# Tester l'authentification
echo "2. Test de l'authentification..."
RESPONSE=$(curl -s -X POST \
    -H "Authorization: Bearer $TOKEN" \
    "$API_URL")

if [ -z "$RESPONSE" ]; then
    echo "❌ Erreur: pas de réponse du serveur"
    exit 1
fi

echo "$RESPONSE" | jq .

# Vérifier le résultat
if echo "$RESPONSE" | jq -e '.authenticated == true' > /dev/null 2>&1; then
    echo ""
    echo "✅ Authentification réussie!"
else
    echo ""
    echo "❌ Authentification échouée"
    exit 1
fi
