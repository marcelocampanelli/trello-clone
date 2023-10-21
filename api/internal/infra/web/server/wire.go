//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
	"github.com/marcelocampanelli/trello-clone/internal/infra/repository"
	"github.com/marcelocampanelli/trello-clone/internal/infra/web/handlers"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/board"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var setUserRepositoryDependency = wire.NewSet(
	repository.NewUserRepository, wire.Bind(new(gateway.UserGateway), new(*repository.UserRepository)),
)

var setBoardRepositoryDependency = wire.NewSet(
	repository.NewBoardRepository, wire.Bind(new(gateway.BoardGateway), new(*repository.BoardRepository)),
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

func InitializeBoardHandler(client *mongo.Client) *handlers.BoardHandler {
	wire.Build(
		setBoardRepositoryDependency,
		handlers.NewBoardHandler,
		wire.NewSet(board.NewBoardCreateUseCase, wire.Bind(new(board.CreateBoardUseCaseInterface), new(*board.CreateBoardUseCase))),
		wire.NewSet(board.NewBoardFindAllUseCase, wire.Bind(new(board.FindAllBoardUseCaseInterface), new(*board.FindAllBoardUseCase))),
		wire.NewSet(board.NewBoardFindByIDUseCase, wire.Bind(new(board.FindByIDBoardUseCaseInterface), new(*board.FindByIDBoardUseCase))),
		wire.NewSet(board.NewBoardUpdateUseCase, wire.Bind(new(board.UpdateBoardUseCaseInterface), new(*board.UpdateBoardUseCase))),
		wire.NewSet(board.NewBoardDeleteUseCase, wire.Bind(new(board.DeleteBoardUseCaseInterface), new(*board.DeleteBoardUseCase))),
	)

	return &handlers.BoardHandler{}
}
