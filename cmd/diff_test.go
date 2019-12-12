package cmd

import (
	// "fmt"
	set "github.com/deckarep/golang-set"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	"testing"
)

// TestGetComparisonSets - generate sets of common, added and removed services and relationships from  input YAML
// and compare to manually created sets with a known or expected outcome
func TestGetComparisonSets(t *testing.T) {

	dsGraphOrg := dagconfigservice.NewDagConfigService()
	err := dsGraphOrg.LoadDagConfigFromString(testOrgDagStr)
	if err != nil {
		t.Errorf("dagConfigService failed to load dag from test string %s", err)
	}

	dsGraphNew := dagconfigservice.NewDagConfigService()
	err = dsGraphNew.LoadDagConfigFromString(testNewDagStr)
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
