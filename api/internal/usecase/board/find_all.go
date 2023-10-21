package board

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type FindAllBoardInputDTO struct {
	UserID string `json:"user_id"`
}

type FindAllBoardOutputDTO struct {
	Boards []*BoardReturnDTO `json:"boards"`
}

type BoardReturnDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type FindAllBoardUseCaseInterface interface {
	Execute(input *FindAllBoardInputDTO) (*FindAllBoardOutputDTO, error)
}

type FindAllBoardUseCase struct {
	BoardGateway gateway.BoardGateway
}

func NewBoardFindAllUseCase(boardGateway gateway.BoardGateway) *FindAllBoardUseCase {
	return &FindAllBoardUseCase{BoardGateway: boardGateway}
}

func (useCase *FindAllBoardUseCase) Execute(input *FindAllBoardInputDTO) (*FindAllBoardOutputDTO, error) {
	boards, err := useCase.BoardGateway.FindAll(input.UserID)
	if err != nil {
		return nil, err
	}

	var boardsReturn []*BoardReturnDTO

	for _, board := range boards {
		boardsReturn = append(boardsReturn, &BoardReturnDTO{
			ID:          "",
			Name:        board.Name,
			Description: board.Description,
			CreatedAt:   board.CreatedAt.String(),
			UpdatedAt:   board.UpdatedAt.String(),
		})
	}

	return &FindAllBoardOutputDTO{
		Boards: boardsReturn,
	}, nil
}
