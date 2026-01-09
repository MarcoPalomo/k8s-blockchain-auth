#!/bin/bash
set -e

echo "Building k8s-blockchain-auth components..."

# Build apiserver
echo "Building apiserver..."
CGO_ENABLED=0 go build -o bin/apiserver ./cmd/apiserver

# Build kubectl-wallet
echo "Building kubectl-wallet..."
CGO_ENABLED=0 go build -o bin/kubectl-wallet ./cmd/kubectl-wallet

# Build permission-manager
echo "Building permission-manager..."
CGO_ENABLED=0 go build -o bin/permission-manager ./cmd/permission-manager

echo "Build complete! Binaries are in ./bin/"
ls -lh bin/
