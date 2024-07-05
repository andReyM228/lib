package rabbit

import (
	"context"
)

type Rabbit interface {
	Publish(queueName string, message interface{}) error
	Consume(ctx context.Context, queueName string, handler func([]byte) error) error
	ConsumeAll(ctx context.Context, queues map[string]func([]byte) error) error
	CloseConnection() error
	Reply(queueName string, statusCode int64, message interface{}) error
	Request(queueName string, message interface{}) (ResponseModel, error)
	PreparePublish(queueName string, message interface{}) error
	Close() error
}
