# Tempest Platform Justfile
# This file provides common development commands

# List available commands
default:
    @just --list

# Run all dependencies in Docker
dev-deps:
    docker-compose up -d

# Stop all dependencies
stop-deps:
    docker-compose down

# Run the accounts service
accounts-run:
    cd services/accounts && go run accounts.go

# Build all services
build-all:
    cd services/accounts && go build -o ../../bin/accounts

# Run all tests
test-all:
    cd services/accounts && go test -v ./...

# Update all Go dependencies
deps-update:
    go get -u ./...
    go mod tidy

# Generate protobuf files
proto-gen:
    cd services/accounts && protoc --go_out=. --go-grpc_out=. accounts.proto

# Format all Go code
fmt:
    gofmt -w .

# Check code with go vet
vet:
    go vet ./...