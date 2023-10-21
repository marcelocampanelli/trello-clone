package board

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestFindByIDBoardUseCase_Execute(t *testing.T) {
	m := &mocks.BoardGateway{}

	m.On("FindByID", mock.Anything).Return(&entity.Board{
		Name:        "Board 1",
		Description: "Board 1 description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil)

	useCase := NewBoardFindByIDUseCase(m)

	input := FindByIDBoardInputDTO{
		ID: "1",
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
}
