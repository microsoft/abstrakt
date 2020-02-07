package cmd

import (
	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestMain does setup or teardown (tests run when m.Run() is called)
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Test your code here
func TestVersion(t *testing.T) {
	expected := "0.0.1"
	version := Version()
	assert.Equal(t, expected, version)
}

func TestVersionCmd(t *testing.T) {
	expected := "0.0.1"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newVersionCmd().cmd)

	assert.NoError(t, err)
	assert.Contains(t, hook.LastEntry().Message, expected)
}
