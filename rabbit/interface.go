package rabbit

type Rabbit interface {
	Publish(queueName string, message interface{}) error
	Consume(queueName string, handler func([]byte) error) error
	CloseConnection() error
}
