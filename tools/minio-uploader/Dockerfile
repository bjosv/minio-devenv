FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod /build
COPY go.sum /build
COPY minio-uploader.go /build
RUN go build -ldflags '-extldflags -static' -a minio-uploader.go

FROM golang:alpine
COPY --from=builder /build/minio-uploader /app/
CMD ["/app/minio-uploader"]
