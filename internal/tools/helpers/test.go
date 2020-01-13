package helpers

import (
	"bytes"
	"fmt"
	set "github.com/deckarep/golang-set"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func ExecuteCommand(root *cobra.Command, args ...string) (output string, err error) {
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

func CheckStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func PrepareRealFilesForTest(t *testing.T) (string, string, string) {
	tdir, err := ioutil.TempDir("./", "output-")
	if err != nil {
		t.Fatal(err)
	}

	cwd, err2 := os.Getwd()
	if err2 != nil {
		t.Fatal(err2)
	}

	fmt.Print(cwd)

	constellationPath := path.Join(cwd, "../sample/constellation/sample_constellation.yaml")
	mapsPath := path.Join(cwd, "../sample/constellation/sample_constellation_maps.yaml")

	return constellationPath, mapsPath, tdir
}

func CleanTempTestFiles(t *testing.T, temp string) {
	err := os.RemoveAll(temp)
	if err != nil {
		t.Fatal(err)
	}
}

// CompareGraphOutputAsSets - the graphviz library does not always output the result string with nodes and edges
// in the same order (it can vary between calls). This does not impact using the result but makes testing the result a
// headache as the assumption is that the expected string and the produced string would match exactly. When the sequence
// changes they dont match. This function converts the strings into sets of lines and compares if the lines in the two outputs
// are the same
func CompareGraphOutputAsSets(expected, produced string) bool {
	lstExpected := strings.Split(expected, "\n")
	lstProduced := strings.Split(produced, "\n")

	setExpected := set.NewSet()
	setProduced := set.NewSet()

	for l := range lstExpected {
		setExpected.Add(l)
	}

	for l := range lstProduced {
		setProduced.Add(l)
	}

	return setProduced.Equal(setExpected)
}
