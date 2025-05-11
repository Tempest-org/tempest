FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /app/bin/access ./access/access.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/access /app/access

# Copy configuration
COPY --from=builder /app/access/etc /app/etc

# Expose the port
EXPOSE 8081

# Run the service
CMD ["/app/access", "-f", "/app/etc/access.yaml"]