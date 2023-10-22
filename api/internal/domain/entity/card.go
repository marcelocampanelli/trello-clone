package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Card struct {
	ID             primitive.ObjectID `json:"id"            bson:"_id"           validate:"-"`
	Name           string             `json:"name"          bson:"name"          validate:"required"`
	UserAssignedID string             `json:"user_assigned" bson:"user_assigned" validate:"required"`
	Position       int                `json:"position"      bson:"position"      validate:"required"`
	CreatedAt      time.Time          `json:"created_at"    bson:"created_at"    validate:"-"`
	UpdatedAt      time.Time          `json:"updated_at"    bson:"updated_at"    validate:"-"`
}

func NewCard(name, userAssignedID string, position int) (*Card, error) {
	card := Card{
		ID:             primitive.NewObjectID(),
		Name:           name,
		UserAssignedID: userAssignedID,
		Position:       position,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := card.isValid()
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (card *Card) Modify(name, userAssignedID string, position int) error {
	card.Name = name
	card.UpdatedAt = time.Now()
	card.UserAssignedID = userAssignedID
	card.Position = position
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
