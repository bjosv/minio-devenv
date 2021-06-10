# MinIO devenv

https://github.com/minio/minio

## Prepare
```
Download client
wget https://dl.min.io/client/mc/release/linux-amd64/mc
chmod +x mc

helm repo add minio https://helm.min.io/
helm repo update
```

## Install distributed minio on K8s (4 K8s nodes, 4 pods default)
```
kind create cluster --config manifests/kind-config.yaml
k get all -A

helm install minio --set mode=distributed --set accessKey=minioadmin,secretKey=minioadmin minio/minio
k get all

# Check status using client
mc admin info play

# Open UI
kubectl port-forward service/minio 9000
xdg-open http://127.0.0.1:9000

# Check data
kubectl exec -it minio-786df9fcdb-6jtrr -- ls -la /export
```



## Other: Local deployment
```
# Start server and create folder to put config and buckets
wget https://dl.minio.io/server/minio/release/linux-amd64/minio
chmod +x minio
sudo ./minio server /minio

# Open UI
xdg-open http://127.0.0.1:9000
> User password: minioadmin minioadmin
> Click red plus, create a bucket, then upload a file
```

## Other: Docker deployment
```
docker run -p 9000:9000 -e "MINIO_ACCESS_KEY=minioadmin" -e "MINIO_SECRET_KEY=minioadmin" -v /tmp/minio:/data  minio/minio server /data

# Open UI
xdg-open http://127.0.0.1:9000
> User password: minioadmin minioadmin
> Click red plus, create a bucket, then upload a file
```

## Links
### Docs
Deploy in k8s
https://docs.min.io/docs/deploy-minio-on-kubernetes.html

### Run on minikube
https://gist.github.com/balamurugana/c59e868a36bb8a549fe863d22d6f0678
https://stackoverflow.com/questions/52719116/how-do-run-object-storage-minio-in-a-minikube-cluster

https://artificialintelligence.oodles.io/dev-blogs/Step-by-Step-Guide-to-Deploying-MinIO-on-Docker-and-Kubernetes

https://faun.pub/minio-object-storage-deployment-on-kubernetes-83f81fba1d03
https://blog.min.io/object_storage_as_a_service_on_minio/

### Helm
https://github.com/minio/charts
Helm charts archived, should use an operator instead
First version: 6.0.0 on Aug 8, 2020

https://github.com/helm/charts/tree/master/stable/minio
https://docs.gitlab.com/charts/charts/minio/

### Other related
https://min.io/resources/docs/CPG-MinIO-implementation-guide.pdf
