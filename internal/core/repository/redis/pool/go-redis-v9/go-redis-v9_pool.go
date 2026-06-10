package core_goredisv9_pool

import (
	"context"
	"fmt"
	core_redis_pool "restapi/internal/core/repository/redis/pool"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redisv9Pool struct {
	client *redis.Client
	TTL    time.Duration
}

func NewRedisv9Pool(ctx context.Context, config Config) (*Redisv9Pool, error) {
	options := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	}
	client := redis.NewClient(&options)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}
	return &Redisv9Pool{
		client: client,
		TTL:    config.TTL,
	}, nil
}

func (p *Redisv9Pool) Close() error {
	return p.client.Close()
}

func (p *Redisv9Pool) GetTTL() time.Duration {
	return p.TTL
}

func (p *Redisv9Pool) HGet(ctx context.Context, key string, field string) core_redis_pool.StringCmd {
	stringCmd := p.client.HGet(ctx, key, field)
	return StringCmd{stringCmd}
}

func (p *Redisv9Pool) HSet(ctx context.Context, key string, values ...interface{}) core_redis_pool.IntCmd {
	intCmd := p.client.HSet(ctx, key, values...)
	return IntCmd{intCmd}
}

func (p *Redisv9Pool) Get(ctx context.Context, key string) core_redis_pool.StringCmd {
	stringCmd := p.client.Get(ctx, key)
	return StringCmd{stringCmd}
}

func (p *Redisv9Pool) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) core_redis_pool.StatusCmd {
	statusCmd := p.client.Set(ctx, key, value, expiration)
	return statusCmd
}

func (p *Redisv9Pool) Del(ctx context.Context, keys ...string) core_redis_pool.IntCmd {
	intCmd := p.client.Del(ctx, keys...)
	return IntCmd{intCmd}
}
