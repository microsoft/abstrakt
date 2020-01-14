package cmd

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComposeCommandReturnsErrorIfTemplateTypeIsInvalid(t *testing.T) {
	templateType := "ble"
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-if-template-type-is-invalid", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)
	assert.Errorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestComposeCommandDoesNotErrorIfTemplateTypeIsEmptyOrHelm(t *testing.T) {
	templateType := ""
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)
	assert.NoErrorf(t, err, "Did not expect error:\n %v\n output: %v", err, output)

	templateType = "helm"

	output, err = helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)
	assert.NoErrorf(t, err, "Did not expect error:\n %v\n output: %v", err, output)
}

func TestComposeCommandReturnsErrorWithInvalidFilePaths(t *testing.T) {
	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-with-invalid-files", "-f", "invalid", "-m", "invalid", "-o", "invalid")
	assert.Errorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestComposeCmdVerifyRequiredFlags(t *testing.T) {
	expected := "required flag(s) \"constellationFilePath\", \"mapsFilePath\", \"outputPath\" not set"

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "")

	if err != nil {
		helpers.CheckStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Expecting error: \n %v\nGot:\n %v\n", expected, output)
	}
}

func TestComposeCmdWithValidFlags(t *testing.T) {
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-with-flags", "-f", constellationPath, "-m", mapsPath, "-o", tdir)
	assert.NoErrorf(t, err, "error: \n %v\noutput:\n %v\n", err, output)
}

func TestComposeWithRealFiles(t *testing.T) {
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-with-real-files", "-f", constellationPath, "-m", mapsPath, "-o", tdir)
	assert.NoErrorf(t, err, "error: \n %v\noutput:\n %v\n", err, output)
}
