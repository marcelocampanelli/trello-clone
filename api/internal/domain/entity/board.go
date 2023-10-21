package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Board struct {
	ID          primitive.ObjectID `json:"id"            bson:"_id"          validate:"-"`
	Name        string             `json:"name"          bson:"name"         validate:"required"`
	Description string             `json:"description"   bson:"description"  validate:"-"`
	UserFounder string             `json:"user_founder"  bson:"user_founder" validate:"required"`
	ListsIds    []string           `json:"lists_ids"     bson:"lists_ids"    validate:"-"`
	CreatedAt   time.Time          `json:"created_at"    bson:"created_at"   validate:"-"`
	UpdatedAt   time.Time          `json:"updated_at"    bson:"updated_at"   validate:"-"`
}

func NewBoard(name, description, user_founder string) (*Board, error) {
	board := Board{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
		UserFounder: user_founder,
		ListsIds:    []string{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := board.isValid()
	if err != nil {
		return nil, err
	}

	return &board, nil
}

func (board *Board) Modify(name, description string) error {
	board.Name = name
	board.Description = description
	board.UpdatedAt = time.Now()

	err := board.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (board *Board) isValid() error {
	validate := validator.New()

	err := validate.Struct(board)

	if err != nil {
		return err
	}

	return nil
}
