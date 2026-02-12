[![SCM Compliance](https://scm-compliance-api.radix.equinor.com/repos/equinor/radix-common/badge)](https://developer.equinor.com/governance/scm-policy/)

# Radix Common

A shared Go library providing common types, utilities, and framework integrations for the [Radix](https://www.radix.equinor.com) platform — Equinor's Platform-as-a-Service (PaaS) for cloud-native applications.

This library is designed for Radix developers building platform services, or anyone interested in exploring the codebase.

## Table of Contents

- [Installation](#installation)
- [Packages](#packages)
  - [models](#models)
  - [net](#net)
  - [pkg](#pkg)
  - [utils](#utils)
- [Development](#development)
  - [Development Process](#development-process)
  - [Release Process](#release-process)
- [Contributing](#contributing)
- [Security](#security)

## Installation

```bash
go get github.com/equinor/radix-common
```

## Packages

### models

Core domain models for authentication and HTTP request routing.

| Type | Description |
|------|-------------|
| `Accounts` | Holds user token and impersonation details for Kubernetes API access |
| `Impersonation` | User and group information for K8s impersonation |
| `Controller` | Interface pattern for REST/stream controllers |
| `Route` / `Routes` | Route definitions with path, method, and handler |
| `RadixHandlerFunc` | Handler function signature accepting Accounts, ResponseWriter, and Request |

```go
import "github.com/equinor/radix-common/models"

accounts := models.NewAccounts(bearerToken, impersonation, inClusterClient, outClusterClient)
upn, err := accounts.GetUserAccountUserPrincipleName()
```

### net

HTTP utilities for request/response handling and middleware.

**`net/http`** — Request parsing and response formatting:
- `GetBearerTokenFromHeader()` — Extract JWT from Authorization header
- `GetImpersonationFromHeader()` — Parse Impersonate-User/Group headers
- `JSONResponse()`, `StringResponse()`, `ByteArrayResponse()` — Response writers
- `ErrorResponse()` — Maps errors to HTTP status codes

**`net/radix_middleware.go`** — Middleware for authentication and CORS:
- `RadixMiddleware` — Extracts bearer tokens and impersonation from headers
- Sets CORS headers and manages authentication flow

```go
import radixhttp "github.com/equinor/radix-common/net/http"

token, err := radixhttp.GetBearerTokenFromHeader(request)
radixhttp.JSONResponse(writer, request, data)
```

### pkg

Framework integrations for common Go libraries.

| Package | Description |
|---------|-------------|
| `pkg/gin` | Zerolog middleware for Gin — request logging with unique request IDs |
| `pkg/gorm` | Zerolog logger for GORM — SQL query logging with elapsed time |
| `pkg/docker` | Docker registry auth config models for Kubernetes secrets |

```go
import "github.com/equinor/radix-common/pkg/gin"

router.Use(gin.SetZerologLogger(logger))
router.Use(gin.ZerologRequestLogger())
```

### utils

Comprehensive utility functions for common operations.

| Package | Functions |
|---------|-----------|
| `utils/slice` | `Map()`, `Reduce()`, `Any()`, `All()`, `FindAll()`, `FindFirst()`, `FindIndex()` |
| `utils/pointers` | `Ptr[T]()`, `Val[T]()` — Generic pointer/value conversion |
| `utils/maps` | `GetKeysFromMap()`, `MergeMaps()`, `FromString()`, `ToString()` |
| `utils/json` | `Save()`, `Load()`, `Pretty()` — Thread-safe JSON file I/O |
| `utils/timewindow` | `TimeWindow` — Cron-like schedule validation (day + time range) |
| `utils/errors` | Custom error utilities |
| `utils` | String helpers, validation, random generation, time utilities |

```go
import "github.com/equinor/radix-common/utils/slice"
import "github.com/equinor/radix-common/utils/pointers"

doubled := slice.Map(numbers, func(n int) int { return n * 2 })
ptr := pointers.Ptr("value")
val := pointers.Val(ptr)
```

## Development

### Development Process

The `radix-common` project follows a **trunk-based development** approach.

#### Workflow

- **External contributors** should:
  - Fork the repository
  - Create a feature branch in their fork

- **Maintainers** may create feature branches directly in the main repository.

#### Merging Changes

All changes must be merged into the `main` branch using **pull requests** with **squash commits**.

The squash commit message must follow the [Conventional Commits](https://www.conventionalcommits.org/en/about/) specification.

### Release Process

Merging a pull request into `main` triggers the **Prepare release pull request** workflow.
This workflow analyzes commit messages to determine version bumps (major, minor, or patch).

It creates a pull request for the new stable version (e.g., `1.2.3`).
Merging this request triggers the **Create releases and tags** workflow, which reads the version from `version.txt`, creates a GitHub release, and tags it accordingly.

## Contributing

Want to contribute? Read our [contributing guidelines](./CONTRIBUTING.md).

## Security

[How to handle security issues](./SECURITY.md)
