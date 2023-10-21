package user

import (
	"errors"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserGateway{}
	m.On("Create", mock.Anything).Return(nil)
	m.On("FindByEmail", mock.Anything).Return(nil, errors.New("not found"))
	m.On("FindByCPF", mock.Anything).Return(nil, errors.New("not found"))

	useCase := NewUserCreateUseCase(m)

	input := CreateUserInputDTO{
		FirstName: "Jhon",
		LastName:  "Doe",
		Nickname:  "jhondoe",
		Email:     "jhon@doe.com",
		Password:  "102030",
		CPF:       "35050817013",
	}

	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	m.AssertNumberOfCalls(t, "FindByEmail", 1)
	m.AssertNumberOfCalls(t, "FindByCPF", 1)
}

func TestCreateUserUseCase_Execute_EmailAlreadyExists(t *testing.T) {
	m := &mocks.UserGateway{}
	m.On("FindByEmail", mock.Anything).Return(&entity.User{}, nil)

	useCase := NewUserCreateUseCase(m)

	input := CreateUserInputDTO{
		FirstName: "Jhon",
		LastName:  "Doe",
		Nickname:  "jhondoe",
		Email:     "jhon@doe.com",
		Password:  "102030",
		CPF:       "35050817013",
	}

	output, err := useCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, ErrEmailAlreadyExists, err)
	assert.Equal(t, "email already exists", err.Error())

	m.AssertNumberOfCalls(t, "FindByEmail", 1)

}

func TestCreateUserUseCase_Execute_CPFAlreadyExists(t *testing.T) {
	m := &mocks.UserGateway{}
	m.On("FindByEmail", mock.Anything).Return(nil, errors.New("not found"))
	m.On("FindByCPF", mock.Anything).Return(&entity.User{}, nil)

	useCase := NewUserCreateUseCase(m)

	input := CreateUserInputDTO{
		FirstName: "Jhon",
		LastName:  "Doe",
		Nickname:  "jhondoe",
		Email:     "jhon@doe.com",
		Password:  "102030",
		CPF:       "35050817013",
	}

	output, err := useCase.Execute(input)
	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, ErrCPFAlreadyExists, err)
	assert.Equal(t, "cpf already exists", err.Error())

	m.AssertNumberOfCalls(t, "FindByCPF", 1)
}
