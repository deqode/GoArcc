```
helm upgrade --install postgres bitnami/postgresql -f deployment/k8s/postgres/values.override.yaml --create-namespace --namespace postgres
```

```
PostgreSQL can be accessed via port 5432 on the following DNS name from within your cluster:

    postgres.postgres.svc.cluster.local - Read/Write connection

To get the password for "postgres" run:

    export POSTGRES_PASSWORD=$(kubectl get secret --namespace postgres postgres -o jsonpath="{.data.postgresql-password}" | base64 --decode)

To connect to your database run the following command:

    kubectl run postgres-client --rm --tty -i --restart='Never' --namespace postgres --image docker.io/bitnami/postgresql:11.12.0-debian-10-r1 --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host postgres -U postgres -d postgres -p 5432



To connect to your database from outside the cluster execute the following commands:

    kubectl port-forward --namespace postgres svc/postgres 5432:5432 &
    PGPASSWORD="$POSTGRES_PASSWORD" psql --host 127.0.0.1 -U postgres -d postgres -p 5432
```