package card

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type FindByIDCardInputDTO struct {
	ID string `json:"id"`
}

type FindByIDCardOutputDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	UserAssignedID string `json:"user_assigned_id"`
	Position       int    `json:"position"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type FindByIDCardUseCaseInterface interface {
	Execute(input *FindByIDCardInputDTO) (*FindByIDCardOutputDTO, error)
}

type FindByIDCardUseCase struct {
	CardGateway gateway.CardGateway
}

func NewCardFindByIDUseCase(cardGateway gateway.CardGateway) *FindByIDCardUseCase {
	return &FindByIDCardUseCase{CardGateway: cardGateway}
}

func (useCase *FindByIDCardUseCase) Execute(input *FindByIDCardInputDTO) (*FindByIDCardOutputDTO, error) {
	card, err := useCase.CardGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindByIDCardOutputDTO{
		ID:             card.ID.Hex(),
		Name:           card.Name,
		UserAssignedID: card.UserAssignedID,
		Position:       card.Position,
		CreatedAt:      card.CreatedAt.String(),
		UpdatedAt:      card.UpdatedAt.String(),
	}, nil
}
