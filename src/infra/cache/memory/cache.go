package memory

import (
	"context"
	"sync"
	"time"
)

type Config struct {
	DefaultTTL time.Duration
}

type Cache struct {
	config *Config
	data   map[string]cacheItem
	mu     sync.RWMutex
}

type cacheItem struct {
	value  interface{}
	expiry time.Time
}

func NewCache() *Cache {
	cfg := &Config{
		DefaultTTL: 5 * time.Minute,
	}

	return &Cache{
		config: cfg,
		data:   make(map[string]cacheItem),
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheItem{
		value:  value,
		expiry: time.Now().Add(c.config.DefaultTTL),
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[key]
	if !exists || time.Now().After(item.expiry) {
		return nil, nil
	}
	return item.value, nil
}
