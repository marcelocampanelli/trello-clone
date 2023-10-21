package gateway

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
)

type UserGateway interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id int32) (*entity.User, error)
}
