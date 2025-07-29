package stripe

import (
	"context"
	"log"
	"os"
)

type Config struct {
	APIKey  string
	BaseURL string
	Timeout string
}

type Client struct {
	config *Config
}

func NewClient() *Client {
	cfg := &Config{
		APIKey:  getEnv("STRIPE_API_KEY", "sk_test_demo"),
		BaseURL: getEnv("STRIPE_BASE_URL", "https://api.stripe.com"),
		Timeout: getEnv("STRIPE_TIMEOUT", "30s"),
	}

	return &Client{
		config: cfg,
	}
}

func (c *Client) ProcessPayment(ctx context.Context, amount int, currency string) error {
	log.Printf("[Stripe] Processing payment: %d %s", amount, currency)
	log.Printf("[Stripe] Using API: %s", c.config.BaseURL)
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
