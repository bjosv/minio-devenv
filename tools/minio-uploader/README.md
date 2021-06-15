# minio-uploader

Upload a bunch of files to a minio deployment

## Build

docker build -t minio-uploader:0.1.0 .

## Upload to kind

kind load docker-image minio-uploader:0.1.0
