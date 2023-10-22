package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/marcelocampanelli/trello-clone/internal/usecase/card"
	"net/http"
)

type CardHandler struct {
	ucCreate        card.CreateCardUseCaseInterface
	ucUpdate        card.UpdateCardUseCaseInterface
	ucDelete        card.DeleteCardUseCaseInterface
	ucFindAllByList card.FindAllByListCardUseCaseInterface
	ucFindByID      card.FindByIDCardUseCaseInterface
}

func NewCardHandler(ucCreate card.CreateCardUseCaseInterface, ucUpdate card.UpdateCardUseCaseInterface, ucDelete card.DeleteCardUseCaseInterface, ucFindAllByList card.FindAllByListCardUseCaseInterface, ucFindByID card.FindByIDCardUseCaseInterface) *CardHandler {
	return &CardHandler{
		ucCreate:        ucCreate,
		ucUpdate:        ucUpdate,
		ucDelete:        ucDelete,
		ucFindAllByList: ucFindAllByList,
		ucFindByID:      ucFindByID,
	}
}

func (handler *CardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto card.CreateCardInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := handler.ucCreate.Execute(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (handler *CardHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto card.UpdateCardInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dto.ID = chi.URLParam(r, "id")

	output, err := handler.ucUpdate.Execute(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *CardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var dto card.DeleteCardInputDTO

	dto.ID = chi.URLParam(r, "id")

	err := handler.ucDelete.Execute(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *CardHandler) FindAllByList(w http.ResponseWriter, r *http.Request) {
	var dto card.FindAllByListCardInputDTO

	dto.ListID = chi.URLParam(r, "list_id")

	output, err := handler.ucFindAllByList.Execute(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (handler *CardHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	var dto card.FindByIDCardInputDTO

	dto.ID = chi.URLParam(r, "id")

	output, err := handler.ucFindByID.Execute(&dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
