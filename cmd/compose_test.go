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

func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func TestComposeCmdVerifyRequiredFlags(t *testing.T) {
	expected := "required flag(s) \"constellationFilePath\", \"mapsFilePath\", \"outputPath\" not set"
	output, err := executeCommand(composeCmd, "")
	if err != nil {
		checkStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Expecting error: \n %v\nGot:\n %v\n", expected, output)
	}

}

func TestComposeCmdWithValidFlags(t *testing.T) {

	output, err := executeCommand(composeCmd, "-f", "constellationFilePath", "-m", "mapsFilePath", "-o", "outputPath")
	if err != nil {
		t.Errorf("error: \n %v\noutput:\n %v\n", err, output)
	}

}

func TestComposeWithRealFiles(t *testing.T) {
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

	output, err := executeCommand(composeCmd, "-f", constellationPath, "-m", mapsPath, "-o", tdir)

	if err != nil {
		t.Errorf("error: \n %v\noutput:\n %v\n", err, output)
	}

}