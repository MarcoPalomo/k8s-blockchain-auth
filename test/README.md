# Tests

Ce dossier contient tous les tests du projet.

## Structure

- `integration/` - Tests d'intégration entre composants
- `e2e/` - Tests end-to-end complets
- `fixtures/` - Données de test et fixtures

## Exécution des tests

```bash
# Tests unitaires
go test ./...

# Tests d'intégration
go test ./test/integration/...

# Tests E2E
./scripts/test-e2e.sh
```

## Couverture

Pour générer un rapport de couverture:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
