package cmd

import (
	// "bufio"
	"fmt"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {

	tdir, err := ioutil.TempDir("", "helm-")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err = os.RemoveAll(tdir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	//Setup variables and content for test
	testValidFilename := filepath.Join(tdir, "testVisualise.out")
	testInvalidFilename := filepath.Join(tdir, "nonexistant.out")
	testData := []byte("A file to test with")

	//Create a file to test against
	err = ioutil.WriteFile(testValidFilename, testData, 0644)
	if err != nil {
		fmt.Println("Could not create output testing file, cannot proceed")
		t.Error(err)
	}

	//Test that a valid file (created above) can be seen
	var result bool = fileExists(testValidFilename) //Expecting true - file does exists
	if result == false {
		t.Errorf("Test file does exist but testFile returns that it does not")
	}

	//Test that an invalid file (does not exist) is not seen
	result = fileExists(testInvalidFilename) //Expecting false - file does not exist
	if result != false {
		t.Errorf("Test file does not exist but testFile says it does")
	}

	err = os.Remove(testValidFilename)
	if err != nil {
		panic(err)
	}

	result = fileExists(testValidFilename) //Expecting false - file has been removed
	if result == true {
		t.Errorf("Test file has been removed but fileExists is finding it")
	}

}

func TestGenerateGraph(t *testing.T) {

	retConfig := dagconfigservice.NewDagConfigService()
	err := retConfig.LoadDagConfigFromString(test01DagStr)
	if err != nil {
		panic(err)
	}

	cmpString := test02ConstGraphString
	retString := generateGraph(retConfig)
	if retString != cmpString {
		t.Errorf("Input graph did not generate expected output graphviz representation")
		t.Errorf("Expected:\n%v \nGot:\n%v", cmpString, retString)
	}

}

func TestParseYaml(t *testing.T) {

	retConfig := dagconfigservice.NewDagConfigService()
	err := retConfig.LoadDagConfigFromString(testValidYAMLString)
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

// Sample DAG file data
const test01DagStr = `Name: "Azure Event Hubs Sample"
Id: "d6e4a5e9-696a-4626-ba7a-534d6ff450a5"
Services:
- Id: "Event Generator"
  Type: "EventGenerator"
  Properties: {}
- Id: "Azure Event Hub"
  Type: "EventHub"
  Properties: {}
- Id: "Event Logger"
  Type: "EventLogger"
  Properties: {}
- Id: "Event Logger"
  Type: "EventLogger"
  Properties: {}
Relationships:
- Id: "Generator to Event Hubs Link"
  Description: "Event Generator to Event Hub connection"
  From: "Event Generator"
  To: "Azure Event Hub"
  Properties: {}
- Id: "Event Hubs to Event Logger Link"
  Description: "Event Hubs to Event Logger connection"
  From: "Azure Event Hub"
  To: "Event Logger"
  Properties: {}
- Id: "Event Hubs to Event Logger Link Repeat"
  Description: "Event Hubs to Event Logger connection"
  From: "Azure Event Hub"
  To: "Event Logger"
  Properties: {}`

const test02ConstGraphString = `digraph Azure_Event_Hubs_Sample {
	rankdir=LR;
	"Event_Generator"->"Azure_Event_Hub";
	"Azure_Event_Hub"->"Event_Logger";
	"Azure_Event_Hub"->"Event_Logger";
	"Azure_Event_Hub";
	"Event_Generator";
	"Event_Logger";

}
`
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
