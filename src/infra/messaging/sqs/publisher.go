package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"minhex/src/domain/shared/ports"
	"os"
)

type Config struct {
	QueueURL string
}

type Publisher struct {
	config *Config
}

func NewPublisher() ports.EventPublisher {
	cfg := &Config{
		QueueURL: getEnv("SQS_QUEUE_URL", "https://sqs.us-east-1.amazonaws.com/123456789/demo-queue"),
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

	log.Printf("[SQS] Publishing event to topic '%s': %s", topic, string(eventData))
	log.Printf("[SQS] Queue URL: %s", p.config.QueueURL)

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
