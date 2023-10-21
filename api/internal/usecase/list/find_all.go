package list

import "github.com/marcelocampanelli/trello-clone/internal/domain/gateway"

type FindAllListInputDTO struct {
	BoardID string `json:"board_id"`
}

type FindAllListOutputDTO struct {
	Lists []*ListCollection `json:"lists"`
}

type ListCollection struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Position int      `json:"position"`
	CardsIDs []string `json:"cards_ids"`
}

type FindAllListUseCaseInterface interface {
	Execute(input *FindAllListInputDTO) (*FindAllListOutputDTO, error)
}

type FindAllListUseCase struct {
	ListGateway gateway.ListGateway
}

func NewListFindAllUseCase(listGateway gateway.ListGateway) *FindAllListUseCase {
	return &FindAllListUseCase{ListGateway: listGateway}
}

func (useCase *FindAllListUseCase) Execute(input *FindAllListInputDTO) (*FindAllListOutputDTO, error) {
	lists, err := useCase.ListGateway.FindAll(input.BoardID)
	if err != nil {
		return nil, err
	}

	var listsCollection []*ListCollection
	for _, list := range lists {
		listsCollection = append(listsCollection, &ListCollection{
			ID:       list.ID.Hex(),
			Name:     list.Name,
			Position: list.Position,
			CardsIDs: list.CardsIDs,
		})
	}

	return &FindAllListOutputDTO{Lists: listsCollection}, nil
}
