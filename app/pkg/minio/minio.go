package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Config struct {
	MinioEndpoint  string `validate:"required" name:"MINIO_ENDPOINT"`
	MinioAccessKey string `validate:"required" name:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `validate:"required" name:"MINIO_SECRET_KEY"`
	BucketName     string `validate:"required" name:"BUCKET_NAME"`
}

func InitiateMinio(config *Config) (*minio.Client, error) {
	ctx := context.Background()
	// Initialize minio client object.
	minioClient, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: true,
	})
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	err = minioClient.MakeBucket(ctx, config.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, config.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", config.BucketName)
		} else {
			//log.Fatalln(err)
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", config.BucketName)
	}
	return minioClient, nil
}
