package board

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUpdateBoardUseCase_Execute(t *testing.T) {
	m := &mocks.BoardGateway{}

	m.On("FindByID", mock.Anything).Return(&entity.Board{
		Name:        "Board 1",
		Description: "Board 1 description",
		UserFounder: "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil)

	m.On("Update", mock.Anything, mock.Anything).Return(nil)

	useCase := NewBoardUpdateUseCase(m)

	input := UpdateBoardInputDTO{
		ID:          "1",
		Name:        "Board 1 updated",
		Description: "Board 1 description updated",
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Description, output.Description)

	m.AssertNumberOfCalls(t, "FindByID", 1)
	m.AssertNumberOfCalls(t, "Update", 1)
}
