package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Card struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"          validate:"-"`
	Name      string             `json:"name"         bson:"name"         validate:"required"`
	CreatedAt time.Time          `json:"createdAt"    bson:"createdAt"    validate:"-"`
	UpdatedAt time.Time          `json:"updatedAt"    bson:"updatedAt"    validate:"-"`
}

func NewCard(name string) (*Card, error) {
	card := Card{
		ID:        primitive.NewObjectID(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := card.isValid()
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (card *Card) Modify(name string, user User) error {
	card.Name = name
	card.UpdatedAt = time.Now()

	err := card.isValid()

	if err != nil {
		return err
	}

	return nil
}

func (card *Card) isValid() error {
	validate := validator.New()

	err := validate.Struct(card)

	if err != nil {
		return err
	}

	return nil
}
