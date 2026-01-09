#!/bin/bash

set -e

# Script pour générer les bindings Go depuis le smart contract Solidity
# Nécessite: solc (Solidity compiler) et abigen (go-ethereum)

echo "=== Génération des bindings Go pour K8sAccessControl ==="

# Couleurs pour l'output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Vérifier que les outils nécessaires sont installés
check_requirements() {
    echo "Vérification des prérequis..."

    # Check for solc using npx (works with global and local installs)
    if ! npx solc --version &> /dev/null; then
        echo -e "${RED}✗ solc n'est pas installé${NC}"
        echo "Installation: npm install -g solc"
        exit 1
    fi

    # Check for abigen in GOPATH/bin and ~/go/bin
    if ! command -v abigen &> /dev/null && ! [ -x "$HOME/go/bin/abigen" ]; then
        echo -e "${RED}✗ abigen n'est pas installé${NC}"
        echo "Installation: go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
        exit 1
    fi

    # Use abigen from ~/go/bin if not in PATH
    if ! command -v abigen &> /dev/null && [ -x "$HOME/go/bin/abigen" ]; then
        export PATH="$HOME/go/bin:$PATH"
    fi

    echo -e "${GREEN}✓ Tous les prérequis sont installés${NC}"
}

# Directories
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONTRACTS_DIR="$PROJECT_ROOT/contracts"
BUILD_DIR="$CONTRACTS_DIR/artifacts/contracts/K8sAccessControl.sol"
OUTPUT_DIR="$PROJECT_ROOT/pkg/blockchain"

echo "Project root: $PROJECT_ROOT"
echo "Contracts dir: $CONTRACTS_DIR"
echo "Output dir: $OUTPUT_DIR"

# Vérifier les prérequis
check_requirements

# Se placer dans le dossier contracts
cd "$CONTRACTS_DIR"

# Compiler les smart contracts avec Hardhat
echo -e "\n${YELLOW}Étape 1: Compilation des smart contracts...${NC}"
npx hardhat compile

if [ ! -f "$BUILD_DIR/K8sAccessControl.json" ]; then
    echo -e "${RED}✗ Le fichier de build n'a pas été généré${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Compilation réussie${NC}"

# Créer le dossier de sortie s'il n'existe pas
mkdir -p "$OUTPUT_DIR"

# Extraire l'ABI et le Bytecode
echo -e "\n${YELLOW}Étape 2: Extraction de l'ABI et du bytecode...${NC}"

ABI_FILE="$OUTPUT_DIR/K8sAccessControl.abi"
BIN_FILE="$OUTPUT_DIR/K8sAccessControl.bin"

# Utiliser jq pour extraire proprement l'ABI et le bytecode
cat "$BUILD_DIR/K8sAccessControl.json" | jq -r '.abi' > "$ABI_FILE"
cat "$BUILD_DIR/K8sAccessControl.json" | jq -r '.bytecode' | sed 's/^0x//' > "$BIN_FILE"

echo -e "${GREEN}✓ ABI et bytecode extraits${NC}"
echo "  - ABI: $ABI_FILE"
echo "  - Bytecode: $BIN_FILE"

# Générer les bindings Go avec abigen
echo -e "\n${YELLOW}Étape 3: Génération des bindings Go...${NC}"

BINDING_FILE="$OUTPUT_DIR/k8s_access_control_bindings.go"

abigen \
    --abi "$ABI_FILE" \
    --bin "$BIN_FILE" \
    --pkg blockchain \
    --type K8sAccessControl \
    --out "$BINDING_FILE"

echo -e "${GREEN}✓ Bindings Go générés${NC}"
echo "  - Output: $BINDING_FILE"

# Nettoyer les fichiers temporaires
echo -e "\n${YELLOW}Étape 4: Nettoyage...${NC}"
rm -f "$ABI_FILE" "$BIN_FILE"
echo -e "${GREEN}✓ Fichiers temporaires supprimés${NC}"

# Afficher un résumé
echo -e "\n${GREEN}================================${NC}"
echo -e "${GREEN}✓ Génération terminée avec succès!${NC}"
echo -e "${GREEN}================================${NC}"
echo ""
echo "Fichier généré: $BINDING_FILE"
echo ""
echo "Pour utiliser les bindings dans votre code Go:"
echo ""
echo "  import \"github.com/marcopalomo/k8s-blockchain-auth/pkg/blockchain\""
echo ""
echo "  client, err := ethclient.Dial(rpcURL)"
echo "  contract, err := blockchain.NewK8sAccessControl(contractAddress, client)"
echo ""
echo -e "${YELLOW}Note:${NC} Pensez à mettre à jour les imports dans client.go si nécessaire"
