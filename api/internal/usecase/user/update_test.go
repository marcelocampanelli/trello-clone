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
	m.On("FindByID", "1").Return(&entity.User{}, nil)
	m.On("Update", mock.Anything).Return(nil)

	useCase := NewUserUpdateUseCase(m)

	input := UpdateUserInputDTO{
		ID:        "1",
		FirstName: "Jhon",
		LastName:  "Doe",
		Nickname:  "jhondoe",
	}

	output, err := useCase.Execute(&input)

	assert.NotNil(t, output)
	assert.Nil(t, err)

	m.AssertNumberOfCalls(t, "FindByID", 1)
	m.AssertNumberOfCalls(t, "Update", 1)
}
