package list

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestCreateListUseCase_Execute(t *testing.T) {
	m := &mocks.ListGateway{}
	m.On("Create", mock.Anything).Return(&entity.List{
		ID:       primitive.NewObjectID(),
		BoardID:  "1",
		Name:     "List 1",
		Position: 1,
	}, nil)

	useCase := NewListCreateUseCase(m)

	input := CreateListInputDTO{
		BoardID:  "1",
		Name:     "List 1",
		Position: 1,
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
}
