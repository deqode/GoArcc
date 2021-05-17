```
helm repo add scylla https://scylla-operator-charts.storage.googleapis.com/stable
```
```
helm repo update
```
```
helm install scylla-operator scylla/scylla-operator --create-namespace --namespace scylla-operator
helm install scylla-manager scylla/scylla-manager --create-namespace --namespace scylla-manager
helm install scylla scylla/scylla --create-namespace --namespace scylla
```