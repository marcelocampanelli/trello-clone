package gateway

import "github.com/marcelocampanelli/trello-clone/internal/domain/entity"

type ListGateway interface {
	FindAll(boardID string) ([]*entity.List, error)
	FindByID(id string) (*entity.List, error)
	Create(list *entity.List) (*entity.List, error)
	Update(id string, list *entity.List) (*entity.List, error)
	Delete(id string) error
}
