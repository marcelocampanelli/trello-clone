package list

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestFindByIDListUseCase_Execute(t *testing.T) {
	m := &mocks.ListGateway{}

	m.On("FindByID", "1").Return(&entity.List{
		ID:        primitive.NewObjectID(),
		BoardID:   "1",
		Name:      "List 1",
		Position:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	useCase := NewListFindByIDUseCase(m)

	input := FindByIDListInputDTO{
		ID: "1",
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
}
