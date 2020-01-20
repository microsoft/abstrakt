package helpers_test

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExistsNoFile(t *testing.T) {
	assert.False(t, helpers.FileExists("does-not-exist"))
}

func TestFileExists(t *testing.T) {
	assert.True(t, helpers.FileExists("../../../README.md"))
}

func TestFileExistsFolder(t *testing.T) {
	assert.False(t, helpers.FileExists("../helpers"))
	assert.False(t, helpers.FileExists("../helpers/"))
}
