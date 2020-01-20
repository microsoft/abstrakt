package mapper_test

import (
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaPass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.NoError(t, err)
}

func TestDuplicateChartNamePass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	duplicate := testData.DuplicateChartName()
	assert.Nil(t, duplicate)
}

func TestDuplicateTypesPass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	duplicate := testData.DuplicateType()
	assert.Nil(t, duplicate)
}

func TestDuplicateLocationsPass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	duplicate := testData.DuplicateLocation()
	assert.Nil(t, duplicate)
}

func TestSchemaFailMissingChartNameValue(t *testing.T) {
	expected := "Validation error in field \"ChartName\" of type \"string\" using validator \"empty=false\""
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/missing/chartNameValue.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
	assert.EqualError(t, err, expected, "Model validation should return error")
}

func TestSchemaFailMissingVersionProperty(t *testing.T) {
	expected := "Validation error in field \"Version\" of type \"string\" using validator \"empty=false\""
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/missing/versionProperty.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
	assert.EqualError(t, err, expected, "Model validation should return error")
}

func TestSchemaFailMissingMap(t *testing.T) {
	expected := "Validation error in field \"Maps\" of type \"[]mapper.Info\" using validator \"empty=false\""
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/missing/map.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should return error")
	assert.EqualError(t, err, expected, "Model validation should return error")
}

func TestDuplicateChartName(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/duplicate/chartNames.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.NoError(t, err, "Model validation should not return error")

	duplicate := testData.DuplicateChartName()
	assert.NotNil(t, duplicate)
	assert.Equal(t, 2, len(duplicate))
	assert.Equal(t, "event_hub_sample_event_generator", duplicate[0])
}

func TestDuplicateTypes(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/duplicate/types.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.NoError(t, err, "Model validation should not return error")

	duplicate := testData.DuplicateType()
	assert.NotNil(t, duplicate)
	assert.Equal(t, 2, len(duplicate))
	assert.Equal(t, "EventHub", duplicate[0])
}

func TestDuplicateLocation(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/duplicate/locations.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.NoError(t, err, "Model validation should not return error")

	duplicate := testData.DuplicateLocation()
	assert.NotNil(t, duplicate)
	assert.Equal(t, 2, len(duplicate))
	assert.Equal(t, "../../helm/basictest", duplicate[0])
}
