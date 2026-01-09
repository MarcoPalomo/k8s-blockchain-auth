# Multi-stage build pour minimiser la taille de l'image

# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Installer les dépendances de build
RUN apk add --no-cache git make

WORKDIR /build

# Copier les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apiserver cmd/apiserver/main.go

# Stage 2: Runtime
FROM alpine:latest

# Installer les certificats CA pour les connexions HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copier le binaire depuis le stage de build
COPY --from=builder /build/apiserver .

# Exposer le port
EXPOSE 8080

# Créer un utilisateur non-root
RUN adduser -D -u 1000 appuser
USER appuser

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/healthz || exit 1

# Commande de démarrage
CMD ["./apiserver"]
