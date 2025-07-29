package redis

import (
	"context"
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Password string
}

type Cache struct {
	config *Config
}

func NewCache() *Cache {
	cfg := &Config{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     getEnv("REDIS_PORT", "6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
	}

	return &Cache{
		config: cfg,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
	log.Printf("[Redis] Setting key '%s' on %s:%s", key, c.config.Host, c.config.Port)
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	log.Printf("[Redis] Getting key '%s' from %s:%s", key, c.config.Host, c.config.Port)
	return nil, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
