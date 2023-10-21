package user

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUpdateUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserGateway{}

	m.On("FindByID", mock.Anything).Return(&entity.User{
		FirstName: "Jhon",
		LastName:  "Doe",
		Nickname:  "jhondoe",
		Email:     "jhon@doe.com",
		Password:  "102030",
		CPF:       "35050817013",
	}, nil)

	m.On("Update", mock.Anything, mock.Anything).Return(nil)

	useCase := NewUserUpdateUseCase(m)

	input := UpdateUserInputDTO{
		ID:        "1",
		FirstName: "Jhon",
		LastName:  "Doe",
		Nickname:  "jhondoe",
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

}
