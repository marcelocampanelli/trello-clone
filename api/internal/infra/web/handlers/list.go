package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/list"
	"net/http"
)

type ListHandler struct {
	ucFindAll  list.FindAllListUseCaseInterface
	ucFindByID list.FindByIDListUseCaseInterface
	ucCreate   list.CreateListUseCaseInterface
	ucUpdate   list.UpdateListUseCaseInterface
	ucDelete   list.DeleteListUseCaseInterface
}

func NewListHandler(ucFindAll list.FindAllListUseCaseInterface, ucFindByID list.FindByIDListUseCaseInterface, ucCreate list.CreateListUseCaseInterface, ucUpdate list.UpdateListUseCaseInterface, ucDelete list.DeleteListUseCaseInterface) *ListHandler {
	return &ListHandler{
		ucFindAll:  ucFindAll,
		ucFindByID: ucFindByID,
		ucCreate:   ucCreate,
		ucUpdate:   ucUpdate,
		ucDelete:   ucDelete,
	}
}

func (handler *ListHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	var input list.FindAllListInputDTO

	input.BoardID = chi.URLParam(r, "boardID")

	output, err := handler.ucFindAll.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *ListHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	var input list.FindByIDListInputDTO

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

func (handler *ListHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input list.CreateListInputDTO

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (handler *ListHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input list.UpdateListInputDTO

	input.ID = chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := handler.ucUpdate.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *ListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input list.DeleteListInputDTO

	input.ID = chi.URLParam(r, "id")

	err := handler.ucDelete.Execute(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
