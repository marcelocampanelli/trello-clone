package list

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
)

type CreateListInputDTO struct {
	Name     string `json:"name"`
	BoardID  string `json:"board_id"`
	Position int    `json:"position"`
}

type CreateListOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoardID   string `json:"board_id"`
	Position  int    `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateListUseCaseInterface interface {
	Execute(input *CreateListInputDTO) (*CreateListOutputDTO, error)
}

type CreateListUseCase struct {
	ListGateway gateway.ListGateway
}

func NewListCreateUseCase(listGateway gateway.ListGateway) *CreateListUseCase {
	return &CreateListUseCase{ListGateway: listGateway}
}

func (useCase *CreateListUseCase) Execute(input *CreateListInputDTO) (*CreateListOutputDTO, error) {
	list, err := entity.NewList(input.Name, input.BoardID, input.Position)
	if err != nil {
		return nil, err
	}

	result, err := useCase.ListGateway.Create(list)
	if err != nil {
		return nil, err
	}

	return &CreateListOutputDTO{
		ID:        result.ID.String(),
		Name:      result.Name,
		BoardID:   result.BoardID,
		Position:  result.Position,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	}, nil
}
