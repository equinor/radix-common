[![SCM Compliance](https://scm-compliance-api.radix.equinor.com/repos/equinor/radix-common/badge)](https://developer.equinor.com/governance/scm-policy/)  

# Radix common types and utils

[Radix](https://www.radix.equinor.com) is a PaaS. This document is for Radix developers, or anyone interested in poking around..

## Development Process

The `radix-common` project follows a **trunk-based development** approach.

### 🔁 Workflow

- **External contributors** should:
  - Fork the repository
  - Create a feature branch in their fork

- **Maintainers** may create feature branches directly in the main repository.

### ✅ Merging Changes

All changes must be merged into the `main` branch using **pull requests** with **squash commits**.

The squash commit message must follow the [Conventional Commits](https://www.conventionalcommits.org/en/about/) specification.

## Release Process

Merging a pull request into `main` triggers the **Prepare release pull request** workflow.  
This workflow analyzes the commit messages to determine whether the version number should be bumped — and if so, whether it's a major, minor, or patch change.  

It then creates a pull request for releasing a new stable version (e.g. `1.2.3`):
Merging this request triggers the **Create releases and tags** workflow, which reads the version stored in `version.txt`, creates a GitHub release, and tags it accordingly.

## Contributing

Want to contribute? Read our [contributing guidelines](./CONTRIBUTING.md)

## Security

[How to handle security issues](./security.md)
