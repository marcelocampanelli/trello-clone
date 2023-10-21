package list

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestUpdateListUseCase_Execute(t *testing.T) {
	m := &mocks.ListGateway{}

	m.On("FindByID", mock.Anything).Return(&entity.List{
		ID:        primitive.NewObjectID(),
		BoardID:   "1",
		Name:      "List 1",
		CardsIDs:  []string{},
		Position:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	m.On("Update", mock.Anything, mock.Anything).Return(&entity.List{
		ID:        primitive.NewObjectID(),
		BoardID:   "1",
		Name:      "List 1",
		Position:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	useCase := NewListUpdateUseCase(m)

	input := UpdateListInputDTO{
		ID:       "1",
		Name:     "List 1",
		Position: 1,
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
}
