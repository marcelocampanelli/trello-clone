package board

import (
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateBoardUseCase_Execute(t *testing.T) {
	userID := "1"
	boardID := "2"

	m := &mocks.BoardGateway{}
	m.On("Create", mock.Anything).Return(&boardID, nil)

	useCase := NewBoardCreateUseCase(m)

	input := CreateBoardInputDTO{
		Name:        "Test",
		Description: "Test",
		UserFounder: userID,
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, boardID, output.ID)

	m.AssertNumberOfCalls(t, "Create", 1)
}
