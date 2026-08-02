# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## AI Skills

Follow the practices defined in `~/Projects/SiteNetSoft/ai-skills/`:
- `dev-practices/golang/` — Go style, error handling, functions, testing, linting
- `dev-practices/git/` — Git authorship rules, multi-repo workspace patterns

## Project Overview

Doorman is a CLI tool for managing secrets within the Amadla ecosystem. It discovers `doorman-*` plugins on PATH and delegates secret retrieval to the appropriate plugin. Each plugin handles a specific secret store (KeePassXC, Vault, Bitwarden, etc.).

**Ecosystem context:** Doorman manages Secret entities (schema in `Entities/Secret/`). It operates alongside the main pipeline (`raise` -> `lay` -> `enjoin` -> `weaver` -> `waiter`), providing secrets to any tool that needs them.

## Build Commands

```bash
make build              # Build for current platform
make test               # Run tests with coverage
make lint               # Run golangci-lint
make clean              # Remove build artifacts
```

## Architecture

**UNIX Plugin Protocol:**
```
doorman get <key> --from <plugin>    →  doorman-<plugin> get <key>
doorman plugins                      →  scans PATH for doorman-* binaries
```

**Package Structure:**
- `main.go` - CLI entry point (Cobra)
- `plugin/` - Plugin discovery and execution (PATH scanning, subprocess delegation)
- `cmd/` - CLI commands (get, plugins, settings)

**Plugin Protocol (doorman-* binaries):**
- `info` subcommand → JSON metadata (name, version, engine, description)
- `get <key>` subcommand → retrieves secret value, outputs to stdout
- Exit codes: 0 success, 1 failure, 2 usage error
- Data to stdout, diagnostics to stderr
