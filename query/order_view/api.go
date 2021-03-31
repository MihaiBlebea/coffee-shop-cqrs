package orderv

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	client *mongo.Client
}

func New(client *mongo.Client) *Service {
	return &Service{client}
}

func (s *Service) AddUser(firstName, lastName, token string, age int) error {
	collection := s.client.Database("coffee_shop").Collection("order_view")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	ctx := context.Background()

	_, err := collection.InsertOne(ctx, bson.D{
		{"firstName", firstName},
		{"lastName", lastName},
		{"token", token},
		{"age", age},
	})
	if err != nil {
		return err
	}
	// id, ok := res.InsertedID.(string)
	return nil
}

func (s *Service) AddOrder(userToken, coffeeName string, price uint) error {
	collection := s.client.Database("coffee_shop").Collection("order_view")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ctx := context.Background()

	_, err := collection.UpdateOne(ctx, bson.D{{"token", userToken}}, bson.D{
		{"$push", bson.M{
			"orders": bson.M{
				"coffeeName": coffeeName,
				"price":      price,
			},
		}},
	})
	if err != nil {
		return err
	}
	// id, ok := res.InsertedID.(string)
	return nil
}
