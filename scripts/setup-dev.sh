#!/bin/bash
set -e

echo "Setting up k8s-blockchain-auth development environment..."

# Check prerequisites
echo "Checking prerequisites..."
command -v go >/dev/null 2>&1 || { echo "Go is required"; exit 1; }
command -v npm >/dev/null 2>&1 || { echo "npm is required"; exit 1; }
command -v docker >/dev/null 2>&1 || { echo "Docker is required"; exit 1; }

# Install Go dependencies
echo "Installing Go dependencies..."
go mod download
go mod tidy

# Install contract dependencies
echo "Installing smart contract dependencies..."
cd contracts
npm install
cd ..

# Generate Go bindings from smart contract
echo "Generating Go bindings..."
./scripts/generate-bindings.sh

# Build binaries
echo "Building binaries..."
./scripts/build.sh

# Create bin directory
mkdir -p bin

echo ""
echo "Development environment setup complete!"
echo ""
echo "Next steps:"
echo "1. Start a local blockchain (e.g., hardhat node)"
echo "2. Deploy the contract: ./scripts/deploy-contract.sh localhost"
echo "3. Set environment variables:"
echo "   export BLOCKCHAIN_RPC_URL=http://localhost:8545"
echo "   export CONTRACT_ADDRESS=<deployed_address>"
echo "4. Run the apiserver: ./bin/apiserver"
