package storage

import (
	"context"
	"fmt"
	minio2 "github.com/EmirShimshir/marketplace/pkg/minio"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/testcontainers/testcontainers-go"
	testminio "github.com/testcontainers/testcontainers-go/modules/minio"
)

var (
	minioConfig = minio2.Config{
		User:       "key",
		Password:   "secretKey",
		BucketName: "test",
	}
)

func newMinioContainer(ctx context.Context) (*testminio.MinioContainer, error) {
	minioContainer, err := testminio.RunContainer(ctx,
		testcontainers.WithImage("docker.io/minio/minio"),
		testminio.WithUsername(minioConfig.User),
		testminio.WithPassword(minioConfig.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start minio container: %s", err)
	}
	return minioContainer, nil
}

func newMinioClient(url string) (*minio.Client, error) {
	minioClient, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.User, minioConfig.Password, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %s", err)
	}

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, minioConfig.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, minioConfig.BucketName)
		if errBucketExists != nil || !exists {
			return nil, fmt.Errorf("failed to make minio bucket: %s", errBucketExists)
		}
	}
	policy := fmt.Sprintf(`{
		"Version":"2012-10-17",
		"Statement":[{
			"Effect":"Allow",
			"Principal":"*",
			"Action":["s3:GetObject"],
			"Resource":["arn:aws:s3:::%s/*"]}
		]}`, minioConfig.BucketName)
	err = minioClient.SetBucketPolicy(ctx, minioConfig.BucketName, policy)
	if err != nil {
		return nil, fmt.Errorf("failed to set bucket public policy: %s", err)
	}

	return minioClient, nil
}
