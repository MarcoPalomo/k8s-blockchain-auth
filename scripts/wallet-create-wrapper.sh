#!/bin/bash
# Wrapper pour wallet-create qui sauvegarde automatiquement

OUTPUT=$(./bin/kubectl-wallet create-wallet)
echo "$OUTPUT"

# Extraire l'adresse et la clé privée
WALLET_ADDRESS=$(echo "$OUTPUT" | grep "Adresse:" | awk '{print $2}')
PRIVATE_KEY=$(echo "$OUTPUT" | grep "Clé privée:" | awk '{print $3}')

if [ -n "$WALLET_ADDRESS" ] && [ -n "$PRIVATE_KEY" ]; then
    ./scripts/save-env.sh WALLET_ADDRESS "$WALLET_ADDRESS"
    ./scripts/save-env.sh PRIVATE_KEY "$PRIVATE_KEY"
    echo ""
    echo "✓ Wallet sauvegardé dans .env"
fi
