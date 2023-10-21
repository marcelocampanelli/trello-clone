//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
	"github.com/marcelocampanelli/trello-clone/internal/infra/repository"
	"github.com/marcelocampanelli/trello-clone/internal/infra/web/handlers"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/board"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/list"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var setUserRepositoryDependency = wire.NewSet(
	repository.NewUserRepository, wire.Bind(new(gateway.UserGateway), new(*repository.UserRepository)),
)

var setBoardRepositoryDependency = wire.NewSet(
	repository.NewBoardRepository, wire.Bind(new(gateway.BoardGateway), new(*repository.BoardRepository)),
)

var setListRepositoryDependency = wire.NewSet(
	repository.NewListRepository, wire.Bind(new(gateway.ListGateway), new(*repository.ListRepository)),
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

func InitializeListHandler(client *mongo.Client) *handlers.ListHandler {
	wire.Build(
		setListRepositoryDependency,
		handlers.NewListHandler,
		wire.NewSet(list.NewListCreateUseCase, wire.Bind(new(list.CreateListUseCaseInterface), new(*list.CreateListUseCase))),
		wire.NewSet(list.NewListFindByIDUseCase, wire.Bind(new(list.FindByIDListUseCaseInterface), new(*list.FindByIDListUseCase))),
		wire.NewSet(list.NewListFindAllUseCase, wire.Bind(new(list.FindAllListUseCaseInterface), new(*list.FindAllListUseCase))),
		wire.NewSet(list.NewListUpdateUseCase, wire.Bind(new(list.UpdateListUseCaseInterface), new(*list.UpdateListUseCase))),
		wire.NewSet(list.NewListDeleteUseCase, wire.Bind(new(list.DeleteListUseCaseInterface), new(*list.DeleteListUseCase))),
	)
	return &handlers.ListHandler{}
}
