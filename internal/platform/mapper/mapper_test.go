package mapper_test

import (
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/microsoft/abstrakt/tools/guid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMapFromString(t *testing.T) {
	contentBytes, err := ioutil.ReadFile("testdata/mapper.yaml")
	if nil != err {
		t.Fatal(err)
	}

	type args struct {
		yamlString string
	}
	tests := []struct {
		name    string
		args    args
		wantRet *mapper.Config
		wantErr bool
	}{
		{
			name:    "Test.01",
			args:    args{yamlString: string(contentBytes)},
			wantRet: &buildMap01,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := &mapper.Config{}
			err := mapper.LoadString(tt.args.yamlString)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadMapFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Truef(t, reflect.DeepEqual(mapper, tt.wantRet), "LoadMapFromString() = %v, want %v", mapper, tt.wantRet)
		})
	}
}

func TestMapLoadFile(t *testing.T) {
	mapper := &mapper.Config{}

	err := mapper.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	assert.Truef(t, reflect.DeepEqual(&buildMap01, mapper), "Expected: %v\nGot: %v", &buildMap01, mapper)
}

func TestIsEmptyTrue(t *testing.T) {
	dag := &mapper.Config{}
	assert.True(t, dag.IsEmpty())
}

func TestIsEmptyFalse(t *testing.T) {
	dag := buildMap01
	assert.False(t, dag.IsEmpty())
}

func TestSchemaPass(t *testing.T) {
	testData := new(mapper.Config)

	err := testData.LoadFile("testdata/mapper.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.NoError(t, err)
}

var buildMap01 = mapper.Config{
	Name: "Basic Azure Event Hubs maps",
	ID:   guid.GUID("a5a7c413-a020-44a2-bd23-1941adb7ad58"),
	Maps: []mapper.Info{
		{
			ChartName: "event_hub_sample_event_generator",
			Type:      "EventGenerator",
			Location:  "../../helm/basictest",
			Version:   "1.0.0",
		},
		{
			ChartName: "event_hub_sample_event_logger",
			Type:      "EventLogger",
			Location:  "../../helm/basictest2",
			Version:   "1.0.0",
		},
		{
			ChartName: "event_hub_sample_event_hub",
			Type:      "EventHub",
			Location:  "../../helm/basictest3",
			Version:   "1.0.0",
		},
	},
}
