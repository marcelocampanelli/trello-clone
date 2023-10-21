package board

import (
	"fmt"
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
)

type CreateBoardInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserFounder string `json:"user_founder"`
}

type CreateBoardOutputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserFounder string `json:"user_founder"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateBoardUseCaseInterface interface {
	Execute(input *CreateBoardInputDTO) (*CreateBoardOutputDTO, error)
}

type CreateBoardUseCase struct {
	BoardGateway gateway.BoardGateway
}

func NewBoardCreateUseCase(boardGateway gateway.BoardGateway) *CreateBoardUseCase {
	return &CreateBoardUseCase{BoardGateway: boardGateway}
}

func (useCase *CreateBoardUseCase) Execute(input *CreateBoardInputDTO) (*CreateBoardOutputDTO, error) {
	board, err := entity.NewBoard(input.Name, input.Description, input.UserFounder)
	if err != nil {
		fmt.Println("BOARD USECASE: EXECUTE NEW BOARD", err)
		return nil, err
	}

	result, err := useCase.BoardGateway.Create(board)
	if err != nil {
		fmt.Println("BOARD USECASE: EXECUTE CREATE", err)
		return nil, err
	}

	output := &CreateBoardOutputDTO{
		ID:          *result,
		Name:        board.Name,
		Description: board.Description,
		UserFounder: board.UserFounder,
		CreatedAt:   board.CreatedAt.String(),
		UpdatedAt:   board.UpdatedAt.String(),
	}

	return output, nil
}
