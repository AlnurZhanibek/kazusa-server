package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"os"
)

type FileService struct {
	client *s3.S3
}

func NewFileService() *FileService {
	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String(os.Getenv("SPACES_ENDPOINT")),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("us-east-1"),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	return &FileService{client: s3Client}
}

func (fs *FileService) Put(filename string, file multipart.File) (*string, error) {
	object := s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("SPACES_BUCKET")),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"),
		Metadata: map[string]*string{
			"x-amz-meta-my-key": aws.String(filename),
		},
	}

	_, err := fs.client.PutObject(&object)
	if err != nil {
		return nil, err
	}

	url := os.Getenv("SPACES_CDN_ENDPOINT") + "/" + filename

	return &url, nil
}
