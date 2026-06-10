package core_redis_pool

import (
	"context"
	"time"
)

type Pool interface {
	HGet(ctx context.Context, key string, field string) StringCmd
	HSet(ctx context.Context, key string, values ...interface{}) IntCmd
	Get(ctx context.Context, key string) StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusCmd
	Del(ctx context.Context, keys ...string) IntCmd
	Close() error
	GetTTL() time.Duration
}

type StringCmd interface {
	Bytes() ([]byte, error)
}

type IntCmd interface {
	Err() error
}

type StatusCmd interface {
	Err() error
}
