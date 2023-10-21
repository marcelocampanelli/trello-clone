package handlers

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := handler.ucCreate.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (handler *BoardHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	var input board.FindAllBoardInputDTO

	input.UserID = chi.URLParam(r, "user_id")

	output, err := handler.ucFindAll.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (handler *BoardHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	var input board.FindByIDBoardInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input.ID = chi.URLParam(r, "id")

	output, err := handler.ucFindByID.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}
