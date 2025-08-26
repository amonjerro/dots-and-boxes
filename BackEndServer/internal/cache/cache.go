package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient interface {
	Set(context.Context, string, interface{}, time.Duration) error
	HSet(context.Context, string, ...interface{}) error
	Get(context.Context, string) (string, error)
	HGet(context.Context, string, string) (string, error)
}

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := r.client.Set(ctx, key, value, duration).Err()
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	return value, err
}

func (r *RedisClient) HSet(ctx context.Context, key string, values ...interface{}) error {
	err := r.client.HSet(ctx, key, values).Err()
	return err
}

func (r *RedisClient) HGet(ctx context.Context, key string, field string) (string, error) {
	value, err := r.client.HGet(ctx, key, field).Result()
	return value, err
}
