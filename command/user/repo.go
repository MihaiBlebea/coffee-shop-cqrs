package user

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
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			age INT(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
		`)

	return result.Error
}

func (s *store) save(user *User) {
	s.db.Create(user)
}

func (s *store) GetByID(ID string) (*User, error) {
	user := User{}
	err := s.db.Where("id = ?", ID).Find(&user).Error

	return &user, err
}
