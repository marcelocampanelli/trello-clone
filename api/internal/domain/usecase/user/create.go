package user

import (
	"errors"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCPFAlreadyExists   = errors.New("cpf already exists")
	ErrCreateUser         = errors.New("error on create user")
)

type CreateUserInputDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
}

type CreateUserOutputDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUserUseCaseInterface interface {
	Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error)
}

type CreateUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserCreateUseCase(userGateway gateway.UserGateway) *CreateUserUseCase {
	return &CreateUserUseCase{UserGateway: userGateway}
}

func (useCase *CreateUserUseCase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	_, err := useCase.UserGateway.FindByEmail(input.Email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}

	_, err = useCase.UserGateway.FindByCPF(input.CPF)
	if err == nil {
		return nil, ErrCPFAlreadyExists
	}

	user, err := entity.NewUser(input.FirstName, input.LastName, input.Nickname, input.Email, input.Password, input.CPF)
	if err != nil {
		return nil, ErrCreateUser
	}

	err = useCase.UserGateway.Create(user)

	return &CreateUserOutputDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CPF:       user.CPF,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
