package list

import (
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteListUseCase_Execute(t *testing.T) {
	m := &mocks.ListGateway{}

	m.On("Delete", mock.Anything).Return(nil)

	useCase := NewListDeleteUseCase(m)

	input := DeleteListInputDTO{
		ID: "1",
	}

	err := useCase.Execute(&input)

	assert.Nil(t, err)
}
