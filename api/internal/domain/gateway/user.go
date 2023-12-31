package gateway

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
)

type UserGateway interface {
	Create(user *entity.User) (*string, error)
	Update(id string, user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByCPF(cpf string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}
