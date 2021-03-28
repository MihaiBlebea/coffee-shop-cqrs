package coffee

import (
	"encoding/json"
	"errors"

	_ "embed"
)

//go:embed products.json
var jsonCoffee []byte

type Service struct {
	products []Coffee
}

func New() (*Service, error) {
	var products []Coffee
	service := Service{}

	err := json.Unmarshal(jsonCoffee, &products)
	if err != nil {
		return &service, err
	}

	service.products = products

	return &service, nil
}

func (s *Service) GetCoffeeByID(id string) (*Coffee, error) {
	for _, coffee := range s.products {
		if coffee.ID == id {
			return &coffee, nil
		}
	}

	return &Coffee{}, errors.New("Could not find coffee by id")
}
