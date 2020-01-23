package cmd

import (
	helper "github.com/microsoft/abstrakt/internal/tools/test"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCommandNoArgs(t *testing.T) {
	_, err := helper.ExecuteCommand(newValidateCmd().cmd)
	assert.Error(t, err)
	assert.EqualError(t, err, "no flags were set")
}

func TestValidateCommandConstellationAndMapper(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"
	mapPath := "testdata/mapper/valid.yaml"

	output, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath, "-m", mapPath)
	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestValidateCommandConstellationExist(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"

	output, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)
	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestValidateCommandMapExist(t *testing.T) {
	mapPath := "testdata/mapper/valid.yaml"

	output, err := helper.ExecuteCommand(newValidateCmd().cmd, "-m", mapPath)
	assert.NoErrorf(t, err, "Did not received expected error. \nGot:\n %v", output)
}

func TestValidateCommandConstellationFail(t *testing.T) {
	expected := "Constellation: open does-not-exist: no such file or directory"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", "does-not-exist")

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
	expected := "Mapper: open does-not-exist: no such file or directory"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-m", "does-not-exist")

	entries := helper.GetAllLogs(hook.AllEntries())

	if err != nil {
		assert.Contains(t, entries, expected)
	} else {
		t.Errorf("Did not received expected error. \nExpected: %v\nGot:\n %v", expected, err.Error())
	}
}

func TestValidateCommandConstellationInvalidSchema(t *testing.T) {
	mapPath := "testdata/mapper/valid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", mapPath)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "Constellation: invalid schema")
	assert.EqualError(t, err, "Invalid configuration(s)")
}

func TestValidateCommandNapperInvalidSchema(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-m", constellationPath)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "Mapper: invalid schema")
	assert.EqualError(t, err, "Invalid configuration(s)")
}

func TestValidateDeploymentFail(t *testing.T) {
	constellationPath := "testdata/constellation/valid.yaml"
	mapPath := "testdata/mapper/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath, "-m", mapPath)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "Service `EventLogger` does not exist in map")
	assert.EqualError(t, err, "Invalid configuration(s)")
}

func TestValidateMapperDuplicates(t *testing.T) {
	mapPath := "testdata/mapper/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-m", mapPath)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "Duplicate `ChartName` present in config")
	assert.Contains(t, entries, "Duplicate `Type` present in config")
	assert.Contains(t, entries, "Duplicate `Location` present in config")
	assert.Contains(t, entries, "Mapper: invalid")
	assert.EqualError(t, err, "Invalid configuration(s)")
}

func TestValidateConstellationDuplicateIDs(t *testing.T) {
	constellationPath := "testdata/constellation/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "Duplicate `ID` present in config")
	assert.Contains(t, entries, "Constellation: invalid")
	assert.EqualError(t, err, "Invalid configuration(s)")
}

func TestValidateConstellationMissingServices(t *testing.T) {
	constellationPath := "testdata/constellation/invalid.yaml"

	hook := test.NewGlobal()
	_, err := helper.ExecuteCommand(newValidateCmd().cmd, "-f", constellationPath)

	entries := helper.GetAllLogs(hook.AllEntries())

	assert.Error(t, err)
	assert.Contains(t, entries, "Relationship 'Event Hubs to Event Logger Link' has missing `Services`:")
	assert.Contains(t, entries, "Constellation: invalid")
	assert.EqualError(t, err, "Invalid configuration(s)")
}
