package cmd

import (
	"testing"

	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestComposeCommandReturnsErrorIfTemplateTypeIsInvalid(t *testing.T) {
	templateType := "ble"
	constellationPath, mapsPath, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	output, err := helper.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-if-template-type-is-invalid", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)
	assert.Errorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestComposeCommandDoesNotErrorIfTemplateTypeIsEmptyOrHelm(t *testing.T) {
	templateType := ""
	constellationPath, mapsPath, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	output, err := helper.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)
	assert.NoErrorf(t, err, "Did not expect error:\n %v\n output: %v", err, output)

	templateType = "helm"

	output, err = helper.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)
	assert.NoErrorf(t, err, "Did not expect error:\n %v\n output: %v", err, output)
}

func TestComposeCommandReturnsErrorWithInvalidFilePaths(t *testing.T) {
	output, err := helper.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-with-invalid-files", "-f", "invalid", "-m", "invalid", "-o", "invalid")
	assert.Errorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestComposeCmdVerifyRequiredFlags(t *testing.T) {
	expected := "required flag(s) \"constellationFilePath\", \"mapsFilePath\", \"outputPath\" not set"

	output, err := helper.ExecuteCommand(newComposeCmd().cmd, "")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Expecting error: \n %v\nGot:\n %v\n", expected, output)
	}
}

func TestComposeCmdWithValidFlags(t *testing.T) {
	constellationPath, mapsPath, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	output, err := helper.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-with-flags", "-f", constellationPath, "-m", mapsPath, "-o", tdir)
	assert.NoErrorf(t, err, "error: \n %v\noutput:\n %v\n", err, output)
}

func TestComposeWithRealFiles(t *testing.T) {
	constellationPath, mapsPath, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	output, err := helper.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-with-real-files", "-f", constellationPath, "-m", mapsPath, "-o", tdir)
	assert.NoErrorf(t, err, "error: \n %v\noutput:\n %v\n", err, output)
}
