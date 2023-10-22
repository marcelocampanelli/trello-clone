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

func TestFindByIDCardUseCase_Execute(t *testing.T) {
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

	useCase := NewCardFindByIDUseCase(m)

	input := &FindByIDCardInputDTO{
		ID: cardID.Hex(),
	}

	output, err := useCase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, cardID.Hex(), output.ID)
	assert.Equal(t, "Card 1", output.Name)
	assert.Equal(t, "5f7b1a9b9b9b9b9b9b9b9b9b", output.UserAssignedID)
	assert.Equal(t, 1, output.Position)
}
