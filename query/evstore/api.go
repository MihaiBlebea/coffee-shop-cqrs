package evstore

import (
	"fmt"

	goes "github.com/jetbasrawi/go.geteventstore"
)

type OrderViewStore interface {
	AddUser(firstName, lastName, token string, age int) error
	AddOrder(userToken, coffeeName string, price uint) error
}

type Service struct {
	client         *goes.Client
	orderViewStore OrderViewStore
}

func New(orderViewStore OrderViewStore) (*Service, error) {
	s := Service{}

	client, err := goes.NewClient(nil, "http://localhost:2113")
	if err != nil {
		return &s, err
	}

	s.client = client
	s.orderViewStore = orderViewStore

	return &s, nil
}

func (s *Service) Listen() {
	reader := s.client.NewStreamReader("FooStream")

	for reader.Next() {
		if reader.Err() != nil {
			if _, ok := reader.Err().(*goes.ErrNoMoreEvents); ok {
				reader.LongPoll(15)
			}
		} else {
			fooMeta := make(map[string]string)
			ev := UserCreated{}
			_ = reader.Scan(&ev, &fooMeta)

			if err := s.orderViewStore.AddUser(ev.FirstName, ev.LastName, ev.ID, int(ev.Age)); err != nil {
				fmt.Println(err)
			}

			fmt.Println(ev)
		}
	}
}
