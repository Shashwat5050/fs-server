package backblaze

import (
	"context"
	"errors"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	// ErrMissingEnv is returned when one of the required environment variables is not set.
	ErrMissingEnv = errors.New("missing environment variable")
)

type client struct {
	session *session.Session
}

/*
NewClient creates a new client for the Backblaze B2 API.
It requires the following environment variables to be set:
  - B2_KEY_ID
  - B2_APP_KEY
  - B2_ENDPOINT
  - B2_BUCKET
*/
func NewClient() (client, error) {
	keyID := os.Getenv("B2_KEY_ID")
	appKey := os.Getenv("B2_APP_KEY")
	endpoint := os.Getenv("B2_ENDPOINT")
	bucket := os.Getenv("B2_BUCKET")
	log.Println(keyID, appKey, endpoint, bucket, "*****************************")
	if keyID == "" || appKey == "" || endpoint == "" || bucket == "" {
		return client{}, ErrMissingEnv
	}

	s3Conf := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(keyID, appKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
	}

	s3session := session.New(s3Conf)

	return client{session: s3session}, nil
}

func (c client) WriteObject(ctx context.Context, bucket, key string, body io.Reader) error {
	_, err := s3manager.NewUploader(c.session).UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c client) GetObject(ctx context.Context, bucket, key string) (io.ReadCloser, int64, error) {
	svc := s3.New(c.session)

	out, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, 0, err
	}

	return out.Body, aws.Int64Value(out.ContentLength), nil
}

func (c client) GetObjectSize(ctx context.Context, bucket, key string) (int64, error) {
	svc := s3.New(c.session)

	out, err := svc.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return 0, err
	}

	return aws.Int64Value(out.ContentLength), nil
}

func (c client) DeleteObject(ctx context.Context, bucket, key string) error {
	svc := s3.New(c.session)

	_, err := svc.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}
