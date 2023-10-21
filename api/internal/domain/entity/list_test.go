package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_NewList(t *testing.T) {
	list, err := NewList("test", "1", 1)

	assert.Nil(t, err)
	assert.NotNil(t, list)
}

func TestList_NewList_Invalid(t *testing.T) {
	list, err := NewList("", "", 0)

	assert.NotNil(t, err)
	assert.Nil(t, list)
}
