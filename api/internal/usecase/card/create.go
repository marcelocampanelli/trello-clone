package card

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
)

type CreateCardInputDTO struct {
	Name           string `json:"name"`
	UserAssignedID string `json:"user_assigned_id"`
	ListID         string `json:"list_id"`
	Position       int    `json:"position"`
}

type CreateCardOutputDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	UserAssignedID string `json:"user_assigned_id"`
	ListID         string `json:"list_id"`
	Position       int    `json:"position"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CreateCardUseCaseInterface interface {
	Execute(input *CreateCardInputDTO) (*CreateCardOutputDTO, error)
}

type CreateCardUseCase struct {
	CardGateway gateway.CardGateway
}

func NewCardCreateUseCase(cardGateway gateway.CardGateway) *CreateCardUseCase {
	return &CreateCardUseCase{CardGateway: cardGateway}
}

func (useCase *CreateCardUseCase) Execute(input *CreateCardInputDTO) (*CreateCardOutputDTO, error) {
	card, err := entity.NewCard(input.Name, input.UserAssignedID, input.ListID, input.Position)
	if err != nil {
		return nil, err
	}

	result, err := useCase.CardGateway.Create(card)
	if err != nil {
		return nil, err
	}

	return &CreateCardOutputDTO{
		ID:             result.ID.Hex(),
		Name:           result.Name,
		UserAssignedID: result.UserAssignedID,
		Position:       result.Position,
		ListID:         result.ListID,
		CreatedAt:      result.CreatedAt.String(),
		UpdatedAt:      result.UpdatedAt.String(),
	}, nil
}
