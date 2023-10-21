package handlers

import "github.com/marcelocampanelli/trello-clone/internal/usecase/user"

type UserHandler struct {
	ucCreate user.CreateUserUseCaseInterface
}
