package card

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/gateway"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FindAllByListCardInputDTO struct {
	ListID string `json:"list_id"`
}

type FindAllByListCardOutputDTO struct {
	Cards []*Card `json:"cards"`
}

type Card struct {
	ID             primitive.ObjectID `json:"id"`
	Name           string             `json:"name"`
	UserAssignedID string             `json:"user_assigned_id"`
	ListID         string             `json:"list_id"`
	Position       int                `json:"position"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type FindAllByListCardUseCaseInterface interface {
	Execute(input *FindAllByListCardInputDTO) (*FindAllByListCardOutputDTO, error)
}

type FindAllByListCardUseCase struct {
	CardGateway gateway.CardGateway
}

func NewCardFindAllByListUseCase(cardGateway gateway.CardGateway) *FindAllByListCardUseCase {
	return &FindAllByListCardUseCase{CardGateway: cardGateway}
}

func (useCase *FindAllByListCardUseCase) Execute(input *FindAllByListCardInputDTO) (*FindAllByListCardOutputDTO, error) {
	cards, err := useCase.CardGateway.FindAllByList(input.ListID)
	if err != nil {
		return nil, err
	}

	var cardsDTO []*Card
	for _, card := range cards {
		cardsDTO = append(cardsDTO, &Card{
			ID:             card.ID,
			Name:           card.Name,
			UserAssignedID: card.UserAssignedID,
			ListID:         card.ListID,
			Position:       card.Position,
			CreatedAt:      card.CreatedAt,
			UpdatedAt:      card.UpdatedAt,
		})
	}

	return &FindAllByListCardOutputDTO{
		Cards: cardsDTO,
	}, nil
}
