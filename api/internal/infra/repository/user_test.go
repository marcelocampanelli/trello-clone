package repository

import (
	"github.com/marcelocampanelli/trello-clone/internal/domain/entity"
	"github.com/marcelocampanelli/trello-clone/internal/infra/database/mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestUserRepository(client)

	user, err := entity.NewUser("Jhon", "Doe", "jhondoe", "jhon@doe.com", "123456", "20271216093")
	assert.Nil(t, err)

	userID, err := repository.Create(user)
	assert.Nil(t, err)
	assert.NotNil(t, userID)
	assert.IsType(t, userID, new(string))
}

func TestUSerRepository_Update(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestUserRepository(client)

	user, err := entity.NewUser("Jhon", "Doe", "jhondoe", "jhon@doe.com", "123456", "20271216093")
	assert.Nil(t, err)

	userID, err := repository.Create(user)

	assert.Nil(t, err)
	assert.NotNil(t, userID)

	err = user.Modify("Ronaldo", "Carioca", "Ronaldinho")
	assert.Nil(t, err)

	err = repository.Update(*userID, user)

	assert.Nil(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestUserRepository(client)

	user, err := entity.NewUser("Jhon", "Doe", "jhondoe", "jhon@doe.com", "123456", "20271216093")
	assert.Nil(t, err)

	_, err = repository.Create(user)
	assert.Nil(t, err)

	user, err = repository.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, user, new(entity.User))
}

func TestUserRepository_FindByID(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestUserRepository(client)

	user, err := entity.NewUser("Jhon", "Doe", "jhondoe", "jhon@doe.com", "123456", "20271216093")
	assert.Nil(t, err)

	userID, err := repository.Create(user)
	assert.Nil(t, err)

	user, err = repository.FindByID(*userID)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, user, new(entity.User))
}

func TestUserRepository_FindByCPF(t *testing.T) {
	client, err := mongodb.ConnectMongoDB()
	assert.Nil(t, err)

	repository := NewTestUserRepository(client)

	user, err := entity.NewUser("Jhon", "Doe", "jhondoe", "jhon@doe.com", "123456", "20271216093")
	assert.Nil(t, err)

	_, err = repository.Create(user)
	assert.Nil(t, err)

	user, err = repository.FindByCPF(user.CPF)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, user, new(entity.User))
}
