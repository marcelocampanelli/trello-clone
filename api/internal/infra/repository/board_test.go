package repository

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/infra/database/mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoardRepository_Create(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestBoardRepository(client)

	board, err := entity.NewBoard("Board 1", "Test", "5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	result, err := repository.Create(board)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.IsType(t, result, new(string))

	client.Database("trello-clone-test").Drop(nil)
}

func TestBoardRepository_Update(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestBoardRepository(client)

	board, err := entity.NewBoard("Board 1", "Test", "5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	result, err := repository.Create(board)
	assert.Nil(t, err)

	boardFinded, err := repository.FindByID(*result)
	assert.Nil(t, err)

	err = boardFinded.Modify("Board 2", "Test 2")
	assert.Nil(t, err)

	err = repository.Update(*result, boardFinded)
	assert.Nil(t, err)

	client.Database("trello-clone-test").Drop(nil)
}

func TestBoardRepository_FindByID(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestBoardRepository(client)

	board, err := entity.NewBoard("Board 1", "Test", "5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	result, err := repository.Create(board)
	assert.Nil(t, err)

	boardFinded, err := repository.FindByID(*result)
	assert.Nil(t, err)
	assert.NotNil(t, boardFinded)

	client.Database("trello-clone-test").Drop(nil)
}

func TestBoardRepository_Delete(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestBoardRepository(client)

	board, err := entity.NewBoard("Board 1", "Test", "5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	result, err := repository.Create(board)
	assert.Nil(t, err)

	err = repository.Delete(*result)
	assert.Nil(t, err)

	client.Database("trello-clone-test").Drop(nil)
}

func TestBoardRepository_FindAll(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestBoardRepository(client)

	board1, err := entity.NewBoard("Board 1", "Test", "5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	board2, err := entity.NewBoard("Board 2", "Test", "5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	_, err = repository.Create(board1)
	assert.Nil(t, err)

	_, err = repository.Create(board2)
	assert.Nil(t, err)

	boards, err := repository.FindAll("5f9b3b3b1c9d440000b9c4a0")
	assert.Nil(t, err)

	assert.NotNil(t, boards)
	assert.Equal(t, len(boards), 2)

	client.Database("trello-clone-test").Drop(nil)
}
