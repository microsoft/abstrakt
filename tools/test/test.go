package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	set "github.com/deckarep/golang-set"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ExecuteCommand is used to run a cobra command with arguments and return its value and error
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

// PrepareRealFilesForTest creates a new temporary folder with new map and constellation files.
// Returns path to new temp folder
func PrepareRealFilesForTest(t *testing.T) (string, string, string) {
	tdir, err := ioutil.TempDir("./", "output-")
	if err != nil {
		t.Fatal(err)
	}

	cwd, err2 := os.Getwd()
	if err2 != nil {
		t.Fatal(err2)
	}

	exampleConstellation := "../examples/constellation/sample_constellation.yaml"
	exampleMap := "../examples/constellation/sample_constellation_maps.yaml"

	if runtime.GOOS == "windows" {
		exampleConstellation = strings.ReplaceAll(exampleConstellation, "/", "\\")
		exampleMap = strings.ReplaceAll(exampleMap, "/", "\\")
	}

	constellationPath := path.Join(cwd, exampleConstellation)
	mapsPath := path.Join(cwd, exampleMap)

	return constellationPath, mapsPath, tdir
}

// PrepareTwoRealConstellationFilesForTest - this use the two required for the diff command
func PrepareTwoRealConstellationFilesForTest(t *testing.T) (string, string, string, string) {
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

	exampleConstellation := "../examples/constellation/sample_constellation.yaml"
	exampleConstellationChanged := "../examples/constellation/sample_constellation_changed.yaml"
	exampleMap := "../examples/constellation/sample_constellation_maps.yaml"

	if runtime.GOOS == "windows" {
		exampleConstellation = strings.ReplaceAll(exampleConstellation, "/", "\\")
		exampleConstellationChanged = strings.ReplaceAll(exampleConstellationChanged, "/", "\\")
		exampleMap = strings.ReplaceAll(exampleMap, "/", "\\")
	}

	constellationPathOrg := path.Join(cwd, exampleConstellation)
	constellationPathNew := path.Join(cwd, exampleConstellationChanged)
	mapsPath := path.Join(cwd, exampleMap)

	return constellationPathOrg, constellationPathNew, mapsPath, tdir
}

// CleanTempTestFiles removes a specific folder and all its contents from disk
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

// GetAllLogs loops through logrus entries and returns messages as []string
func GetAllLogs(logs []*logrus.Entry) (entries []string) {
	for _, i := range logs {
		entries = append(entries, i.Message)
	}
	return
}
