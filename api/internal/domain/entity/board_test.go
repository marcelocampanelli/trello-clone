package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_NewBoardIsValid(t *testing.T) {
	board, err := NewBoard("Test", "Test", "131")

	assert.Nil(t, err)
	assert.NotNil(t, board)
}

func TestBoard_Modify(t *testing.T) {
	board, err := NewBoard("Test1", "Test1", "131")

	assert.Nil(t, err)
	assert.NotNil(t, board)

	err = board.Modify("Test", "Test")

	assert.Nil(t, err)

	assert.Equal(t, "Test", board.Name)
	assert.Equal(t, "Test", board.Description)
}

func TestBoard_NewBoardInvalid(t *testing.T) {
	board, err := NewBoard("", "", "")

	assert.NotNil(t, err)
	assert.Nil(t, board)
}
