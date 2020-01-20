package helpers_test

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindElementInSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f"}

	index, found := helpers.Find(slice, "d")

	assert.True(t, found)
	assert.NotEqual(t, -1, index)
	assert.Equal(t, 3, index)
}

func TestFindElementNotInSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f"}

	index, found := helpers.Find(slice, "g")

	assert.False(t, found)
	assert.Equal(t, -1, index)
}
