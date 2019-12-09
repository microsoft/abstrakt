package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	return
}

func TestComposeCommandReturnsErrorIfTemplateTypeIsInvalid(t *testing.T) {
	templateType := "ble"
	constellationPath, mapsPath, tdir := PrepareRealFilesForTest(t)

	output, err := executeCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-if-template-type-is-invalid", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)

	if err == nil {
		t.Errorf("Did not received expected error. \nGot:\n %v", output)
	}
}

func TestComposeCommandDoesNotErrorIfTemplateTypeIsEmptyOrHelm(t *testing.T) {
	templateType := ""
	constellationPath, mapsPath, tdir := PrepareRealFilesForTest(t)

	output, err := executeCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)

	if err != nil {
		t.Errorf("Did not expect error:\n %v\n output: %v", err, output)
	}
	templateType = "helm"
	output, err = executeCommand(newComposeCmd().cmd, "test-compose-cmd-does-not-error-if-template-type-is-empty-or-helm", "-f", constellationPath, "-m", mapsPath, "-t", templateType, "-o", tdir)

	if err != nil {
		t.Errorf("Did not expect error:\n %v\n output: %v", err, output)
	}
}

func TestComposeCommandReturnsErrorWithInvalidFilePaths(t *testing.T) {
	output, err := executeCommand(newComposeCmd().cmd, "test-compose-cmd-returns-error-with-invalid-files", "-f", "invalid", "-m", "invalid", "-o", "invalid")

	if err == nil {
		t.Errorf("Did not received expected error. \nGot:\n %v", output)
	}
}

// TODO bug #43: figure out how to make this test work reliably.
// Something weird is making this test fail when run along with other tests in the package.
// It passes whenever it runs on it's own.
func TestComposeCmdVerifyRequiredFlags(t *testing.T) {
	expected := "required flag(s) \"constellationFilePath\", \"mapsFilePath\", \"outputPath\" not set"

	output, err := executeCommand(newComposeCmd().cmd, "")
	if err != nil {
		checkStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Expecting error: \n %v\nGot:\n %v\n", expected, output)
	}
}

func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func TestComposeCmdWithValidFlags(t *testing.T) {

	constellationPath, mapsPath, tdir := PrepareRealFilesForTest(t)

	output, err := executeCommand(newComposeCmd().cmd, "test-compose-cmd-with-flags", "-f", constellationPath, "-m", mapsPath, "-o", tdir)
	if err != nil {
		t.Errorf("error: \n %v\noutput:\n %v\n", err, output)
	}

}

func TestComposeWithRealFiles(t *testing.T) {
	constellationPath, mapsPath, tdir := PrepareRealFilesForTest(t)
	output, err := executeCommand(newComposeCmd().cmd, "test-compose-cmd-with-real-files", "-f", constellationPath, "-m", mapsPath, "-o", tdir)

	if err != nil {
		t.Errorf("error: \n %v\noutput:\n %v\n", err, output)
	}

}

func PrepareRealFilesForTest(t *testing.T) (string, string, string) {
	tdir, err := ioutil.TempDir("./", "output-")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err = os.RemoveAll(tdir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	cwd, err2 := os.Getwd()
	if err2 != nil {
		t.Fatal(err2)
	}

	fmt.Print(cwd)

	constellationPath := path.Join(cwd, "../sample/constellation/sample_constellation.yaml")
	mapsPath := path.Join(cwd, "../sample/constellation/sample_constellation_maps.yaml")

	return constellationPath, mapsPath, tdir
}
