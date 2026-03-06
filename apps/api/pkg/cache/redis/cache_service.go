package redis

import (
	"context"
	"strings"
	"time"
)

type CacheOptions struct {
	DefaultTTL time.Duration
	KeyPrefix  string
}

type CacheService struct {
	repo       Repository
	defaultTTL time.Duration
	keyPrefix  string
}

func NewCacheService(repo Repository, opts CacheOptions) *CacheService {
	ttl := opts.DefaultTTL
	if ttl == 0 {
		ttl = 5 * time.Minute
	}

	return &CacheService{
		repo:       repo,
		defaultTTL: ttl,
		keyPrefix:  strings.Trim(opts.KeyPrefix, ":"),
	}
}

func (s *CacheService) Get(ctx context.Context, key string) (string, error) {
	return s.repo.Get(ctx, s.buildKey(key))
}

func (s *CacheService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if ttl == 0 {
		ttl = s.defaultTTL
	}
	return s.repo.Set(ctx, s.buildKey(key), value, int(ttl.Seconds()))
}

func (s *CacheService) Delete(ctx context.Context, key string) error {
	return s.repo.Delete(ctx, s.buildKey(key))
}

func (s *CacheService) Exists(ctx context.Context, key string) (bool, error) {
	return s.repo.Exists(ctx, s.buildKey(key))
}

func (s *CacheService) buildKey(key string) string {
	if s.keyPrefix == "" {
		return key
	}
	return s.keyPrefix + ":" + key
}
