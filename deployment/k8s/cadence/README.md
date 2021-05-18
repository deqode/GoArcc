```
helm install cadence --set cassandra.config.cluster_size=1 banzaicloud-stable/cadence --create-namespace --namespace cadence
```

```
ubercadence/cli:master --do simple-domain domain register --rd 10
```