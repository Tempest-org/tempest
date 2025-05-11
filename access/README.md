# Access Service

This service handles RBAC (Role-Based Access Control) management using Casbin with PostgreSQL as the storage backend.

## Overview

The Access Service provides API endpoints for:

1. Checking if a subject (user) has permission to perform an action on an object in an organization
2. Granting permission to a subject to perform an action on an object in an organization
3. Revoking permission from a subject to perform an action on an object in an organization
4. Retrieving all access permissions a subject has in an organization
5. Retrieving all subjects that have access to perform an action on an object in an organization

## Architecture

The service uses the following components:

- **Casbin**: A powerful and efficient open-source access control library
- **PostgreSQL**: For persistent storage of access policies
- **pgx driver**: For PostgreSQL connectivity
- **go-zero**: As the RPC framework
- **gRPC**: For API communication

The service can operate in two modes:
1. With PostgreSQL backend (production mode)
2. With file-based storage (development/testing mode)

## Configuration

The service is configured through YAML files in the `etc` directory:

- `access.yaml`: For Docker/production deployment
- `access.local.yaml`: For local development 

The configuration includes:
- gRPC server settings
- Etcd settings for service discovery
- Database connection settings

## Database Schema

This service uses PostgreSQL with the pgx driver. The database schema includes:

- `casbin_rule` table to store access policies with the structure:
  ```
  id SERIAL PRIMARY KEY,
  ptype VARCHAR(100) NOT NULL,
  v0 VARCHAR(100),  -- organization
  v1 VARCHAR(100),  -- subject
  v2 VARCHAR(100),  -- object
  v3 VARCHAR(100),  -- action
  v4 VARCHAR(100),
  v5 VARCHAR(100)
  ```

## API Reference

### Check Access

Check if a subject has permission to perform an action on an object in an organization.

```protobuf
rpc Check(CheckAccessRequest) returns(CheckAccessResponse);
```

```bash
grpcurl -plaintext -d '{"organization_id": "org1", "subject_id": "user1", "object": "resource1", "action": "read"}' \
  localhost:8081 access.Access/Check
```

### Grant Access

Grant permission to a subject to perform an action on an object in an organization.

```protobuf
rpc Grant(GrantAccessRequest) returns(Empty);
```

```bash
grpcurl -plaintext -d '{"organization_id": "org1", "subject_id": "user1", "object": "resource1", "action": "write"}' \
  localhost:8081 access.Access/Grant
```

### Revoke Access

Revoke permission from a subject to perform an action on an object in an organization.

```protobuf
rpc Revoke(RevokeAccessRequest) returns(Empty);
```

```bash
grpcurl -plaintext -d '{"organization_id": "org1", "subject_id": "user1", "object": "resource1", "action": "write"}' \
  localhost:8081 access.Access/Revoke
```

### Get Subject Access

Get all access permissions for a subject in an organization.

```protobuf
rpc GetSubjectAccess(GetSubjectAccessRequest) returns(GetSubjectAccessResponse);
```

```bash
grpcurl -plaintext -d '{"organization_id": "org1", "subject_id": "user1"}' \
  localhost:8081 access.Access/GetSubjectAccess
```

### Get Object Subjects

Get all subjects that have access to perform an action on an object in an organization.

```protobuf
rpc GetObjectSubjects(GetObjectSubjectsRequest) returns(GetObjectSubjectsResponse);
```

```bash
grpcurl -plaintext -d '{"organization_id": "org1", "object": "resource1", "action": "read"}' \
  localhost:8081 access.Access/GetObjectSubjects
```

## Using the Client

The service provides a client package for easy integration with other services:

```go
import "github.com/tempest-org/tempest/access/accessclient"

// Create a new client
client := accessclient.NewAccessClient("localhost:8081")

// Check access
allowed, err := client.Check(ctx, "org123", "user456", "resource789", "read")
if err != nil {
    log.Fatal(err)
}
if allowed {
    fmt.Println("Access allowed")
} else {
    fmt.Println("Access denied")
}

// Grant access
err = client.Grant(ctx, "org123", "user456", "resource789", "write")
if err != nil {
    log.Fatal(err)
}

// Get all permissions for a user
permissions, err := client.GetSubjectAccess(ctx, "org123", "user456")
if err != nil {
    log.Fatal(err)
}
for _, perm := range permissions {
    fmt.Printf("Object: %s, Action: %s\n", perm.Object, perm.Action)
}
```

## Running the Service

### Local Development

```bash
# Run with local configuration
cd access
go run access.go -f etc/access.local.yaml

# Test the service
./test_access.sh
```

### Docker

```bash
# Build and run with Docker Compose
docker-compose up -d access
```