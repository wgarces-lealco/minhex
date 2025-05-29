package cache

import "context"

type RedisCache struct{}

func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value string) error {
	return nil
}
