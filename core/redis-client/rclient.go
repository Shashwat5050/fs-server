package rclient

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	coreErr "iceline-hosting.com/core/error"
)

const (
	leftDirection = "left"
)

type rClient struct {
	db *redis.Client
}

func NewRedisClient(redisURL, password string) (rClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: password,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return rClient{}, err
	}

	return rClient{
		db: rdb,
	}, nil
}

func (r rClient) Set(ctx context.Context, key string, value interface{}, expTime time.Duration) error {
	err := r.db.Set(ctx, key, value, expTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r rClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.db.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r rClient) Delete(ctx context.Context, key string) error {
	err := r.db.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r rClient) Push(ctx context.Context, queueName string, value ...interface{}) error {
	err := r.db.LPush(ctx, queueName, value...).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r rClient) Pop(ctx context.Context, queueName string, batchSize int) ([]string, error) {
	_, vals, err := r.db.BLMPop(ctx, time.Second, leftDirection, int64(batchSize), queueName).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("%w - %s", coreErr.ErrEmptyQueue, queueName)
	} else if err != nil {
		return nil, err
	}

	return vals, nil
}
