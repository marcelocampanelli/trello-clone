package entity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard_NewCard(t *testing.T) {
	card, err := NewCard("test")

	assert.Nil(t, err)
	assert.NotNil(t, card)

	fmt.Println(card)
}

func TestCard_NewCard_Invalid(t *testing.T) {
	card, err := NewCard("")

	assert.NotNil(t, err)
	assert.Nil(t, card)
}
