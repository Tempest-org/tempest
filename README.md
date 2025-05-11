# Accounts Microservice

This microservice provides user account management functionality.

## Features

- User registration with token generation
- More features to be added...

## Getting Started

### Prerequisites

- Go 1.24.2 or higher
- Docker and Docker Compose

### Running the Service with Dependencies

1. Start the required dependencies (ETCD and NATS) using Docker Compose:

```bash
docker-compose up -d
```

2. Run the accounts service:

```bash
cd accounts
go run accounts.go
```

### Testing

Use the provided test script to test the Register endpoint:

```bash
./test-register.sh
```

## API Reference

### Register

Registers a new user and returns access and refresh tokens.

**Request:**
```proto
message RegisterRequest {
  string username = 1;
  string email = 2;
  string phone = 3; // optional
  string password = 4;
}
```

**Response:**
```proto
message TokenResponse {
  string access_token = 1;
  string refresh_token = 2;
}
```

## Configuration

The service configuration is stored in `accounts/etc/accounts.yaml`. Make sure to update the JWT secrets for production use.