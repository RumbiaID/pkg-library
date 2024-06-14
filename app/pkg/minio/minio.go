package minio

import (
	"context"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"pkg-library/app/config"
)

func InitiateMinio(config *config.Config) (*minio.Client, error) {
	ctx := context.Background()
	// Initialize minio client object.
	minioClient, err := minio.New(config.Storage.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Storage.MinioAccessKey, config.Storage.MinioSecretKey, ""),
		Secure: true,
	})
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	err = minioClient.MakeBucket(ctx, config.Storage.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, config.Storage.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", config.Storage.BucketName)
		} else {
			//log.Fatalln(err)
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", config.Storage.BucketName)
	}
	return minioClient, nil
}
