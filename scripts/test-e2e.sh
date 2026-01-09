#!/bin/bash
set -e

echo "Running end-to-end tests..."

# Check prerequisites
command -v kubectl >/dev/null 2>&1 || { echo "kubectl is required"; exit 1; }

echo "1. Testing wallet creation..."
./bin/kubectl-wallet create-wallet > /tmp/wallet.txt
PRIVATE_KEY=$(grep "Clé privée" /tmp/wallet.txt | awk '{print $3}')
echo "Wallet created: $PRIVATE_KEY"

echo "2. Generating auth token..."
TOKEN=$(./bin/kubectl-wallet generate "$PRIVATE_KEY")
echo "Token generated"

echo "3. Testing authentication endpoint..."
curl -X POST http://localhost:8080/authenticate \
  -H "Authorization: Bearer $TOKEN" \
  -v

echo "End-to-end tests completed!"
