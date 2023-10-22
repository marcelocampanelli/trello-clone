package card

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestUpdateCardUseCase_Execute(t *testing.T) {
	cardID := primitive.NewObjectID()

	m := &mocks.CardGateway{}

	m.On("FindByID", mock.Anything).Return(&entity.Card{
		ID:             cardID,
		Name:           "Card 1",
		UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
		Position:       1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil)

	m.On("Update", mock.Anything, mock.Anything).Return(&entity.Card{
		ID:             cardID,
		Name:           "Card Updated",
		UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
		Position:       1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil)

	useCase := NewCardUpdateUseCase(m)

	input := &UpdateCardInputDTO{
		ID:             cardID.Hex(),
		Name:           "Card Updated",
		UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
		Position:       1,
	}

	result, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
