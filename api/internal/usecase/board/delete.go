package board

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type DeleteBoardInputDTO struct {
	ID string `json:"id"`
}

type DeleteBoardUseCaseInterface interface {
	Execute(input *DeleteBoardInputDTO) error
}

type DeleteBoardUseCase struct {
	BoardGateway gateway.BoardGateway
}

func NewBoardDeleteUseCase(boardGateway gateway.BoardGateway) *DeleteBoardUseCase {
	return &DeleteBoardUseCase{BoardGateway: boardGateway}
}

func (useCase *DeleteBoardUseCase) Execute(input *DeleteBoardInputDTO) error {
	err := useCase.BoardGateway.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
