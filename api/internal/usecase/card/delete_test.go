package card

import (
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteCardUseCase_Execute(t *testing.T) {
	m := &mocks.CardGateway{}

	m.On("Delete", mock.Anything).Return(nil)

	useCase := NewDeleteUseCase(m)

	input := DeleteCardInputDTO{
		ID: "1",
	}

	err := useCase.Execute(&input)
	assert.Nil(t, err)
}
