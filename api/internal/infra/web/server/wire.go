//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
	"github.com/marcelocampanelli/trello-clone/internal/infra/repository"
	"github.com/marcelocampanelli/trello-clone/internal/infra/web/handlers"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var setUserRepositoryDependency = wire.NewSet(
	repository.NewUserRepository, wire.Bind(new(gateway.UserGateway), new(*repository.UserRepository)),
)

func InitializeUserHandler(client *mongo.Client) *handlers.UserHandler {
	wire.Build(
		setUserRepositoryDependency,
		handlers.NewUserHandler,
		wire.NewSet(user.NewUserCreateUseCase, wire.Bind(new(user.CreateUserUseCaseInterface), new(*user.CreateUserUseCase))),
		wire.NewSet(user.NewUserUpdateUseCase, wire.Bind(new(user.UpdateUserUseCaseInterface), new(*user.UpdateUserUseCase))),
	)

	return &handlers.UserHandler{}
}
