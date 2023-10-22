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

func TestCreateCardUseCase_Execute(t *testing.T) {
	m := &mocks.CardGateway{}

	m.On("Create", mock.Anything).Return(&entity.Card{
		ID:             primitive.NewObjectID(),
		Name:           "Card 1",
		UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
		ListID:         "5f7b1a9b9b9b9b9b9b9b9b9b",
		Position:       1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil)

	useCase := NewCardCreateUseCase(m)

	input := &CreateCardInputDTO{
		Name:           "Card 1",
		UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
		ListID:         "5f7b1a9b9b9b9b9b9b9b9b9b",
		Position:       1,
	}

	result, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
