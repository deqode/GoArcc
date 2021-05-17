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

