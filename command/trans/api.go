package trans

import (
	"gorm.io/gorm"
)

type Service struct {
	repo *store
}

func New(db *gorm.DB) *Service {
	s := newStore(db)
	return &Service{repo: s}
}

func (s *Service) Migrate() error {
	err := s.repo.createTable()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) NewTransaction(userID, coffeeID string) (*Transaction, error) {
	t, err := new(userID, coffeeID)
	if err != nil {
		return t, err
	}

	s.repo.save(t)

	return t, nil
}
