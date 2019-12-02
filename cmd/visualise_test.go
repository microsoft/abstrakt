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
	}

}

func TestParseYaml(t *testing.T) {

	testValidYAMLString := `
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
- Name: "Event Generator"
  Id: "9e1bcb3d-ff58-41d4-8779-f71e7b8800f8"
  Type: "EventGenerator"
  Properties: {}
- Name: "Azure Event Hub"
  Id: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  Type: "EventHub"
  Properties: {}
- Name: "Event Logger"
  Id: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Type: "EventLogger"
  Properties: {}
- Name: "Event Logger"
  Id: "1d0255d4-5b8c-4a52-b0bb-ac024cda37e5"
  Type: "EventLogger"
  Properties: {}
Relationships:
- Name: "Generator to Event Hubs Link"
  Id: "211a55bd-5d92-446c-8be8-190f8f0e623e"
  Description: "Event Generator to Event Hub connection"
  From: "9e1bcb3d-ff58-41d4-8779-f71e7b8800f8"
  To: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  Properties: {}
- Name: "Event Hubs to Event Logger Link"
  Id: "08ccbd67-456f-4349-854a-4e6959e5017b"
  Description: "Event Hubs to Event Logger connection"
  From: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  To: "1d0255d4-5b8c-4a52-b0bb-ac024cda37e5"
  Properties: {}
- Name: "Event Hubs to Event Logger Link Repeat"
  Id: "c8a719e0-164d-408f-9ed1-06e08dc5abbe"
  Description: "Event Hubs to Event Logger connection"
  From: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  To: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Properties: {}`

const test02ConstGraphString = `digraph Azure_Event_Hubs_Sample {
	Event_Generator->Azure_Event_Hub;
	Azure_Event_Hub->Event_Logger;
	Azure_Event_Hub->Event_Logger;
	Azure_Event_Hub;
	Event_Generator;
	Event_Logger;

}
`
