# MinIO devenv

https://github.com/minio/minio

## Prepare
```
# Download latest client
wget https://dl.min.io/client/mc/release/linux-amd64/mc
chmod +x mc

# Download old release and build (make)
https://github.com/minio/mc/releases/tag/RELEASE.2020-02-20T23-49-54Z
# Move to ~/bin/mc-2020-02-20

helm repo add minio https://helm.min.io/
helm repo update
```

## Install distributed minio on K8s (4 K8s nodes, 4 pods default)
```
kind create cluster --config manifests/kind-config.yaml
k get all -A

helm install minio --set mode=distributed --set accessKey=minioadmin,secretKey=minioadmin minio/minio
k get all

# Alternative: Install using specific tag, like
helm install minio --set image.tag=RELEASE.2020-02-20T22-51-23Z --set mode=distributed --set accessKey=minioadmin,secretKey=minioadmin minio/minio

# Alternative: Install own build minio (via repo)
make docker
kind load docker-image minio/minio:RELEASE.2020-06-22T03-12-50Z-18-g9b1876c56
helm install minio --set image.tag=RELEASE.2020-06-22T03-12-50Z-18-g9b1876c56 --set mode=distributed --set accessKey=minioadmin,secretKey=minioadmin minio/minio

# Forward ports for access
kubectl port-forward service/minio 9000 &

# Create a config for the minio client `mc`
mc alias set local http://127.0.0.1:9000 minioadmin minioadmin
# Alternative for older mc:
mc-2020-02-20 config host add local http://127.0.0.1:9000 minioadmin minioadmin
```

# Open UI
xdg-open http://127.0.0.1:9000

# Check status and data
mc admin info local
# or more details by
mc admin info --json local

kubectl exec -it pod/minio-0 -- ls -la /export

```

### Upload files

Build tool, see: tools/minio-uploader/README.md

```
kind load docker-image minio-uploader:0.1.0
kubectl create -f manifests/upload-job.yaml
mc admin info local

# Takes 12-18 min
# Latest:                                       -> Used: 8.45 GB
# RELEASE.2020-02-20T22-51-23Z: Used: 931.80 GB -> Used: 1,009.65 GB

# mc:
# latest: 9.4 GiB Used, 10 Buckets, 16,000 Objects
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

### Auto-heal
https://github.com/minio/minio/issues/6782
Jun 19, 2019 - heals namespace every 24hrs

### Docs
Deploy in k8s
https://docs.min.io/docs/deploy-minio-on-kubernetes.html

See Healing section in:
https://docs.min.io/docs/minio-server-configuration-guide.html

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
