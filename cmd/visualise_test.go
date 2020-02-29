package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/tools/file"
	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestVisualiseCmdWithAllRequirementsNoError(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newVisualiseCmd().cmd, "-f", constellationPath)

	assert.NoError(t, err)
	assert.True(t, helper.CompareGraphOutputAsSets(validGraphString, hook.LastEntry().Message))
}

func TestVisualiseCmdFailYaml(t *testing.T) {
	expected := "Constellation config failed to load file"

	_, err := helper.ExecuteCommand(newVisualiseCmd().cmd, "-f", "constellationPath")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expected)
}

func TestVisualiseCmdFailNotYaml(t *testing.T) {
	expected := "Constellation config failed to load file"

	_, err := helper.ExecuteCommand(newVisualiseCmd().cmd, "-f", "visualise.go")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), expected)
}

func TestFileExists(t *testing.T) {
	_, _, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	//Setup variables and content for test
	testValidFilename := filepath.Join(tdir, "testVisualise.out")
	testInvalidFilename := filepath.Join(tdir, "nonexistant.out")
	testData := []byte("A file to test with")

	//Create a file to test against
	err := ioutil.WriteFile(testValidFilename, testData, 0644)
	assert.NoError(t, err, "Could not create output testing file, cannot proceed")

	//Test that a valid file (created above) can be seen
	var result bool = file.Exists(testValidFilename) //Expecting true - file does exists
	assert.True(t, result, "Test file does exist but testFile returns that it does not")

	//Test that an invalid file (does not exist) is not seen
	result = file.Exists(testInvalidFilename) //Expecting false - file does not exist
	assert.False(t, result, "Test file does not exist but testFile says it does")

	err = os.Remove(testValidFilename)
	assert.NoError(t, err)

	result = file.Exists(testValidFilename) //Expecting false - file has been removed
	assert.False(t, result, "Test file has been removed but fileExists is finding it")
}

func TestParseYaml(t *testing.T) {
	retConfig := new(constellation.Config)
	err := retConfig.LoadString(testValidYAMLString)
	assert.NoError(t, err)

	errMsg := "YAML did not parse correctly and it should have"

	assert.Equalf(t, retConfig.Name, "Azure Event Hubs Sample", errMsg)
	assert.EqualValuesf(t, retConfig.ID, "211a55bd-5d92-446c-8be8-190f8f0e623e", errMsg)
	assert.Equalf(t, len(retConfig.Services), 1, errMsg)
	assert.Equalf(t, len(retConfig.Relationships), 1, errMsg)
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
