package repository

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewBoardRepository(client *mongo.Client) *BoardRepository {
	return &BoardRepository{
		Client:     client,
		Collection: client.Database("trello-clone").Collection("boards"),
	}
}

func (repository *BoardRepository) Create(board *entity.Board) (*string, error) {
	result, err := repository.Collection.InsertOne(context.Background(), board)
	if err != nil {
		return nil, err
	}

	boardID := result.InsertedID.(primitive.ObjectID).Hex()

	return &boardID, nil
}

func (repository *BoardRepository) Update(id string, board *entity.Board) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := primitive.M{"_id": objectID}
	update := primitive.M{"$set": board}

	_, err = repository.Collection.UpdateOne(context.Background(), filter, update)

	return nil
}

func (repository *BoardRepository) FindAll(userID string) ([]*entity.Board, error) {
	var boards []*entity.Board

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := primitive.M{"user_id": userObjectID}

	results, err := repository.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for results.Next(context.Background()) {
		var board entity.Board
		err := results.Decode(&board)
		if err != nil {
			return nil, err
		}
		boards = append(boards, &board)
	}

	return boards, nil
}

func (repository *BoardRepository) FindByID(id string) (*entity.Board, error) {
	var board entity.Board

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := primitive.M{"_id": objectID}

	err = repository.Collection.FindOne(context.Background(), filter).Decode(&board)
	if err != nil {
		return nil, err
	}

	return &board, nil
}

func (repository *BoardRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := primitive.M{"_id": objectID}

	_, err = repository.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
