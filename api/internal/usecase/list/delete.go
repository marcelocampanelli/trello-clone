package list

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type DeleteListInputDTO struct {
	ID string `json:"id"`
}

type DeleteListUseCaseInterface interface {
	Execute(input *DeleteListInputDTO) error
}

type DeleteListUseCase struct {
	ListGateway gateway.ListGateway
}

func NewListDeleteUseCase(listGateway gateway.ListGateway) *DeleteListUseCase {
	return &DeleteListUseCase{ListGateway: listGateway}
}

func (useCase *DeleteListUseCase) Execute(input *DeleteListInputDTO) error {
	err := useCase.ListGateway.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
