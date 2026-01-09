#!/bin/bash
# Script pour sauvegarder les variables d'environnement

ENV_FILE=".env"
KEY="$1"
VALUE="$2"

if [ -z "$KEY" ] || [ -z "$VALUE" ]; then
    echo "Usage: $0 <key> <value>"
    exit 1
fi

# Créer .env s'il n'existe pas
if [ ! -f "$ENV_FILE" ]; then
    cp .env.template "$ENV_FILE"
fi

# Mettre à jour ou ajouter la variable
if grep -q "^${KEY}=" "$ENV_FILE"; then
    # Remplacer la ligne existante
    sed -i "s|^${KEY}=.*|${KEY}=${VALUE}|" "$ENV_FILE"
else
    # Ajouter la nouvelle ligne
    echo "${KEY}=${VALUE}" >> "$ENV_FILE"
fi

echo "✓ ${KEY} sauvegardé dans .env"
