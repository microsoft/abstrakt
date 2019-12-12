package cmd

import (
	"bytes"
	logger "github.com/microsoft/abstrakt/internal/tools/logger"
	"os"
	"testing" // based on standard golang testing library https://golang.org/pkg/testing/
)

// TestMain does setup or teardown (tests run when m.Run() is called)
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Test your code here
func TestVersion(t *testing.T) {
	expected := "0.0.1"
	version := Version()

	if version != expected {
		t.Errorf("Did not find correct abstrakt version. Expected %v, got %v", expected, version)
	}
}

func TestVersionCmd(t *testing.T) {
	expected := "0.0.1"
	buf := bytes.Buffer{}

	logger.SetOutput(&buf)
	_, err := executeCommand(newVersionCmd().cmd)

	if err != nil {
		t.Error(err)
	} else {
		checkStringContains(t, buf.String(), expected)
	}
}
