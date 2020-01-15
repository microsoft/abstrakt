package cmd

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCommand(t *testing.T) {
	expected := "open does-not-exist: no such file or directory"

	constellationPath, _, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)

	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)

	_, err = helpers.ExecuteCommand(newValidateCmd().cmd, "-f", "does-not-exist")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not received expected error. \nExpected: %v\nGot:\n %v", expected, err.Error())
	}
}
