package handler

import (
	"net/http"
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

type Service struct {
	logger Logger
}

func New(logger Logger) *Service {
	return &Service{logger}
}

func (s *Service) OrdersEndpoint() http.Handler {
	return ordersEndpoint(s.logger)
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint(s.logger)
}
