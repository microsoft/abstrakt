package cmd

import (
	"os"
	"testing" // based on standard golang testing library
)

// TestMain does setup or teardown (tests run when m.Run() is called)
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// Test your code here
func TestPrintVersion(t *testing.T) {
	PrintVersion()
	t.Log("All good")
}
