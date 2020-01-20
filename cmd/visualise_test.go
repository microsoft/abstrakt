package cmd

import (
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestVisualiseCmdWithAllRequirementsNoError(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newVisualiseCmd().cmd, "-f", constellationPath)

	if err != nil {
		t.Error("Did not receive output")
	} else {
		assert.Truef(t, helpers.CompareGraphOutputAsSets(validGraphString, hook.LastEntry().Message), "Expcted output and produced output do not match : expected %s produced %s", validGraphString, hook.LastEntry().Message)
	}
}

func TestVisualiseCmdFailYaml(t *testing.T) {
	expected := "Could not open YAML input file for reading"

	output, err := helpers.ExecuteCommand(newVisualiseCmd().cmd, "-f", "constellationPath")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail. Expected: %v \nGot: %v", expected, output)
	}
}

func TestVisualiseCmdFailNotYaml(t *testing.T) {
	expected := "dagConfigService failed to load file"

	output, err := helpers.ExecuteCommand(newVisualiseCmd().cmd, "-f", "visualise.go")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail. Expected: %v \nGot: %v", expected, output)
	}
}

func TestFileExists(t *testing.T) {

	_, _, tdir := helpers.PrepareRealFilesForTest(t)

	defer helpers.CleanTempTestFiles(t, tdir)

	//Setup variables and content for test
	testValidFilename := filepath.Join(tdir, "testVisualise.out")
	testInvalidFilename := filepath.Join(tdir, "nonexistant.out")
	testData := []byte("A file to test with")

	//Create a file to test against
	err := ioutil.WriteFile(testValidFilename, testData, 0644)
	assert.NoError(t, err, "Could not create output testing file, cannot proceed")

	//Test that a valid file (created above) can be seen
	var result bool = helpers.FileExists(testValidFilename) //Expecting true - file does exists
	assert.True(t, result, "Test file does exist but testFile returns that it does not")

	//Test that an invalid file (does not exist) is not seen
	result = helpers.FileExists(testInvalidFilename) //Expecting false - file does not exist
	assert.False(t, result, "Test file does not exist but testFile says it does")

	err = os.Remove(testValidFilename)
	if err != nil {
		panic(err)
	}

	result = helpers.FileExists(testValidFilename) //Expecting false - file has been removed
	assert.False(t, result, "Test file has been removed but fileExists is finding it")
}

func TestParseYaml(t *testing.T) {

	retConfig := new(constellation.Config)
	err := retConfig.LoadString(testValidYAMLString)
	if err != nil {
		panic(err)
	}

	if retConfig.Name != "Azure Event Hubs Sample" &&
		retConfig.ID != "d6e4a5e9-696a-4626-ba7a-534d6ff450a5" &&
		len(retConfig.Services) != 1 &&
		len(retConfig.Relationships) != 1 {
		t.Errorf("YAML did not parse correctly and it should have")
	}
}

const testValidYAMLString = `
Description: "Event Generator to Event Hub connection"
From: 9e1bcb3d-ff58-41d4-8779-f71e7b8800f8
Id: 211a55bd-5d92-446c-8be8-190f8f0e623e
Name: "Azure Event Hubs Sample"
Properties: {}
Relationships: 
  - 
    Name: "Generator to Event Hubs Link"
Services: 
  - 
    Name: "Event Generator"
To: 3aa1e546-1ed5-4d67-a59c-be0d5905b490
Type: EventGenerator
`

const validGraphString = `digraph Azure_Event_Hubs_Sample {
	rankdir=LR;
	"Event_Generator"->"Azure_Event_Hub";
	"Azure_Event_Hub"->"Event_Logger";
	"Azure_Event_Hub";
	"Event_Generator";
	"Event_Logger";

}
`
