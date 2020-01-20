package cmd

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCommandNoArgs(t *testing.T) {
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd)
	assert.Error(t, err)
	assert.EqualError(t, err, "no flags were set")
}

func TestValidateCommandConstellationAndMapper(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"
	mapPath := "testdata/mapper/valid.yaml"

	output, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath, "-m", mapPath)
	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestValidateCommandConstellationExist(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"

	output, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)
	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestValidateCommandMapExist(t *testing.T) {
	mapPath := "testdata/mapper/valid.yaml"

	output, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-m", mapPath)
	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestValidateCommandConstellationFail(t *testing.T) {
	expected := "constellation: open does-not-exist: no such file or directory"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", "does-not-exist")

	entries := []string{}

	for _, i := range hook.AllEntries() {
		entries = append(entries, i.Message)
	}

	if err != nil {
		assert.Contains(t, entries, expected)
	} else {
		t.Errorf("Did not received expected error. \nExpected: %v\nGot:\n %v", expected, err.Error())
	}
}

func TestValidateCommandMapFail(t *testing.T) {
	expected := "mapper: open does-not-exist: no such file or directory"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-m", "does-not-exist")

	entries := helpers.GetAllLogs(hook.AllEntries())

	if err != nil {
		assert.Contains(t, entries, expected)
	} else {
		t.Errorf("Did not received expected error. \nExpected: %v\nGot:\n %v", expected, err.Error())
	}
}

func TestValidateCommandConstellationInvalidSchema(t *testing.T) {
	mapPath := "testdata/mapper/valid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", mapPath)

	entries := helpers.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "constellation: invalid schema")
	assert.EqualError(t, err, "invalid configuration(s)")
}

func TestValidateCommandNapperInvalidSchema(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-m", constellationPath)

	entries := helpers.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "mapper: invalid schema")
	assert.EqualError(t, err, "invalid configuration(s)")
}

func TestValidateDeploymentFail(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"
	mapPath := "testdata/mapper/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath, "-m", mapPath)

	entries := helpers.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "service `EventLogger` does not exist in map")
	assert.EqualError(t, err, "invalid configuration(s)")
}

func TestValidateMapperDuplicates(t *testing.T) {
	mapPath := "testdata/mapper/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-m", mapPath)

	entries := helpers.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "duplicate `ChartName` present in config")
	assert.Contains(t, entries, "duplicate `Type` present in config")
	assert.Contains(t, entries, "duplicate `Location` present in config")
	assert.Contains(t, entries, "mapper: invalid")
	assert.EqualError(t, err, "invalid configuration(s)")
}

func TestValidateConstellationDuplicateIDs(t *testing.T) {
	constellationPath := "testdata/constellation/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)

	entries := helpers.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "duplicate `ID` present in config")
	assert.Contains(t, entries, "constellation: invalid")
	assert.EqualError(t, err, "invalid configuration(s)")
}

func TestValidateConstellationMissingServices(t *testing.T) {
	constellationPath := "testdata/constellation/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helpers.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)

	entries := helpers.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "relationship 'Event Hubs to Event Logger Link' has missing `Services`:")
	assert.Contains(t, entries, "constellation: invalid")
	assert.EqualError(t, err, "invalid configuration(s)")
}
