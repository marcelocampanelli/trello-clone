package repository

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/infra/database/mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListRepository_Create(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	list, err := entity.NewList("List 1", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestListRepository(client)

	list, err = repository.Create(list)
	assert.Nil(t, err)

	err = client.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewListRepository_Update(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	list, err := entity.NewList("List 1", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestListRepository(client)

	list, err = repository.Create(list)
	assert.Nil(t, err)

	err = list.Modify("List updated", 2)
	assert.Nil(t, err)

	list, err = repository.Update(list.ID.Hex(), list)

	assert.Nil(t, err)
	assert.Equal(t, list.Name, "List updated")
	assert.Equal(t, list.Position, 2)

	err = client.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewListRepository_Delete(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	list, err := entity.NewList("List 1", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestListRepository(client)

	list, err = repository.Create(list)
	assert.Nil(t, err)

	err = repository.Delete(list.ID.Hex())
	assert.Nil(t, err)

	err = client.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewListRepository_FindByID(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	list, err := entity.NewList("List 1", "5f9b3b3b1c9d440000b9c4a0", 1)
	assert.Nil(t, err)

	repository := NewTestListRepository(client)

	list, err = repository.Create(list)
	assert.Nil(t, err)

	list, err = repository.FindByID(list.ID.Hex())
	assert.Nil(t, err)
	assert.NotNil(t, list)

	err = client.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}

func TestNewListRepository_FindAll(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestListRepository(client)

	list1, err := entity.NewList("List 1", "5f9b3b3b1c9d440000b9c4a9", 1)
	assert.Nil(t, err)
	assert.NotNil(t, list1)

	list2, err := entity.NewList("List 2", "5f9b3b3b1c9d440000b9c4a9", 1)
	assert.Nil(t, err)
	assert.NotNil(t, list2)

	_, err = repository.Create(list1)
	assert.Nil(t, err)

	_, err = repository.Create(list2)
	assert.Nil(t, err)

	lists, err := repository.FindAll("5f9b3b3b1c9d440000b9c4a9")
	assert.Nil(t, err)
	assert.NotNil(t, lists)
	assert.Equal(t, len(lists), 2)

	err = client.Database("trello-clone-test").Drop(nil)
	assert.Nil(t, err)
}
