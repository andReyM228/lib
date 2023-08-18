package rabbit

import (
	"encoding/json"
	"github.com/andReyM228/lib/errs"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type rabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type RequestModel struct {
	ReplyTopic string
	Payload    json.RawMessage
}

type ResponseModel struct {
	StatusCode int64
	Payload    json.RawMessage
}

func (r rabbitMQ) CloseConnection() error {
	err := r.conn.Close()
	if err != nil {
		return err
	}

	err = r.ch.Close()
	if err != nil {
		return err
	}

	return nil
}

func NewRabbitMQ(url string) (Rabbit, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return rabbitMQ{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return rabbitMQ{}, err
	}

	return rabbitMQ{
		conn: conn,
		ch:   ch,
	}, err
}

func (r rabbitMQ) Request(queueName string, message interface{}) (ResponseModel, error) {
	replyTopic := uuid.New().String()

	payload, err := json.Marshal(message)
	if err != nil {
		return ResponseModel{}, err
	}

	request := RequestModel{
		ReplyTopic: replyTopic,
		Payload:    payload,
	}

	err = r.Publish(queueName, request)
	if err != nil {
		return ResponseModel{}, err
	}

	result, err := r.ConsumeWithResponse(replyTopic)
	if err != nil {
		return ResponseModel{}, err
	}

	if result == nil {
		return ResponseModel{}, errs.InternalError{}
	}

	var resp ResponseModel

	err = json.Unmarshal(result, &resp)
	if err != nil {
		return ResponseModel{}, err
	}

	return resp, nil
}

func (r rabbitMQ) Reply(queueName string, statusCode int64, message interface{}) error {
	var payload json.RawMessage
	var err error

	if message != nil {
		payload, err = json.Marshal(message)
		if err != nil {
			return err
		}
	}

	request := ResponseModel{
		StatusCode: statusCode,
		Payload:    payload,
	}

	err = r.Publish(queueName, request)
	if err != nil {
		return err
	}

	return nil
}

func (r rabbitMQ) Publish(queueName string, message interface{}) error {
	queue, err := r.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = r.ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r rabbitMQ) Consume(queueName string, handler func([]byte) error) error {
	queue, err := r.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := r.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err != nil {
				log.Printf("Failed to process message: %v", err)
			}
		}
	}()

	return nil
}

func (r rabbitMQ) ConsumeWithResponse(queueName string) ([]byte, error) {
	queue, err := r.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := r.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var result []byte

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()

		for msg := range msgs {
			result = msg.Body
		}
	}()
	//***
	wg.Wait()

	return result, nil
}
