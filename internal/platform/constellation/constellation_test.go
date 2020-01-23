package constellation_test

import (
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/tools/guid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNewDagConfigFromString(t *testing.T) {
	contentBytes, err := ioutil.ReadFile("testdata/valid.yaml")
	if nil != err {
		t.Fatal(err)
	}

	type targs struct {
		yamlString string
	}
	tests := []struct {
		name    string
		args    targs
		wantRet *constellation.Config
		wantErr bool
	}{
		{ // TEST START
			name:    "Test.01",
			args:    targs{yamlString: string(contentBytes)},
			wantRet: &test01WantDag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dag := &constellation.Config{}
			err := dag.LoadString(tt.args.yamlString)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDagConfigFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Truef(t, reflect.DeepEqual(dag, tt.wantRet), "LoadDagConfigFromString() =\n%#v,\nWant:\n%#v\n", dag, tt.wantRet)
		})
	}
}

func TestMapLoadFile(t *testing.T) {
	dag := &constellation.Config{}

	err := dag.LoadFile("testdata/valid.yaml")
	assert.NoError(t, err)

	assert.Truef(t, reflect.DeepEqual(&test01WantDag, dag), "Expected: %v\nGot: %v", &test01WantDag, dag)
}

func TestIsEmptyTrue(t *testing.T) {
	dag := &constellation.Config{}
	assert.True(t, dag.IsEmpty())
}

func TestIsEmptyFalse(t *testing.T) {
	dag := test01WantDag
	assert.False(t, dag.IsEmpty())
}

var test01WantDag constellation.Config = constellation.Config{
	Name: "Azure Event Hubs Sample",
	ID:   guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
	Services: []constellation.Service{
		{
			ID:         "Event Generator",
			Type:       "EventGenerator",
			Properties: make(map[string]constellation.Property),
		},
		{
			ID:         "Azure Event Hub",
			Type:       "EventHub",
			Properties: make(map[string]constellation.Property),
		},
		{
			ID:         "Event Logger",
			Type:       "EventLogger",
			Properties: make(map[string]constellation.Property),
		},
	},
	Relationships: []constellation.Relationship{
		{
			ID:          "Generator to Event Hubs Link",
			Description: "Event Generator to Event Hub connection",
			From:        "Event Generator",
			To:          "Azure Event Hub",
			Properties:  make(map[string]constellation.Property),
		},
		{
			ID:          "Event Hubs to Event Logger Link",
			Description: "Event Hubs to Event Logger connection",
			From:        "Azure Event Hub",
			To:          "Event Logger",
			Properties:  make(map[string]constellation.Property),
		},
	},
}
