```
helm upgrade --install backend deployment/k8s/backend/ \
    --set image.repository=gcr.io/alfred-dev-314111/backend \
    --set image.tag=v2 \
    --create-namespace \
    --namespace development

```

```
kubectl create secret tls ingress-tls \
    --namespace development \
    --key tls.key \
    --cert tls.crt
```