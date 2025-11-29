package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type FileService struct {
	client *s3.Client
	bucket string
}

func NewFileService(ctx context.Context) (*FileService, error) {
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("S3_BUCKET")

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return &FileService{
		client: client,
		bucket: bucket,
	}, nil
}

func (fs *FileService) Put(ctx context.Context, key string, r multipart.File) (*string, error) {
	_, err := fs.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(key),
		Body:   r,
	})

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
		os.Getenv("S3_BUCKET"),
		os.Getenv("AWS_REGION"),
		key,
	)

	return &fileURL, err
}
