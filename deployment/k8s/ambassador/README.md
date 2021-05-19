```
helm repo add datawire https://www.getambassador.io
kubectl create namespace ambassador
helm install ambassador --namespace ambassador datawire/ambassador
kubectl -n ambassador wait --for condition=available --timeout=90s deploy -lproduct=aes
```