package board

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
)

type UpdateBoardInputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateBoardOutputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateBoardUseCaseInterface interface {
	Execute(input *UpdateBoardInputDTO) (*UpdateBoardOutputDTO, error)
}

type UpdateBoardUseCase struct {
	BoardGateway gateway.BoardGateway
}

func NewBoardUpdateUseCase(boardGateway gateway.BoardGateway) *UpdateBoardUseCase {
	return &UpdateBoardUseCase{BoardGateway: boardGateway}
}

func (useCase *UpdateBoardUseCase) Execute(input *UpdateBoardInputDTO) (*UpdateBoardOutputDTO, error) {
	board, err := useCase.BoardGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	err = board.Modify(input.Name, input.Description)
	if err != nil {
		return nil, err
	}

	err = useCase.BoardGateway.Update(input.ID, board)
	if err != nil {
		return nil, err
	}

	return &UpdateBoardOutputDTO{
		ID:          input.ID,
		Name:        board.Name,
		Description: board.Description,
		UpdatedAt:   board.UpdatedAt.String(),
	}, nil
}
