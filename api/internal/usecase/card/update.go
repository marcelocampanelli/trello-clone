package card

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type UpdateCardInputDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	UserAssignedID string `json:"user_assigned_id"`
	Position       int    `json:"position"`
}

type UpdateCardOutputDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	UserAssignedID string `json:"user_assigned_id"`
	Position       int    `json:"position"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type UpdateCardUseCaseInterface interface {
	Execute(input *UpdateCardInputDTO) (*UpdateCardOutputDTO, error)
}

type UpdateCardUseCase struct {
	CardGateway gateway.CardGateway
}

func NewCardUpdateUseCase(cardGateway gateway.CardGateway) *UpdateCardUseCase {
	return &UpdateCardUseCase{CardGateway: cardGateway}
}

func (useCase *UpdateCardUseCase) Execute(input *UpdateCardInputDTO) (*UpdateCardOutputDTO, error) {
	card, err := useCase.CardGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	err = card.Modify(input.Name, input.UserAssignedID, input.Position)
	if err != nil {
		return nil, err
	}

	result, err := useCase.CardGateway.Update(input.ID, card)
	if err != nil {
		return nil, err
	}

	return &UpdateCardOutputDTO{
		ID:             result.ID.Hex(),
		Name:           result.Name,
		UserAssignedID: result.UserAssignedID,
		Position:       result.Position,
		CreatedAt:      result.CreatedAt.String(),
		UpdatedAt:      result.UpdatedAt.String(),
	}, nil
}
