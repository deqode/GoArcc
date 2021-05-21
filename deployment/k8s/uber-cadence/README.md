```
helm upgrade --install cadence --set cassandra.config.cluster_size=1 banzaicloud-stable/cadence --create-namespace --namespace cadence -f deployment/k8s/cadence/values.override.yaml
```

```
ubercadence/cli:master --do simple-domain domain register --rd 10
```