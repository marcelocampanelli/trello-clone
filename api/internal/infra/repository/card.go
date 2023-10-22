package repository

import (
	"context"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewCardRepository(client *mongo.Client) *CardRepository {
	return &CardRepository{
		Client:     client,
		Collection: client.Database("trello-clone").Collection("cards"),
	}
}

func NewTestCardRepository(client *mongo.Client) *CardRepository {
	return &CardRepository{
		Client:     client,
		Collection: client.Database("trello-clone-test").Collection("cards"),
	}
}

func (repository *CardRepository) Create(card *entity.Card) (*entity.Card, error) {
	result, err := repository.Collection.InsertOne(context.Background(), card)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": result.InsertedID}

	var cardInserted entity.Card

	err = repository.Collection.FindOne(context.Background(), filter).Decode(&cardInserted)
	if err != nil {
		return nil, err
	}

	return &cardInserted, nil
}

func (repository *CardRepository) Update(id string, card *entity.Card) (*entity.Card, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": card}

	_, err = repository.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (repository *CardRepository) Delete(id string) error {
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

func (repository *CardRepository) FindByID(id string) (*entity.Card, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var card entity.Card

	err = repository.Collection.FindOne(context.Background(), filter).Decode(&card)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (repository *CardRepository) FindAllByList(listID string) ([]*entity.Card, error) {
	var cards []*entity.Card

	filter := bson.M{"list_id": listID}

	cursor, err := repository.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &cards); err != nil {
		return nil, err
	}

	return cards, nil
}
