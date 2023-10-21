package repository

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		Client:     client,
		Collection: client.Database("trello-clone").Collection("users"),
	}
}

func (repository *UserRepository) Create(user *entity.User) error {
	_, err := repository.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
