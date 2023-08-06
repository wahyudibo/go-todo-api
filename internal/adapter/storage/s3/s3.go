package s3storageadapter

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	storageadapter "github.com/wahyudibo/go-todo-api/internal/adapter/storage"
	"github.com/wahyudibo/go-todo-api/internal/config"
)

var _ storageadapter.StorageAdapter = (*s3StorageAdapter)(nil)

type s3StorageAdapter struct {
	config   *config.Config
	s3Client *s3.Client
}

func New(ctx context.Context, cfg *config.Config) (storageadapter.StorageAdapter, error) {
	sdkConfig, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithSharedConfigProfile(cfg.AWSProfile),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	return &s3StorageAdapter{
		config:   cfg,
		s3Client: s3Client,
	}, nil
}

func (sa *s3StorageAdapter) Upload(ctx context.Context, directory, objectKey string, body io.Reader) (string, error) {
	fullPath := path.Join(directory, objectKey)

	_, err := sa.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(sa.config.S3BucketName),
		Key:    aws.String(fullPath),
		Body:   body,
	})
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

func (sa *s3StorageAdapter) GenerateDownloadURL(ctx context.Context, objectKey string) (string, error) {
	return "", errors.New("unimplemented")
}
