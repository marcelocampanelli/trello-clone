package gateway

import "github.com/marcelocampanelli/trello-clone/internal/domain/entity"

type BoardGateway interface {
	FindAll(userID string) ([]*entity.Board, error)
	Create(board *entity.Board) (*string, error)
	Update(id string, board *entity.Board) error
	FindByID(id string) (*entity.Board, error)
	Delete(id string) error
}
