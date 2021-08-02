---
sidebar_position: 1
---

# GoArcc - Download

To download GoArcc locally, you need to clone the repo.
```
git clone https://github.com/deqode/GoArcc.git
```
OR
```
go get github.com/deqode/GoArcc
```

# Requirements
GoArcc depends on several tools and libraries-
- Go 1.16
- [Docker](https://docs.docker.com/install/) 19.03+
- [Docker Compose](https://docs.docker.com/compose/install/) 1.25+

Below plugins are also required to generate code from a proto file. Check the usages.<br/>
- [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [protoc](https://grpc.io/docs/protoc-installation/)
- [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)
- [grpc-graphql-gateway](https://github.com/ysugimoto/grpc-graphql-gateway)

# Running the GoArcc
build the project using ``make build``<br/>

run the binary ``./bin/goarcc``

OR

Run from docker-compose ``docker-compose up``

# Hosts
GoArcc running servers hosts on different-2 ports.

- Jaeger UI: `http://localhost:16686`
- Health Trace: `http://localhost:8083/health/`
- Prometheus UI: `http://localhost:9090`
- Prometheus UI Metrics: `http://localhost:9090/metrics`
- Grpc Server: `http://localhost:8080`
- Graphql Server: `http://localhost:8081`
- Rest Server: `http://localhost:8082`

Server configurations are in config.yml