package cmd

import (
	"runtime"
	"testing"

	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// TestDiffCmdWithAllRequirementsNoError - test diff command parameters
// use valid arguments so expect no failures
func TestDiffCmdWithAllRequirementsNoError(t *testing.T) {
	constellationPathOrg, constellationPathNew, _, _ := helper.PrepareTwoRealConstellationFilesForTest(t)

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newDiffCmd().cmd, "-o", constellationPathOrg, "-n", constellationPathNew)

	assert.NoError(t, err)
	assert.True(t, helper.CompareGraphOutputAsSets(testDiffComparisonOutputString, hook.LastEntry().Message))
}

// TestDffCmdFailYaml - test diff command parameters
// Test both required command line parameters (-o, -n) failing each in turn
func TestDffCmdFailYaml(t *testing.T) {
	expected := "Constellation config failed to load file \"constellationPathOrg\": open constellationPathOrg: no such file or directory"
	winExpacted := "Constellation config failed to load file \"constellationPathOrg\": open constellationPathOrg: The system cannot find the file specified."

	_, err := helper.ExecuteCommand(newDiffCmd().cmd, "-o", "constellationPathOrg", "-n", "constellationPathNew")

	assert.Error(t, err)

	if runtime.GOOS == "windows" {
		assert.EqualError(t, err, winExpacted)
	} else {
		assert.EqualError(t, err, expected)
	}

	expected = "Constellation config failed to load file \"constellationPathNew\": open constellationPathNew: no such file or directory"
	winExpacted = "Constellation config failed to load file \"constellationPathNew\": open constellationPathNew: The system cannot find the file specified."

	_, err = helper.ExecuteCommand(newDiffCmd().cmd, "-o", "../examples/constellation/sample_constellation.yaml", "-n", "constellationPathNew")

	assert.Error(t, err)

	if runtime.GOOS == "windows" {
		assert.EqualError(t, err, winExpacted)
	} else {
		assert.EqualError(t, err, expected)
	}
}

// TestDiffCmdFailNotYaml - test diff command parameters
// Test both required command line parameter files fail when provided with invalid input files (-o, -n) failing each in turn
func TestDiffCmdFailNotYaml(t *testing.T) {
	expected := "Constellation config failed to load file \"diff.go\": yaml: line 26: mapping values are not allowed in this context"

	_, err := helper.ExecuteCommand(newDiffCmd().cmd, "-o", "diff.go", "-n", "diff.go")

	assert.Error(t, err)
	assert.EqualError(t, err, expected)

	_, err = helper.ExecuteCommand(newDiffCmd().cmd, "-o", "../examples/constellation/sample_constellation.yaml", "-n", "diff.go")

	assert.Error(t, err)
	assert.EqualError(t, err, expected)
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
