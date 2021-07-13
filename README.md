# Example Go monolith with embedded microservices and The Clean Architecture

This project shows an example of how to implement monolith with embedded
microservices (a.k.a. modular monolith). This way you'll get many upsides
of monorepo without it complexity and at same time most of upsides of
microservice architecture without some of it complexity.


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
    - [Structure of Go packages](#structure-of-go-packages)
    - [Features](#features)
- [Development](#development)
    - [Requirements](#requirements)
    - [Setup](#setup)
        - [HTTPS](#https)
    - [Usage](#usage)
        - [Cheatsheet](#cheatsheet)
- [Run](#run)
    - [Docker](#docker)
    - [Source](#source)
- [TODO](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

### Structure of Go packages

- `api/*` - definitions of own and 3rd-party (in `api/ext-*`)
  APIs/protocols and related auto-generated code
- `cmd/*` - main application(s)
- `internal/*` - packages shared by embedded microservices, e.g.:
    - `internal/config` - configuration (default values, env) shared by
      embedded microservices' subcommands and tests
    - `internal/dom` - domain types shared by microservices (Entities)
- `ms/*` - embedded microservices, with structure:
    - `internal/config` - configuration(s) (default values, env, flags) for
      microservice's subcommands and tests
    - `internal/app` - define interfaces ("ports") for The Clean
      Architecture (or "Ports and Adapters" architecture) and implements
      business-logic
    - `internal/srv/*` - adapters for served APIs/UI
    - `internal/sub` - adapter for incoming events
    - `internal/dal` - adapter for data storage
    - `internal/migrations` - DB migrations (in both SQL and Go)
    - `internal/svc/*` - adapters for accessing external services
- `pkg/*` - helper packages, not related to architecture and
  business-logic (may be later moved to own modules and/or replaced by
  external dependencies), e.g.:
    - `pkg/def/` - project-wide defaults
- `*/old/*` - contains legacy code which shouldn't be modified - this code
  is supposed to be extracted from `old/` directories (and refactored to
  follow Clean Architecture) when it'll need any non-trivial modification
  which require testing

### Features

- [X] Project structure (mostly) follows
  [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
- [X] Strict but convenient golangci-lint configuration.
- [X] Easily testable code (thanks to The Clean Architecture).
- [X] Avoids (and resists to) using global objects (to ensure embedded
  microservices won't conflict on these global objects).
- [X] CLI subcommands support using [cobra](https://github.com/spf13/cobra).
- [X] Graceful shutdown support.
- [X] Configuration defaults can be overwritten by env vars and flags.
- [X] Example JSON-RPC 2.0 over HTTP API, with CORS support.
- [X] Example gRPC API:
    - [X] External and internal APIs on different host/port.
    - [X] gRPC services with and without token-based authentication.
    - [X] API design (mostly) follows
      [Google API Design Guide](https://cloud.google.com/apis/design) and
      [Google API Improvement Proposals](https://google.aip.dev/).
- [X] Example DAL (data access layer):
    - [X] PostgreSQL 11 (secure schema usage pattern).
- [X] Example tests, both unit and integration.
- [X] Production logging using [structlog](https://github.com/powerman/structlog).
- [X] Production metrics using Prometheus.
- [X] Docker and docker-compose support.
- [X] Smart test coverage report, with optional support for coveralls.io.
- [X] Linters for Dockerfile and shell scripts.
- [X] CI/CD setup for GitHub Actions and CircleCI.

## Development

### Requirements

- Go 1.16
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+

### Setup



### Usage

To develop this project you'll need only standard tools: `go generate`,
`go test`, `go build`, `docker build`. Provided scripts are for
convenience only.


#### Cheatsheet

```sh

```

It's recommended to avoid `docker-compose down` - this command will also
remove docker's network for the project, and next `dc up -d` will create a
new network… repeat this many enough times and docker will exhaust
available networks, then you'll have to restart docker service or reboot.

## Run

### Docker

```
Production Container build
```
pack build backend_api -D backend_api \
--tag backend_api:latest \
--path backend/ \
--buildpack gcr.io/paketo-buildpacks/go \
--builder paketobuildpacks/builder:tiny \
--volume "$(pwd)/common:/common"
```

**Jaeger UI:**

http://localhost:16686

**Health Trace:**

http://localhost:8083/health/

**Prometheus UI:**

http://localhost:9090

**Prometheus UI Metrics:**

http://localhost:9090/metrics

**Grpc Server:**

http://localhost:8080

**Graphql Server:**

http://localhost:8081

**Rest Server:**

http://localhost:8082


ngrok http https://localhost:8443
```
