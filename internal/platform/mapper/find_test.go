package mapper_test

import (
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByName(t *testing.T) {
	mapper := new(mapper.Config)

	err := mapper.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	info := mapper.FindByName("event_hub_sample_event_generator")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_generator", info.ChartName)
	assert.Equal(t, "EventGenerator", info.Type)
	assert.Equal(t, "../../helm/basictest", info.Location)
	assert.Equal(t, "1.0.0", info.Version)

	info = mapper.FindByName("event_hub_sample_event_logger")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_logger", info.ChartName)
	assert.Equal(t, "EventLogger", info.Type)
	assert.Equal(t, "../../helm/basictest2", info.Location)
	assert.Equal(t, "1.0.0", info.Version)

	info = mapper.FindByName("event_hub_sample_event_hub")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_hub", info.ChartName)
	assert.Equal(t, "EventHub", info.Type)
	assert.Equal(t, "../../helm/basictest3", info.Location)
	assert.Equal(t, "1.0.0", info.Version)
}

func TestFindByType(t *testing.T) {
	mapper := new(mapper.Config)

	err := mapper.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	info := mapper.FindByType("EventGenerator")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_generator", info.ChartName)
	assert.Equal(t, "EventGenerator", info.Type)
	assert.Equal(t, "../../helm/basictest", info.Location)
	assert.Equal(t, "1.0.0", info.Version)

	info = mapper.FindByType("EventLogger")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_logger", info.ChartName)
	assert.Equal(t, "EventLogger", info.Type)
	assert.Equal(t, "../../helm/basictest2", info.Location)
	assert.Equal(t, "1.0.0", info.Version)

	info = mapper.FindByType("EventHub")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_hub", info.ChartName)
	assert.Equal(t, "EventHub", info.Type)
	assert.Equal(t, "../../helm/basictest3", info.Location)
	assert.Equal(t, "1.0.0", info.Version)
}

func TestDuplicateChartNamePass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	duplicate := testData.FindDuplicateChartName()
	assert.Nil(t, duplicate)
}

func TestDuplicateTypesPass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	duplicate := testData.FindDuplicateType()
	assert.Nil(t, duplicate)
}

func TestDuplicateLocationsPass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	duplicate := testData.FindDuplicateLocation()
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

	duplicate := testData.FindDuplicateChartName()
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

	duplicate := testData.FindDuplicateType()
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

	duplicate := testData.FindDuplicateLocation()
	assert.NotNil(t, duplicate)
	assert.Equal(t, 2, len(duplicate))
	assert.Equal(t, "../../helm/basictest", duplicate[0])
}
