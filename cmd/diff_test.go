package cmd

import (
	"fmt"
	set "github.com/deckarep/golang-set"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/sirupsen/logrus/hooks/test"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

// TestDiffCmdWithAllRequirementsNoError - test diff command parameters
// use valid arguments so expect no failures
func TestDiffCmdWithAllRequirementsNoError(t *testing.T) {
	constellationPathOrg, constellationPathNew, _, _ := localPrepareRealFilesForTest(t)

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newDiffCmd().cmd, "-o", constellationPathOrg, "-n", constellationPathNew)

	// fmt.Println(hook.LastEntry().Message)
	// fmt.Println(testDiffComparisonOutputString)

	if err != nil {
		t.Error("Did not receive output")
	} else {
		if !helpers.CompareGraphOutputAsSets(testDiffComparisonOutputString, hook.LastEntry().Message) {
			t.Errorf("Expcted output and produced output do not match : expected %s produced %s", testDiffComparisonOutputString, hook.LastEntry().Message)
		}
		// Did use this initially but wont work with the strongs output from the graphviz library as the sequence of entries in the output can change
		// while the sequence may change the result is still valid and the same so am usinga  local comparison function to get around this problem
		// checkStringContains(t, hook.LastEntry().Message, testDiffComparisonOutputString)
	}
}

// localPrepareRealFilesForTest - global function assumes only a single input file, this use the two required for the diff command
func localPrepareRealFilesForTest(t *testing.T) (string, string, string, string) {
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

	constellationPathOrg := path.Join(cwd, "../sample/constellation/sample_constellation.yaml")
	constellationPathNew := path.Join(cwd, "../sample/constellation/sample_constellation_changed.yaml")
	mapsPath := path.Join(cwd, "../sample/constellation/sample_constellation_maps.yaml")

	return constellationPathOrg, constellationPathNew, mapsPath, tdir
}

// TestDffCmdFailYaml - test diff command parameters
// Test both required command line parameters (-o, -n) failing each in turn
func TestDffCmdFailYaml(t *testing.T) {
	expected := "Could not open original YAML input file for reading constellationPathOrg"

	output, err := helpers.ExecuteCommand(newDiffCmd().cmd, "-o", "constellationPathOrg", "-n", "constellationPathNew")

	if err != nil {
		helpers.CheckStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail and it should have. Expected: %v \nGot: %v", expected, output)
	}

	expected = "Could not open new YAML input file for reading constellationPathNew"

	output, err = helpers.ExecuteCommand(newDiffCmd().cmd, "-o", "../sample/constellation/sample_constellation.yaml", "-n", "constellationPathNew")

	if err != nil {
		helpers.CheckStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail and it should have. Expected: %v \nGot: %v", expected, output)
	}

}

// TestDiffCmdFailNotYaml - test diff command parameters
// Test both required command line parameter files fail when provided with invalid input files (-o, -n) failing each in turn
func TestDiffCmdFailNotYaml(t *testing.T) {
	expected := "dagConfigService failed to load file"

	output, err := helpers.ExecuteCommand(newDiffCmd().cmd, "-o", "diff.go", "-n", "diff.go")

	if err != nil {
		helpers.CheckStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail. Expected: %v \nGot: %v", expected, output)
	}

	output, err = helpers.ExecuteCommand(newDiffCmd().cmd, "-o", "../sample/constellation/sample_constellation.yaml", "-n", "diff.go")

	if err != nil {
		helpers.CheckStringContains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail. Expected: %v \nGot: %v", expected, output)
	}
}

// TestGetComparisonSets - generate sets of common, added and removed services and relationships from  input YAML
// and compare to manually created sets with a known or expected outcome
func TestGetComparisonSets(t *testing.T) {

	dsGraphOrg := new(constellation.Config)
	err := dsGraphOrg.LoadString(testOrgDagStr)
	if err != nil {
		t.Errorf("dagConfigService failed to load dag from test string %s", err)
	}

	dsGraphNew := new(constellation.Config)
	err = dsGraphNew.LoadString(testNewDagStr)
	if err != nil {
		t.Errorf("dagConfigService failed to load file %s", err)
	}

	//construct sets struct for loaded constellation
	loadedSets := &setsForComparison{}
	// will populate with expected/known outcomes
	knownSets := &setsForComparison{}

	populateComparisonSets(knownSets)

	// function being tested
	fillComparisonSets(dsGraphOrg, dsGraphNew, loadedSets)

	if !knownSets.setCommonSvcs.Equal(loadedSets.setCommonSvcs) {
		t.Errorf("Common services - did not match between expected result and input yaml")
	}

	if !knownSets.setCommonRels.Equal(loadedSets.setCommonRels) {
		t.Errorf("Common relationships - did not match between expected result and input yaml")
	}

	if !knownSets.setAddedSvcs.Equal(loadedSets.setAddedSvcs) {
		t.Errorf("Added services - did not match between expected result and input yaml")
	}

	if !knownSets.setAddedRels.Equal(loadedSets.setAddedRels) {
		t.Errorf("Added relationships - did not match between expected result and input yaml")
	}

	if !knownSets.setDelSvcs.Equal(loadedSets.setDelSvcs) {
		t.Errorf("Deleted services - did not match between expected result and input yaml")
	}

	if !knownSets.setDelRels.Equal(loadedSets.setDelRels) {
		t.Errorf("Deleted relationships - did not match between expected result and input yaml")
	}

}

// testGraphWithChanges - test diff comparison function
func TestGraphWithChanges(t *testing.T) {

	dsGraphOrg := new(constellation.Config)
	err := dsGraphOrg.LoadString(testOrgDagStr)
	if err != nil {
		t.Errorf("dagConfigService failed to load dag from test string %s", err)
	}

	dsGraphNew := new(constellation.Config)
	err = dsGraphNew.LoadString(testNewDagStr)
	if err != nil {
		t.Errorf("dagConfigService failed to load file %s", err)
	}

	//construct sets struct for loaded constellation
	loadedSets := &setsForComparison{}

	// function being tested
	fillComparisonSets(dsGraphOrg, dsGraphNew, loadedSets)

	resString := createGraphWithChanges(dsGraphNew, loadedSets)

	if !helpers.CompareGraphOutputAsSets(testDiffComparisonOutputString, resString) {
		t.Errorf("Resulting output does not match the reference comparison input \n RESULT \n%s EXPECTED \n%s", resString, testDiffComparisonOutputString)
	}
}

// Utility to populate comparison sets with expected/known result
func populateComparisonSets(target *setsForComparison) {
	target.setCommonSvcs = set.NewSet()
	target.setCommonSvcs.Add("3aa1e546-1ed5-4d67-a59c-be0d5905b490")
	target.setCommonSvcs.Add("1d0255d4-5b8c-4a52-b0bb-ac024cda37e5")
	target.setCommonSvcs.Add("a268fae5-2a82-4a3e-ada7-a52eeb7019ac")

	target.setCommonRels = set.NewSet()
	target.setCommonRels.Add("3aa1e546-1ed5-4d67-a59c-be0d5905b490" + "|" + "1d0255d4-5b8c-4a52-b0bb-ac024cda37e5")

	target.setAddedSvcs = set.NewSet()
	target.setAddedSvcs.Add("9f1bcb3d-ff58-41d4-8779-f71e7b8800f8")
	target.setAddedSvcs.Add("b268fae5-2a82-4a3e-ada7-a52eeb7019ac")

	target.setAddedRels = set.NewSet()
	target.setAddedRels.Add("9f1bcb3d-ff58-41d4-8779-f71e7b8800f8" + "|" + "3aa1e546-1ed5-4d67-a59c-be0d5905b490")
	target.setAddedRels.Add("1d0255d4-5b8c-4a52-b0bb-ac024cda37e5" + "|" + "a268fae5-2a82-4a3e-ada7-a52eeb7019ac")
	target.setAddedRels.Add("a268fae5-2a82-4a3e-ada7-a52eeb7019ac" + "|" + "b268fae5-2a82-4a3e-ada7-a52eeb7019ac")

	target.setDelSvcs = set.NewSet()
	target.setDelSvcs.Add("9e1bcb3d-ff58-41d4-8779-f71e7b8800f8")

	target.setDelRels = set.NewSet()
	target.setDelRels.Add("9e1bcb3d-ff58-41d4-8779-f71e7b8800f8" + "|" + "3aa1e546-1ed5-4d67-a59c-be0d5905b490")
	target.setDelRels.Add("3aa1e546-1ed5-4d67-a59c-be0d5905b490" + "|" + "a268fae5-2a82-4a3e-ada7-a52eeb7019ac")

}

// Sample DAG file data - original file
const testOrgDagStr = `
Name: "Azure Event Hubs Sample"
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

// An updated (new) constellation dag with known differences
const testNewDagStr = `
Name: "Azure Event Hubs Sample Changed"
Id: "d6e4a5e9-696a-4626-ba7a-534d6ff450a5"
Services:
- Name: "Event Generator"
  Id: "9f1bcb3d-ff58-41d4-8779-f71e7b8800f8"
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
- Name: "Event Logger Added"
  Id: "b268fae5-2a82-4a3e-ada7-a52eeb7019ac"
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
  From: "9f1bcb3d-ff58-41d4-8779-f71e7b8800f8"
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
  From: "1d0255d4-5b8c-4a52-b0bb-ac024cda37e5"
  To: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Properties: {}
- Name: "Event Hubs to Event Logger Link Added to the end Repeat"
  Id: "d8a719e0-164d-408f-9ed1-06e08dc5abbe"
  Description: "Event Hubs to Event Logger connection"
  From: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  To: "b268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Properties: {}
  `

const testDiffComparisonOutputString = `digraph Azure_Event_Hubs_Sample_Changed_diff {
	rankdir=LR;
	"9f1bcb3d-ff58-41d4-8779-f71e7b8800f8"->"3aa1e546-1ed5-4d67-a59c-be0d5905b490"[ color="#d8ffa8" ];
	"3aa1e546-1ed5-4d67-a59c-be0d5905b490"->"1d0255d4-5b8c-4a52-b0bb-ac024cda37e5";
	"1d0255d4-5b8c-4a52-b0bb-ac024cda37e5"->"a268fae5-2a82-4a3e-ada7-a52eeb7019ac"[ color="#d8ffa8" ];
	"a268fae5-2a82-4a3e-ada7-a52eeb7019ac"->"b268fae5-2a82-4a3e-ada7-a52eeb7019ac"[ color="#d8ffa8" ];
	"3aa1e546-1ed5-4d67-a59c-be0d5905b490"->"a268fae5-2a82-4a3e-ada7-a52eeb7019ac"[ color="#ff9494" ];
	"9e1bcb3d-ff58-41d4-8779-f71e7b8800f8"->"3aa1e546-1ed5-4d67-a59c-be0d5905b490"[ color="#ff9494" ];
	"1d0255d4-5b8c-4a52-b0bb-ac024cda37e5" [ label="EventLogger
1d0255d4-5b8c-4a52-b0bb-ac024cda37e5", shape=rectangle, style="rounded, filled" ];
	"3aa1e546-1ed5-4d67-a59c-be0d5905b490" [ label="EventHub
3aa1e546-1ed5-4d67-a59c-be0d5905b490", shape=rectangle, style="rounded, filled" ];
	"9e1bcb3d-ff58-41d4-8779-f71e7b8800f8" [ color="#ff9494", shape=rectangle, style="rounded, filled" ];
	"9f1bcb3d-ff58-41d4-8779-f71e7b8800f8" [ color="#d8ffa8", label="EventGenerator
9f1bcb3d-ff58-41d4-8779-f71e7b8800f8", shape=rectangle, style="rounded, filled" ];
	"a268fae5-2a82-4a3e-ada7-a52eeb7019ac" [ label="EventLogger
a268fae5-2a82-4a3e-ada7-a52eeb7019ac", shape=rectangle, style="rounded, filled" ];
	"b268fae5-2a82-4a3e-ada7-a52eeb7019ac" [ color="#d8ffa8", label="EventLogger
b268fae5-2a82-4a3e-ada7-a52eeb7019ac", shape=rectangle, style="rounded, filled" ];

}
`
