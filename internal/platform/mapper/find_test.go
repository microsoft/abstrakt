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
	assert.Equal(t, "../../helm/basictest", info.Location)
	assert.Equal(t, "1.0.0", info.Version)

	info = mapper.FindByName("event_hub_sample_event_hub")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_hub", info.ChartName)
	assert.Equal(t, "EventHub", info.Type)
	assert.Equal(t, "../../helm/basictest", info.Location)
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
	assert.Equal(t, "../../helm/basictest", info.Location)
	assert.Equal(t, "1.0.0", info.Version)

	info = mapper.FindByType("EventHub")

	assert.NotNil(t, info)
	assert.Equal(t, "event_hub_sample_event_hub", info.ChartName)
	assert.Equal(t, "EventHub", info.Type)
	assert.Equal(t, "../../helm/basictest", info.Location)
	assert.Equal(t, "1.0.0", info.Version)
}
