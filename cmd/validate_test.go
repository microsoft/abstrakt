package cmd

import (
	"testing"
)

func TestValidateCommand(t *testing.T) {
	expected := "open does-not-exist: no such file or directory"

	constellationPath, _, _ := PrepareRealFilesForTest(t)
	output, err := executeCommand(newValidateCmd().cmd, "-f", constellationPath)

	if err != nil {
		t.Errorf("Did not received expected error. \nGot:\n %v", output)
	}

	_, err = executeCommand(newValidateCmd().cmd, "-f", "does-not-exist")

	if err != nil {
		checkStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not received expected error. \nExpected: %v\nGot:\n %v", expected, err.Error())
	}
}
