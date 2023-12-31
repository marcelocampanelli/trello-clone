package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/board"
	"net/http"
)

type BoardHandler struct {
	ucFindAll  board.FindAllBoardUseCaseInterface
	ucFindByID board.FindByIDBoardUseCaseInterface
	ucCreate   board.CreateBoardUseCaseInterface
	ucUpdate   board.UpdateBoardUseCaseInterface
	ucDelete   board.DeleteBoardUseCaseInterface
}

func NewBoardHandler(ucFindAll board.FindAllBoardUseCaseInterface, ucFindByID board.FindByIDBoardUseCaseInterface, ucCreate board.CreateBoardUseCaseInterface, ucUpdate board.UpdateBoardUseCaseInterface, ucDelete board.DeleteBoardUseCaseInterface) *BoardHandler {
	return &BoardHandler{
		ucFindAll:  ucFindAll,
		ucFindByID: ucFindByID,
		ucCreate:   ucCreate,
		ucUpdate:   ucUpdate,
		ucDelete:   ucDelete,
	}
}

func (handler *BoardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input board.CreateBoardInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := handler.ucCreate.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (handler *BoardHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	var input board.FindAllBoardInputDTO

	input.UserID = chi.URLParam(r, "userID")

	output, err := handler.ucFindAll.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *BoardHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	var input board.FindByIDBoardInputDTO

	input.ID = chi.URLParam(r, "id")

	output, err := handler.ucFindByID.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *BoardHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input board.UpdateBoardInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input.ID = chi.URLParam(r, "id")

	output, err := handler.ucUpdate.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *BoardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input board.DeleteBoardInputDTO

	input.ID = chi.URLParam(r, "id")

	err := handler.ucDelete.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("")
}
