package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (r *RabbitMQ) NewRabbitMQ(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}

	ch.Close()
	conn.Close()
	return nil
}

func (r *RabbitMQ) Publish(queueName string, message interface{}) error {
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

func (r *RabbitMQ) Consume(queueName string, handler func([]byte) error) error {
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
