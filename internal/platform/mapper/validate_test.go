package mapper_test

import (
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaFailMissingChartNameValue(t *testing.T) {
	testData := new(mapper.Config)
	err := testData.LoadFile("testdata/invalid/missingChartNameValue.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
}

func TestSchemaFailMissingVersionProperty(t *testing.T) {
	testData := new(mapper.Config)
	err := testData.LoadFile("testdata/invalid/missingVersionProperty.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
}

func TestSchemaFailInvalidType(t *testing.T) {
	testData := new(mapper.Config)
	err := testData.LoadFile("testdata/invalid/invalidType.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
}
