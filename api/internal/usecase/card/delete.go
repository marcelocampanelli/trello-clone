package card

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type DeleteCardInputDTO struct {
	ID string `json:"id"`
}

type NewDeleteCardUseCaseInterface interface {
	Execute(input *DeleteCardInputDTO) error
}

type DeleteCardUseCase struct {
	CardGateway gateway.CardGateway
}

func NewCardDeleteUseCase(cardGateway gateway.CardGateway) *DeleteCardUseCase {
	return &DeleteCardUseCase{CardGateway: cardGateway}
}

func (useCase *DeleteCardUseCase) Execute(input *DeleteCardInputDTO) error {
	err := useCase.CardGateway.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
