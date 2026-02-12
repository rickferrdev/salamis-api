# syntax=docker/dockerfile:1

# Builder stage
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
# CGO_ENABLED=0 is important for creating statically linked binaries
# -ldflags="-s -w" reduces the size of the binary by removing debug info
# The application binary is named 'salamis-api' and placed in /usr/local/bin
RUN CGO_ENABLED=0 go build -o /usr/local/bin/salamis-api ./cmd

# Final stage
FROM alpine/git AS runtime_git
FROM alpine:latest AS runtime

# Copy CA certificates for HTTPS
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

# Set up the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /usr/local/bin/salamis-api /usr/local/bin/salamis-api

# Copy the .env.example as a placeholder. In production, environment variables
# should be passed directly or managed via a secret management system.
COPY .env.example .env.example

# Expose the port your application listens on (e.g., 8080 for Gin)
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/usr/local/bin/salamis-api"]
