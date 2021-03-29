package broker

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Service struct {
	conn *amqp.Connection
}

func New() (*Service, error) {
	s := Service{}
	conn, err := amqp.Dial("amqp://admin:pass@localhost:5673")
	if err != nil {
		return &s, err
	}

	s.conn = conn

	return &s, nil
}

func (s *Service) CreateChannel() error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"test.queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) PublishMessage(payload interface{}) error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		"test.queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Listen() error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"test.queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for m := range msgs {
		fmt.Println(string(m.Body))
	}

	return nil
}
