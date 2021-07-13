# Go monolith with embedded microservices including GRPC,REST, graphQL and The Clean Architecture.

When you start writing a Go project, GoArcc helps to set up all the initial code boilerplate for your project. Initial boilerplate code means how you should organize your codebase and how you can write multiple services. We have support for REST, Graphql as well as gRPC.

We have implemented logging, tracing, health check, Jaeger, etc so that any developer will come and write services within a minute.

In short, GoArcc is a boilerplate setup codebase for any monolithic(Architecture) based web/mobile applications which later converted into microservices(Architecture).


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
    - [Structure of Go packages](#structure-of-go-packages)
    - [Features](#features)
- [Development](#development)
    - [Requirements](#requirements)
   
    - [Usage](#usage)
        - [Cheatsheet](#cheatsheet)
- [Run](#run)
    - [Docker](#docker)
    - [Source](#source)
- [TODO](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

### Structure of Go packages



### Features

- [X] Project structure (mostly) follows
  [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
- [X] Strict but convenient golangci-lint configuration.
- [X] Easily testable code (thanks to The Clean Architecture).
- [X] Avoids (and resists to) using global objects (to ensure embedded
  microservices won't conflict on these global objects).
- [X] Graceful shutdown support.
- [X] Configuration defaults can be overwritten by env vars and flags.
- [X] Example gRPC API:
    - [X] External and internal APIs on different host/port.
    - [X] gRPC services with and without token-based authentication.
    - [X] API design (mostly) follows
      [Google API Design Guide](https://cloud.google.com/apis/design) and
      [Google API Improvement Proposals](https://google.aip.dev/).
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
