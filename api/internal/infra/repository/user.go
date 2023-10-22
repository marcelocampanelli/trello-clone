package repository

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func NewTestUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		Client:     client,
		Collection: client.Database("trello-clone-test").Collection("users"),
	}
}

func (repository *UserRepository) Create(user *entity.User) (*string, error) {
	result, err := repository.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()

	return &userID, nil
}

func (repository *UserRepository) Update(id string, user *entity.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": user}

	_, err = repository.Collection.UpdateOne(context.Background(), filter, update)
	return nil
}

func (repository *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	filter := bson.M{"email": email}

	err := repository.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByCPF(cpf string) (*entity.User, error) {
	var user entity.User

	filter := bson.M{"cpf": cpf}

	err := repository.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	err = repository.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
