package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       uint      `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func new(firstName, lastName string, age uint) (*User, error) {
	user := User{}

	token, err := genToken()
	if err != nil {
		return &user, err
	}

	user.ID = token
	user.FirstName = firstName
	user.LastName = lastName
	user.Age = age

	return &user, nil
}

func genToken() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
