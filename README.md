[![SCM Compliance](https://scm-compliance-api.radix.equinor.com/repos/equinor/radix-common/badge)](https://developer.equinor.com/governance/scm-policy/)

# Radix Common

A shared Go library providing common utilities for the [Radix](https://www.radix.equinor.com) platform — Equinor's Platform-as-a-Service (PaaS) for cloud-native applications.

This library is designed for Radix developers building platform services, or anyone interested in exploring the codebase.

## Table of Contents

- [Installation](#installation)
- [Packages](#packages)
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

### utils

Comprehensive utility functions for common operations.

| Package | Functions |
|---------|-----------|
| `utils/slice` | `Map()`, `Reduce()`, `Any()`, `All()`, `FindAll()`, `FindFirst()`, `FindIndex()`, `PointersOf()` |
| `utils/timewindow` | `TimeWindow` — Cron-like schedule validation (day + time range) |

```go
import "github.com/equinor/radix-common/utils/slice"

doubled := slice.Map(numbers, func(n int) int { return n * 2 })
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
