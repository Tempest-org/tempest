# Tempest Platform

Tempest is a modern microservices platform built with Go, designed for scalability and maintainability.

## Project Structure

```
tempest/
├── pkg/                # Shared packages used across services
│   ├── auth/           # Authentication utilities
│   ├── config/         # Configuration utilities 
│   ├── datasource/     # Database connection utilities
│   └── ...
├── services/           # Individual microservices
│   ├── accounts/       # User account management service
│   └── ...
├── scripts/            # Utility scripts
├── docker-compose.yml  # Docker Compose configuration
└── README.md           # This file
```

## Getting Started

### Prerequisites

- Go 1.24.2 or higher
- Docker and Docker Compose

### Running the Platform

1. Start the required dependencies (PostgreSQL, ETCD, and NATS) using Docker Compose:

```bash
docker-compose up -d
```

2. Run a specific service:

```bash
cd services/accounts
go run .
```

## Services

### Accounts Service

The accounts service provides user management functionality:

- User registration with token generation
- Authentication and authorization
- User profile management

## Contributing

1. Create a new branch for your feature or bugfix
2. Make your changes
3. Submit a pull request

## License

[MIT License](LICENSE)