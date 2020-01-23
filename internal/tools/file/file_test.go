package file_test

import (
	"github.com/microsoft/abstrakt/internal/tools/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExistsNoFile(t *testing.T) {
	assert.False(t, file.Exists("does-not-exist"))
}

func TestFileExists(t *testing.T) {
	assert.True(t, file.Exists("../../../README.md"))
}

func TestFileExistsIsFolder(t *testing.T) {
	assert.False(t, file.Exists("../file"))
	assert.False(t, file.Exists("../file/"))
}
