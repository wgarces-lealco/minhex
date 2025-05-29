package messaging

import "context"

type RabbitMQPublisher struct{}

func NewRabbitMQPublisher() *RabbitMQPublisher {
	return &RabbitMQPublisher{}
}

func (p *RabbitMQPublisher) Publish(ctx context.Context, topic string, message []byte) error {
	return nil
}
