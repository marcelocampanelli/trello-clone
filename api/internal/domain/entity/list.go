package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type List struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"          validate:"-"`
	Name      string             `json:"name"         bson:"name"         validate:"required"`
	CardsIDs  []string           `json:"cardsIds"     bson:"cardsIds"     validate:"-"`
	CreatedAt time.Time          `json:"createdAt"    bson:"createdAt"    validate:"-"`
	UpdatedAt time.Time          `json:"updatedAt"    bson:"updatedAt"    validate:"-"`
}

func NewList(name string) (*List, error) {
	list := List{
		ID:        primitive.NewObjectID(),
		Name:      name,
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

func (list *List) Modify(name string) error {
	list.Name = name
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
