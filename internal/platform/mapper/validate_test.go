package mapper_test

import (
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaFailMissingChartNameValue(t *testing.T) {
	expected := "Validation error in field \"ChartName\" of type \"string\" using validator \"empty=false\""
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/invalid/missingChartNameValue.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
	assert.EqualError(t, err, expected, "Model validation should return error")
}

func TestSchemaFailMissingVersionProperty(t *testing.T) {
	expected := "Validation error in field \"Version\" of type \"string\" using validator \"empty=false\""
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/invalid/missingVersionProperty.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
	assert.EqualError(t, err, expected, "Model validation should return error")
}

func TestSchemaFailMissingMap(t *testing.T) {
	expected := "Validation error in field \"Maps\" of type \"[]mapper.Info\" using validator \"empty=false\""
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/invalid/missingMap.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
	assert.EqualError(t, err, expected, "Model validation should return error")
}
