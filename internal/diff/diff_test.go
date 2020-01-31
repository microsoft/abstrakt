package diff_test

import (
	set "github.com/deckarep/golang-set"
	"github.com/microsoft/abstrakt/internal/diff"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetComparisonSets - generate sets of common, added and removed services and relationships from  input YAML
// and compare to manually created sets with a known or expected outcome
func TestGetComparisonSets(t *testing.T) {

	dsGraphOrg := new(constellation.Config)
	err := dsGraphOrg.LoadFile("testdata/original.yaml")
	assert.NoErrorf(t, err, "dagConfigService failed to load dag from test string %s", err)

	dsGraphNew := new(constellation.Config)
	err = dsGraphNew.LoadFile("testdata/modified.yaml")
	assert.NoErrorf(t, err, "dagConfigService failed to load file %s", err)

	// will populate with expected/known outcomes
	knownSets := &diff.ComparisonSet{}

	populateComparisonSets(knownSets)

	// function being tested
	diffSet := diff.Compare{Original: dsGraphOrg, New: dsGraphNew}
	loadedSets := diffSet.FillComparisonSets()

	assert.True(t, knownSets.SetCommonSvcs.Equal(loadedSets.SetCommonSvcs), "Common services - did not match between expected result and input yaml")
	assert.True(t, knownSets.SetCommonRels.Equal(loadedSets.SetCommonRels), "Common relationships - did not match between expected result and input yaml")
	assert.True(t, knownSets.SetAddedSvcs.Equal(loadedSets.SetAddedSvcs), "Added services - did not match between expected result and input yaml")
	assert.True(t, knownSets.SetAddedRels.Equal(loadedSets.SetAddedRels), "Added relationships - did not match between expected result and input yaml")
	assert.True(t, knownSets.SetDelSvcs.Equal(loadedSets.SetDelSvcs), "Deleted services - did not match between expected result and input yaml")
	assert.True(t, knownSets.SetDelRels.Equal(loadedSets.SetDelRels), "Deleted relationships - did not match between expected result and input yaml")
}

// testGraphWithChanges - test diff comparison function
func TestGraphWithChanges(t *testing.T) {

	dsGraphOrg := new(constellation.Config)
	err := dsGraphOrg.LoadFile("testdata/original.yaml")
	if err != nil {
		t.Errorf("dagConfigService failed to load dag from test string %s", err)
	}

	dsGraphNew := new(constellation.Config)
	err = dsGraphNew.LoadFile("testdata/modified.yaml")
	assert.NoErrorf(t, err, "dagConfigService failed to load file %s", err)

	// function being tested
	diffSet := diff.Compare{Original: dsGraphOrg, New: dsGraphNew}
	loadedSets := diffSet.FillComparisonSets()

	resString, err := diff.CreateGraphWithChanges(dsGraphNew, &loadedSets)

	assert.NoError(t, err)
	assert.Truef(t, helper.CompareGraphOutputAsSets(testDiffComparisonOutputString, resString), "Resulting output does not match the reference comparison input \n RESULT \n%s EXPECTED \n%s", resString, testDiffComparisonOutputString)
}

// Utility to populate comparison sets with expected/known result
func populateComparisonSets(target *diff.ComparisonSet) {
	target.SetCommonSvcs = set.NewSet()
	target.SetCommonSvcs.Add("3aa1e546-1ed5-4d67-a59c-be0d5905b490")
	target.SetCommonSvcs.Add("1d0255d4-5b8c-4a52-b0bb-ac024cda37e5")
	target.SetCommonSvcs.Add("a268fae5-2a82-4a3e-ada7-a52eeb7019ac")

	target.SetCommonRels = set.NewSet()
	target.SetCommonRels.Add("3aa1e546-1ed5-4d67-a59c-be0d5905b490" + "|" + "1d0255d4-5b8c-4a52-b0bb-ac024cda37e5")

	target.SetAddedSvcs = set.NewSet()
	target.SetAddedSvcs.Add("9f1bcb3d-ff58-41d4-8779-f71e7b8800f8")
	target.SetAddedSvcs.Add("b268fae5-2a82-4a3e-ada7-a52eeb7019ac")

	target.SetAddedRels = set.NewSet()
	target.SetAddedRels.Add("9f1bcb3d-ff58-41d4-8779-f71e7b8800f8" + "|" + "3aa1e546-1ed5-4d67-a59c-be0d5905b490")
	target.SetAddedRels.Add("1d0255d4-5b8c-4a52-b0bb-ac024cda37e5" + "|" + "a268fae5-2a82-4a3e-ada7-a52eeb7019ac")
	target.SetAddedRels.Add("a268fae5-2a82-4a3e-ada7-a52eeb7019ac" + "|" + "b268fae5-2a82-4a3e-ada7-a52eeb7019ac")

	target.SetDelSvcs = set.NewSet()
	target.SetDelSvcs.Add("9e1bcb3d-ff58-41d4-8779-f71e7b8800f8")

	target.SetDelRels = set.NewSet()
	target.SetDelRels.Add("9e1bcb3d-ff58-41d4-8779-f71e7b8800f8" + "|" + "3aa1e546-1ed5-4d67-a59c-be0d5905b490")
	target.SetDelRels.Add("3aa1e546-1ed5-4d67-a59c-be0d5905b490" + "|" + "a268fae5-2a82-4a3e-ada7-a52eeb7019ac")
}

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
