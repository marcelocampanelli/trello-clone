package gateway

import "github.com/marcelocampanelli/trello-clone/internal/domain/entity"

type CardGateway interface {
	Create(card *entity.Card) (*entity.Card, error)
	Update(id string, card *entity.Card) (*entity.Card, error)
	Delete(id string) error
	FindByID(id string) (*entity.Card, error)
	FindAllByList(listID string) ([]*entity.Card, error)
}
