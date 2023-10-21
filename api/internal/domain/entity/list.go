package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type List struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"          validate:"-"`
	Name      string             `json:"name"         bson:"name"         validate:"required"`
	BoardID   string             `json:"boardId"      bson:"board_id"      validate:"-"`
	Position  int                `json:"position"     bson:"position"     validate:"-"`
	CardsIDs  []string           `json:"cardsIds"     bson:"cards_ids"     validate:"-"`
	CreatedAt time.Time          `json:"createdAt"    bson:"created_at"    validate:"-"`
	UpdatedAt time.Time          `json:"updatedAt"    bson:"updated_at"    validate:"-"`
}

func NewList(name, boardID string, position int) (*List, error) {
	list := List{
		ID:        primitive.NewObjectID(),
		Name:      name,
		BoardID:   boardID,
		Position:  position,
		CardsIDs:  []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := list.isValid()
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (list *List) Modify(name string, position int) error {
	list.Name = name
	list.Position = position
	list.UpdatedAt = time.Now()

	err := list.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (list *List) isValid() error {
	validate := validator.New()

	err := validate.Struct(list)

	if err != nil {
		return err
	}

	return nil
}
