package broker

import (
	"github.com/streadway/amqp"
)

type Service struct {
}

func New() (*Service, error) {
	s := Service{}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:15673/")
	if err != nil {
		return &s, err
	}
	defer conn.Close()

	return &s, nil
}
