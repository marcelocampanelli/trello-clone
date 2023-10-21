package board

import (
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteBoardUseCase_Execute(t *testing.T) {
	m := &mocks.BoardGateway{}

	m.On("Delete", mock.Anything).Return(nil)

	useCase := NewBoardDeleteUseCase(m)

	input := DeleteBoardInputDTO{
		ID: "1",
	}

	err := useCase.Execute(&input)

	assert.Nil(t, err)
}
