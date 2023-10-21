package board

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	mocks "github.com/marcelocampanelli/trello-clone/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindAllBoardUseCase_Execute(t *testing.T) {
	m := &mocks.BoardGateway{}

	m.On("FindAll", "1").Return([]*entity.Board{
		{
			Name:        "Board 1",
			Description: "Board 1 description",
			UserFounder: "1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Board 2",
			Description: "Board 2 description",
			UserFounder: "1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil)

	useCase := NewBoardFindAllUseCase(m)

	input := FindAllBoardInputDTO{
		UserID: "1",
	}

	output, err := useCase.Execute(&input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, 2, len(output.Boards))
}
