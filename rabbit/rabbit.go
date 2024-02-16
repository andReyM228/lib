package rabbit

import (
	"encoding/json"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"sync"
)

type rabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	log  log.Logger
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

func NewRabbitMQ(url string, log log.Logger) (Rabbit, error) {
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
		log:  log,
	}, err
}

func (r rabbitMQ) Request(queueName string, message interface{}) (ResponseModel, error) {
	r.log.Debugf("start request to %s", queueName)

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

	r.log.Debugf("request send")

	result, err := r.ConsumeWithResponse(replyTopic)
	if err != nil {
		return ResponseModel{}, err
	}

	if result == nil {
		r.log.Debugf("result is nil")
		return ResponseModel{}, errs.InternalError{}
	}

	var resp ResponseModel

	r.log.Debugf("start unmarshalling payload")

	err = json.Unmarshal(result, &resp)
	if err != nil {
		return ResponseModel{}, err
	}

	r.log.Debugf("return response")

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
	go func() {
		queue, err := r.ch.QueueDeclare(
			queueName,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {

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
			r.log.Fatalf("Failed to register a consumer: %v", err)
		}

		go func() {
			for msg := range msgs {
				if err := handler(msg.Body); err != nil {
					r.log.Errorf("Failed to process message: %v", err)
				}
			}
		}()
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

	r.log.Debugf("start listening replyTopic")

	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		r.log.Debugf("start listening msgs")

		for msg := range msgs {
			result = msg.Body
			r.log.Debugf("got result from msgs")

			wg.Done()
		}
	}(&wg)
	//***
	wg.Wait()

	return result, nil
}

func (r rabbitMQ) PreparePublish(queueName string, message interface{}) error {
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
			Headers: amqp.Table{
				"action": "prepare",
			},
			Body: body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
