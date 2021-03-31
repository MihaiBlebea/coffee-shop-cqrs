package handler

import (
	"net/http"

	"github.com/MihaiBlebea/coffee-shop-command/user"
)

type Logger interface {
	Info(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type UserService interface {
	NewUser(firstName, lastName string, age uint) (*user.User, error)
	OrderCoffee(userID, coffeeID string) (string, error)
}

type Service struct {
	userService UserService
	logger      Logger
}

func New(userService UserService, logger Logger) *Service {
	return &Service{userService, logger}
}

func (s *Service) AuthenticateEndpoint() http.Handler {
	return authenticateEndpoint(s.userService, s.logger)
}

func (s *Service) OrderEndpoint() http.Handler {
	return orderEndpoint(s.userService, s.logger)
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint(s.logger)
}
