package main

import (
	"context"

	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := getEnv("ENDPOINT", "")
	accessKeyID := getEnv("ACCESS_KEY", "")
	secretAccessKey := getEnv("SECRET_KEY", "")
	useSSL := getEnvBool("USE_SSL", true)

	rand.Seed(time.Now().UnixNano())

	// Initialize minio client
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Create buckets
	buckets := make([]string, 0, 10)
	for i := getEnvInt("CREATED_BUCKETS", 1); i > 0; i-- {
		bucketName := fmt.Sprintf("bucket-%d", i)
		log.Printf("Create bucket: %s\n", bucketName)

		ctx := context.Background()
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			exists, errBucketExists := client.BucketExists(ctx, bucketName)
			if errBucketExists == nil && exists {
				log.Printf("Bucket already created: %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		}
		buckets = append(buckets, bucketName)
	}

	// Create files
	for _, name := range buckets {
		files := getEnvInt("CREATED_10KB_FILES_PER_BUCKET", 0)
		log.Printf("Creating %d 10kB files in bucket %s...\n", files, name)
		for i := files; i > 0; i-- {
			createFile(client, name, 10000)
		}

		files = getEnvInt("CREATED_500KB_FILES_PER_BUCKET", 0)
		log.Printf("Creating %d 500kB files in bucket %s...\n", files, name)
		for i := files; i > 0; i-- {
			createFile(client, name, 500000)
		}

		files = getEnvInt("CREATED_2MB_FILES_PER_BUCKET", 0)
		log.Printf("Creating %d 2MB files in bucket %s...\n", files, name)
		for i := files; i > 0; i-- {
			createFile(client, name, 2000000)
		}
	}
	log.Println("Done")
}

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if envVal, ok := os.LookupEnv(key); ok {
		envInt, err := strconv.ParseInt(envVal, 10, 0)
		if err == nil {
			return int(envInt)
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if envVal, ok := os.LookupEnv(key); ok {
		envBool, err := strconv.ParseBool(envVal)
		if err == nil {
			return envBool
		}
	}
	return defaultVal
}

func createFile(client *minio.Client, bucketName string, filesize int) {
	file := NewFakeFile(filesize)
	// log.Printf("Create file: %s (size=%d)\n", file.Name, file.Size)
	_, err := client.PutObject(context.Background(), bucketName, file.Name, file, int64(file.Size), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatalln(err)
	}
}

type FakeFile struct {
	Name      string
	Size      int
	readIndex int
}

func NewFakeFile(size int) *FakeFile {
	return &FakeFile{strconv.Itoa(rand.Int()), size, 0}
}

func (f *FakeFile) Read(b []byte) (n int, err error) {
	if f.readIndex >= f.Size {
		return 0, io.EOF
	}

	i := 0
	for (f.readIndex+i) < f.Size && i < len(b) {
		b[i] = byte('0' + (i % 10))
		i++
	}
	f.readIndex += i

	if f.readIndex >= f.Size {
		return i, io.EOF
	}
	return i, nil
}
