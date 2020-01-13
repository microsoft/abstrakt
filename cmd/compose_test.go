package cmd

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"testing"
)

func TestComposeCommandReturnsErrorIfTemplateTypeIsInvalid(t *testing.T) {
	templateType := "ble"
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-if-template-type-is-invalid", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)

	if err == nil {
		t.Errorf("Did not received expected error. \nGot:\n %v", output)
	}
}

func TestComposeCommandDoesNotErrorIfTemplateTypeIsEmptyOrHelm(t *testing.T) {
	templateType := ""
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)

	if err != nil {
		t.Errorf("Did not expect error:\n %v\n output: %v", err, output)
	}
	templateType = "helm"
	output, err = helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)

	if err != nil {
		t.Errorf("Did not expect error:\n %v\n output: %v", err, output)
	}
}

func TestComposeCommandReturnsErrorWithInvalidFilePaths(t *testing.T) {
	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-with-invalid-files", "-f", "invalid", "-m", "invalid", "-o", "invalid")

	if err == nil {
		t.Errorf("Did not received expected error. \nGot:\n %v", output)
	}
}

// TODO bug #43: figure out how to make this test work reliably.
// Something weird is making this test fail when run along with other tests in the package.
// It passes whenever it runs on it's own.
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
	if err != nil {
		t.Errorf("error: \n %v\noutput:\n %v\n", err, output)
	}
}

func TestComposeWithRealFiles(t *testing.T) {
	constellationPath, mapsPath, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	output, err := helpers.ExecuteCommand(newComposeCmd().cmd, "test-compose-cmd-with-real-files", "-f", constellationPath, "-m", mapsPath, "-o", tdir)
	if err != nil {
		t.Errorf("error: \n %v\noutput:\n %v\n", err, output)
	}

}
