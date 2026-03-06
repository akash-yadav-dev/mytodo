package redis

import (
	"context"
	"errors"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

var ErrCacheMiss = errors.New("cache miss")

type Repository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttlSeconds int) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type RedisRepository struct {
	client *Client
}

func NewRepository(client *Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key)
	if err == goredis.Nil {
		return "", ErrCacheMiss
	}
	return value, err
}

func (r *RedisRepository) Set(ctx context.Context, key string, value any, ttlSeconds int) error {
	return r.client.Set(ctx, key, value, timeSeconds(ttlSeconds))
}

func (r *RedisRepository) Delete(ctx context.Context, key string) error {
	return r.client.Delete(ctx, key)
}

func (r *RedisRepository) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func timeSeconds(ttlSeconds int) time.Duration {
	if ttlSeconds <= 0 {
		return 0
	}
	return time.Duration(ttlSeconds) * time.Second
}
