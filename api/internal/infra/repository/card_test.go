package repository

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/infra/database/mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCardRepository_Create(t *testing.T) {
	cliet, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	card, err := entity.NewCard("Card 1", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestCardRepository(cliet)

	card, err = repository.Create(card)
	assert.Nil(t, err)
	assert.Equal(t, card.Name, "Card 1")
	assert.Equal(t, card.UserAssignedID, "5f9b3b3b1c9d440000b9c4a0")
	assert.Equal(t, card.ListID, "5f9b3b3b1c9d440000b9c4a0")
	assert.Equal(t, card.Position, 1)

	err = cliet.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewCardRepository_Update(t *testing.T) {
	cliet, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	card, err := entity.NewCard("Card 1", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	err = card.Modify("Card updated", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 2)
	assert.Nil(t, err)

	repository := NewTestCardRepository(cliet)

	card, err = repository.Update(card.ID.Hex(), card)
	assert.Nil(t, err)

	assert.Equal(t, card.Name, "Card updated")

	err = cliet.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewCardRepository_Delete(t *testing.T) {
	cliet, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	card, err := entity.NewCard("Card 1", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestCardRepository(cliet)

	err = repository.Delete(card.ID.Hex())

	assert.Nil(t, err)

	err = cliet.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewCardRepository_FindAllByList(t *testing.T) {
	cliet, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	card1, err := entity.NewCard("Card 1", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	card2, err := entity.NewCard("Card 2", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 2)
	assert.Nil(t, err)

	repository := NewTestCardRepository(cliet)

	_, err = repository.Create(card1)
	assert.Nil(t, err)

	_, err = repository.Create(card2)
	assert.Nil(t, err)

	cards, err := repository.FindAllByList(card1.ListID)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(cards))

	err = cliet.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewCardRepository_FindByID(t *testing.T) {
	cliet, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	card, err := entity.NewCard("Card 1", "5f9b3b3b1c9d440000b9c4a0", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestCardRepository(cliet)

	card, err = repository.Create(card)
	assert.Nil(t, err)

	cardFound, err := repository.FindByID(card.ID.Hex())
	assert.Nil(t, err)

	assert.Equal(t, card.ID, cardFound.ID)

	err = cliet.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}
