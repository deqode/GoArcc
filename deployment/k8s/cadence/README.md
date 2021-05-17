```
helm install cadence --set cassandra.config.cluster_size=1 banzaicloud-stable/cadence --create-namespace --namespace cadence
```