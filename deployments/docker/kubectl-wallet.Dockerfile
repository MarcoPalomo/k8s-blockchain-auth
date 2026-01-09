# kubectl-wallet Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /build

# Install dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build kubectl-wallet
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kubectl-wallet ./cmd/kubectl-wallet

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/kubectl-wallet .

# Create non-root user
RUN addgroup -g 1000 wallet && \
    adduser -D -u 1000 -G wallet wallet && \
    chown -R wallet:wallet /app

USER wallet

ENTRYPOINT ["/app/kubectl-wallet"]
