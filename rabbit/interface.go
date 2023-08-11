package rabbit

type Rabbit interface {
	Publish(queueName string, message interface{}) error
	Consume(queueName string, handler func([]byte) error) error
	CloseConnection() error
	Reply(queueName string, statusCode int64, message interface{}) error
	Request(queueName string, message interface{}) (ResponseModel, error)
}
