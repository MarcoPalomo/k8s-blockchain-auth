#!/bin/bash
set -e

echo "Deploying K8sAccessControl smart contract..."

cd contracts

# Check if .env exists
if [ ! -f .env ]; then
    echo "Error: .env file not found in contracts directory"
    echo "Please create .env with PRIVATE_KEY and RPC URLs"
    exit 1
fi

# Deploy to local network (default)
NETWORK="${1:-localhost}"

echo "Deploying to network: $NETWORK"

npx hardhat run migrations/1_deploy_contracts.js --network "$NETWORK"

echo "Contract deployed successfully!"
echo "Don't forget to update CONTRACT_ADDRESS in your environment"
