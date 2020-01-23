package cmd

import (
	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDiffCmdWithAllRequirementsNoError - test diff command parameters
// use valid arguments so expect no failures
func TestDiffCmdWithAllRequirementsNoError(t *testing.T) {
	constellationPathOrg, constellationPathNew, _, _ := helper.PrepareTwoRealConstellationFilesForTest(t)

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newDiffCmd().cmd, "-o", constellationPathOrg, "-n", constellationPathNew)

	if err != nil {
		t.Error("Did not receive output")
	} else {
		assert.Truef(t, helper.CompareGraphOutputAsSets(testDiffComparisonOutputString, hook.LastEntry().Message), "Expcted output and produced output do not match : expected %s produced %s", testDiffComparisonOutputString, hook.LastEntry().Message)
		// Did use this initially but wont work with the strongs output from the graphviz library as the sequence of entries in the output can change
		// while the sequence may change the result is still valid and the same so am usinga  local comparison function to get around this problem
		// assert.Contains(t, hook.LastEntry().Message, testDiffComparisonOutputString)
	}
}

// TestDffCmdFailYaml - test diff command parameters
// Test both required command line parameters (-o, -n) failing each in turn
func TestDffCmdFailYaml(t *testing.T) {
	expected := "Could not open original YAML input file for reading constellationPathOrg"

	output, err := helper.ExecuteCommand(newDiffCmd().cmd, "-o", "constellationPathOrg", "-n", "constellationPathNew")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail and it should have. Expected: %v \nGot: %v", expected, output)
	}

	expected = "Could not open new YAML input file for reading constellationPathNew"

	output, err = helper.ExecuteCommand(newDiffCmd().cmd, "-o", "../examples/constellation/sample_constellation.yaml", "-n", "constellationPathNew")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail and it should have. Expected: %v \nGot: %v", expected, output)
	}

}

// TestDiffCmdFailNotYaml - test diff command parameters
// Test both required command line parameter files fail when provided with invalid input files (-o, -n) failing each in turn
func TestDiffCmdFailNotYaml(t *testing.T) {
	expected := "dagConfigService failed to load file"

	output, err := helper.ExecuteCommand(newDiffCmd().cmd, "-o", "diff.go", "-n", "diff.go")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail. Expected: %v \nGot: %v", expected, output)
	}

	output, err = helper.ExecuteCommand(newDiffCmd().cmd, "-o", "../examples/constellation/sample_constellation.yaml", "-n", "diff.go")

	if err != nil {
		assert.Contains(t, err.Error(), expected)
	} else {
		t.Errorf("Did not fail. Expected: %v \nGot: %v", expected, output)
	}
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
