package cache

import (
	"context"
	"time"
)

type MockCacheClient struct {
	SetFunc  func(context.Context, string, interface{}, time.Duration) error
	GetFunc  func(context.Context, string) (string, error)
	HGetFunc func(context.Context, string, string) (string, error)
	HSetFunc func(context.Context, string, ...interface{}) error
}

func (m *MockCacheClient) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return m.SetFunc(ctx, key, value, duration)
}

func (m *MockCacheClient) Get(ctx context.Context, key string) (string, error) {
	return m.GetFunc(ctx, key)
}

func (m *MockCacheClient) HGet(ctx context.Context, key string, field string) (string, error) {
	return m.HGetFunc(ctx, key, field)
}

func (m *MockCacheClient) HSet(ctx context.Context, key string, values ...interface{}) error {
	return m.HSetFunc(ctx, key, values)
}
