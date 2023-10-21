package board

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type FindByIDBoardInputDTO struct {
	ID string `json:"id"`
}

type FindByIDBoardOutputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type FindByIDBoardUseCaseInterface interface {
	Execute(input *FindByIDBoardInputDTO) (*FindByIDBoardOutputDTO, error)
}

type FindByIDBoardUseCase struct {
	BoardGateway gateway.BoardGateway
}

func NewBoardFindByIDUseCase(boardGateway gateway.BoardGateway) *FindByIDBoardUseCase {
	return &FindByIDBoardUseCase{BoardGateway: boardGateway}
}

func (useCase *FindByIDBoardUseCase) Execute(input *FindByIDBoardInputDTO) (*FindByIDBoardOutputDTO, error) {
	board, err := useCase.BoardGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindByIDBoardOutputDTO{
		ID:          "",
		Name:        board.Name,
		Description: board.Description,
		CreatedAt:   board.CreatedAt.String(),
		UpdatedAt:   board.UpdatedAt.String(),
	}, nil
}
