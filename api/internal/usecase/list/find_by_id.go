package list

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type FindByIDListInputDTO struct {
	ID string `json:"id"`
}

type FindByIDListOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoardID   string `json:"board_id"`
	Position  int    `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FindByIDListUseCaseInterface interface {
	Execute(input *FindByIDListInputDTO) (*FindByIDListOutputDTO, error)
}

type FindByIDListUseCase struct {
	ListGateway gateway.ListGateway
}

func NewListFindByIDUseCase(listGateway gateway.ListGateway) *FindByIDListUseCase {
	return &FindByIDListUseCase{ListGateway: listGateway}
}

func (useCase *FindByIDListUseCase) Execute(input *FindByIDListInputDTO) (*FindByIDListOutputDTO, error) {
	list, err := useCase.ListGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindByIDListOutputDTO{
		ID:        list.ID.String(),
		Name:      list.Name,
		BoardID:   list.BoardID,
		Position:  list.Position,
		CreatedAt: list.CreatedAt.String(),
		UpdatedAt: list.UpdatedAt.String(),
	}, nil
}
