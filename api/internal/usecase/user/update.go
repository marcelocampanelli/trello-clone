package user

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type UpdateUserInputDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
}

type UpdateUserOutputDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUserUseCaseInterface interface {
	Execute(input *UpdateUserInputDTO) (*UpdateUserOutputDTO, error)
}

type UpdateUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUserUpdateUseCase(userGateway gateway.UserGateway) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserGateway: userGateway}
}

func (useCase *UpdateUserUseCase) Execute(input *UpdateUserInputDTO) (*UpdateUserOutputDTO, error) {
	user, err := useCase.UserGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	user.Modify(input.FirstName, input.LastName, input.Nickname)

	err = useCase.UserGateway.Update(user)
	if err != nil {
		return nil, err
	}

	return &UpdateUserOutputDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CPF:       user.CPF,
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
