package repository

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewListRepository(client *mongo.Client) *ListRepository {
	return &ListRepository{
		Client:     client,
		Collection: client.Database("trello-clone").Collection("lists"),
	}
}

func NewTestListRepository(client *mongo.Client) *ListRepository {
	return &ListRepository{
		Client:     client,
		Collection: client.Database("trello-clone-test").Collection("lists"),
	}
}

func (repository *ListRepository) Create(list *entity.List) (*entity.List, error) {
	result, err := repository.Collection.InsertOne(context.Background(), list)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": result.InsertedID}

	var listInserted entity.List

	repository.Collection.FindOne(context.Background(), filter).Decode(&listInserted)

	return &listInserted, nil
}

func (repository *ListRepository) Update(id string, list *entity.List) (*entity.List, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": list}

	_, err = repository.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (repository *ListRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = repository.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ListRepository) FindAll(boardID string) ([]*entity.List, error) {
	var lists []*entity.List

	filter := bson.M{"board_id": boardID}

	cursor, err := repository.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &lists); err != nil {
		return nil, err
	}

	return lists, nil
}

func (repository *ListRepository) FindByID(id string) (*entity.List, error) {
	var list entity.List

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	err = repository.Collection.FindOne(context.Background(), filter).Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
