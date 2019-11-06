package cmd

import (
	"testing"
)

func TestSome(t *testing.T) {
	PrintVersion()
	t.Log("All good")
	// t.Errorf("not good")
}
