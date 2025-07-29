package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"minhex/src/domain/shared/ports"
	"os"
)

type Config struct {
	ConnectionURL string
}

type Publisher struct {
	config *Config
}

func NewPublisher() ports.EventPublisher {
	cfg := &Config{
		ConnectionURL: getEnv("RABBITMQ_URL", "amqp://localhost:5672"),
	}

	return &Publisher{
		config: cfg,
	}
}

func (p *Publisher) Publish(ctx context.Context, topic string, event interface{}) error {
	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	log.Printf("[RabbitMQ] Publishing event to topic '%s': %s", topic, string(eventData))
	log.Printf("[RabbitMQ] Connection URL: %s", p.config.ConnectionURL)

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
