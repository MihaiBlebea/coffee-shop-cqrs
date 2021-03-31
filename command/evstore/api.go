package evstore

import (
	goes "github.com/jetbasrawi/go.geteventstore"
)

type Service struct {
	client *goes.Client
}

func New() (*Service, error) {
	s := Service{}

	client, err := goes.NewClient(nil, "http://localhost:2113")
	if err != nil {
		return &s, err
	}

	s.client = client

	return &s, nil
}

func (s *Service) Publish(eventType string, data interface{}, meta interface{}) error {
	ev := goes.NewEvent(goes.NewUUID(), eventType, &data, &meta)

	writer := s.client.NewStreamWriter("FooStream")

	err := writer.Append(nil, ev)
	if err != nil {
		return err
	}

	return nil
}
