package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type Config struct {
	Endpoint   string
	User       string
	Password   string
	BucketName string
	Host       string
}

func NewClient(cfg *Config) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.User, cfg.Password, ""),
		Secure: false,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create minio client")
	}

	ctx := context.Background()
	err = minioClient.MakeBucket(ctx, cfg.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, cfg.BucketName)
		if errBucketExists != nil || !exists {
			return nil, errors.Wrap(errBucketExists, "failed to make minio bucket")
		}
	}
	policy := fmt.Sprintf(`{
		"Version":"2012-10-17",
		"Statement":[{
			"Effect":"Allow",
			"Principal":"*",
			"Action":["s3:GetObject"],
			"Resource":["arn:aws:s3:::%s/*"]}
		]}`, cfg.BucketName)
	err = minioClient.SetBucketPolicy(ctx, cfg.BucketName, policy)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set bucket public policy")
	}

	return minioClient, nil
}
