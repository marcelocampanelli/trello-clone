package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/user"
	"net/http"
)

type UserHandler struct {
	ucCreate user.CreateUserUseCaseInterface
	ucUpdate user.UpdateUserUseCaseInterface
}

func NewUserHandler(ucCreate user.CreateUserUseCaseInterface, ucUpdate user.UpdateUserUseCaseInterface) *UserHandler {
	return &UserHandler{
		ucCreate: ucCreate,
		ucUpdate: ucUpdate,
	}
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input user.CreateUserInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("User Handler: error to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := handler.ucCreate.Execute(&input)
	if err != nil {
		fmt.Println("User Handler: error to execute use case", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input user.UpdateUserInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("User Handler: error to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := handler.ucUpdate.Execute(&input)
	if err != nil {
		fmt.Println("User Handler: error to execute use case", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
