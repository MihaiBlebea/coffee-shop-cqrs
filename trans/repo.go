package trans

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func newStore(db *gorm.DB) *store {
	return &store{db}
}

func (s *store) createTable() error {
	result := s.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS transactions (
			id VARCHAR(255) PRIMARY KEY,
			user_id VARCHAR(255) NOT NULL,
			coffee_id VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
		`)

	return result.Error
}

func (s *store) save(transaction *Transaction) {
	s.db.Create(transaction)
}
