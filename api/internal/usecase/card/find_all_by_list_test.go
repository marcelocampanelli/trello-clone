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

func TestNewFindByIDCardUseCase_Execute(t *testing.T) {
	m := &mocks.CardGateway{}

	m.On("FindAllByList", mock.Anything).Return([]*entity.Card{
		{
			ID:             primitive.NewObjectID(),
			Name:           "Card 1",
			UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
			ListID:         "5f7b1a9b9b9b9b9b11244241",
			Position:       1,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             primitive.NewObjectID(),
			Name:           "Card 2",
			UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
			ListID:         "5f7b1a9b9b9b9b9b11244241",
			Position:       1,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             primitive.NewObjectID(),
			Name:           "Card 3",
			UserAssignedID: "5f7b1a9b9b9b9b9b9b9b9b9b",
			ListID:         "5f7b1a9b9b9b9b9b11244241",
			Position:       1,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}, nil)

	useCase := NewCardFindAllByListUseCase(m)

	input := &FindAllByListCardInputDTO{
		ListID: "5f7b1a9b9b9b9b9b11244241",
	}

	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, 3, len(output.Cards))
}
