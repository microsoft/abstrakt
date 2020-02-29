package cmd

import (
	"os"
	"testing"

	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// TestMain does setup or teardown (tests run when m.Run() is called)
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Test your code here
func TestVersion(t *testing.T) {
	expected := "edge"
	version := Version()
	assert.Equal(t, expected, version)
}

func TestVersionCmd(t *testing.T) {
	expected := "abstrakt version: edge"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newVersionCmd().cmd)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.NoError(t, err)
	assert.Contains(t, entries, expected)
}

func TestCommit(t *testing.T) {
	expected := "n/a"
	version := Commit()
	assert.Equal(t, expected, version)
}

func TestCommitCmd(t *testing.T) {
	expected := "abstrakt commit: n/a"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newVersionCmd().cmd)

	assert.NoError(t, err)
	assert.Contains(t, hook.LastEntry().Message, expected)
}
