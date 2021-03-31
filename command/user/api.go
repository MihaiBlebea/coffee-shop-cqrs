package user

import (
	"github.com/MihaiBlebea/coffee-shop-command/coffee"
	"github.com/MihaiBlebea/coffee-shop-command/trans"
	"gorm.io/gorm"
)

type CoffeeService interface {
	GetCoffeeByID(coffeeID string) (*coffee.Coffee, error)
}

type TransactionService interface {
	NewTransaction(userID, coffeeID string) (*trans.Transaction, error)
}

type EventStore interface {
	Publish(eventType string, data interface{}, meta interface{}) error
}

type Service struct {
	repo               *store
	coffeeService      CoffeeService
	transactionService TransactionService
	eventStore         EventStore
}

func New(db *gorm.DB, coffeeService CoffeeService, transactionService TransactionService, eventStore EventStore) *Service {
	s := newStore(db)
	return &Service{s, coffeeService, transactionService, eventStore}
}

func (s *Service) Migrate() error {
	err := s.repo.createTable()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) NewUser(firstName, lastName string, age uint) (*User, error) {
	u, err := new(firstName, lastName, age)
	if err != nil {
		return u, err
	}

	s.repo.save(u)

	meta := make(map[string]string)
	meta["Name"] = "user.created"

	s.eventStore.Publish("user.created", u, meta)

	return u, nil
}

func (s *Service) OrderCoffee(userID, coffeeID string) (string, error) {
	return orderCoffee(s.repo, s.coffeeService, s.transactionService, userID, coffeeID)
}
