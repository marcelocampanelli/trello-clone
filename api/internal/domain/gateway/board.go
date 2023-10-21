package gateway

import "github.com/marcelocampanelli/trello-clone/internal/domain/entity"

type BoardGateway interface {
	FindAll(userID string) ([]*entity.Board, error)
	Create(board *entity.Board) (*entity.Board, error)
	Update(id string, board *entity.Board) (*entity.Board, error)
	FindOne(id string) (*entity.Board, error)
	Delete(id string) error
}
