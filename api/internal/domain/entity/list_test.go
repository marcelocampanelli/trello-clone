package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_NewList(t *testing.T) {
	list, err := NewList("test")

	assert.Nil(t, err)
	assert.NotNil(t, list)
}

func TestList_NewList_Invalid(t *testing.T) {
	list, err := NewList("")

	assert.NotNil(t, err)
	assert.Nil(t, list)
}
