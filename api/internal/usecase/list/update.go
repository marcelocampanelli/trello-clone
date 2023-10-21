package list

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type UpdateListInputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type UpdateListOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

type UpdateListUseCaseInterface interface {
	Execute(input *UpdateListInputDTO) (*UpdateListOutputDTO, error)
}

type UpdateListUseCase struct {
	ListGateway gateway.ListGateway
}

func NewListUpdateUseCase(listGateway gateway.ListGateway) *UpdateListUseCase {
	return &UpdateListUseCase{ListGateway: listGateway}
}

func (useCase *UpdateListUseCase) Execute(input *UpdateListInputDTO) (*UpdateListOutputDTO, error) {
	list, err := useCase.ListGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	err = list.Modify(input.Name, input.Position)
	if err != nil {
		return nil, err
	}

	result, err := useCase.ListGateway.Update(input.ID, list)
	if err != nil {
		return nil, err
	}

	return &UpdateListOutputDTO{
		ID:       result.ID.String(),
		Name:     result.Name,
		Position: result.Position,
	}, nil

}
