package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard_NewCard(t *testing.T) {
	card, err := NewCard("test", "01", "1", 1)

	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCard_NewCard_Invalid(t *testing.T) {
	card, err := NewCard("", "", "", 1)

	assert.NotNil(t, err)
	assert.Nil(t, card)
}

func TestCard_Modify(t *testing.T) {
	card, err := NewCard("test", "01", "1", 1)

	assert.Nil(t, err)
	assert.NotNil(t, card)

	err = card.Modify("test2", "02", "1", 1)

	assert.Nil(t, err)
	assert.Equal(t, "test2", card.Name)
	assert.Equal(t, "02", card.UserAssignedID)
}
