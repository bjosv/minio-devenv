# Operator

## Prepare (once)
```
kubectl krew update
kubectl krew install minio
```

## Install
```
# Install operator in cluster
kubectl minio init

# Create minio installation
kubectl minio tenant create minio-tenant-1 --servers 4 --volumes 4 --capacity 4Gi --namespace default --storage-class standard
```
