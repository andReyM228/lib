package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type rabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
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
