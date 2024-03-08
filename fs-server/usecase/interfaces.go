package usecase

import (
	"context"
	"io"
)

type objectStorage interface {
	WriteObject(ctx context.Context, bucket, key string, body io.Reader) error
	GetObject(ctx context.Context, bucket, key string) (io.ReadCloser, int64, error)
	GetObjectSize(ctx context.Context, bucket, key string) (int64, error)
	DeleteObject(ctx context.Context, bucket, key string) error
}
