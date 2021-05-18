```

helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update

helm install nginx-ingress nginx-stable/nginx-ingress --create-namespace --namespace nginx-ingress

```