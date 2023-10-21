package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_NewUser(t *testing.T) {
	user, err := NewUser("John", "Doe", "johndoe", "jhon@doe.com", "123456", "35050817013")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.Password)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)

	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "johndoe", user.Nickname)
	assert.Equal(t, "35050817013", user.CPF)

	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}

func TestUser_NewUser_InvalidCPF(t *testing.T) {
	user, err := NewUser("John", "Doe", "johndoe", "jhon@doe.com", "123456", "12345678910")
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "invalid CPF", err.Error())
}

func TestUser_NewUser_InvalidEmail(t *testing.T) {
	user, err := NewUser("John", "Doe", "johndoe", "jhon@doe", "123456", "35050817013")
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag", err.Error())
}
