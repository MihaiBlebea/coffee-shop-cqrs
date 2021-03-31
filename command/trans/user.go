package trans

import (
	"time"

	"github.com/gofrs/uuid"
)

type Transaction struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	CoffeeID  string    `json:"coffee_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func new(userID, coffeeID string) (*Transaction, error) {
	transaction := Transaction{}

	token, err := genToken()
	if err != nil {
		return &transaction, err
	}

	transaction.ID = token
	transaction.UserID = userID
	transaction.CoffeeID = coffeeID

	return &transaction, nil
}

func genToken() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
