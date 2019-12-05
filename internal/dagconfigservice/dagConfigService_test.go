package dagconfigservice

import (
	"github.com/microsoft/abstrakt/internal/tools/guid"
	"reflect"
	"testing"
)

func TestRelationshipFinding(t *testing.T) {
	dag := &DagConfigService{}
	_ = dag.LoadDagConfigFromString(test01DagStr)
	rel1 := dag.FindRelationshipByFromName("Event Generator")
	rel2 := dag.FindRelationshipByToName("Azure Event Hub")

	if rel1[0].From != rel2[0].From || rel1[0].To != rel2[0].To {
		t.Error("Relationships were not correctly resolved")
	}

}

func TestNewDagConfigFromString(t *testing.T) {
	type targs struct {
		yamlString string
	}
	tests := []struct {
		name    string
		args    targs
		wantRet *DagConfigService
		wantErr bool
	}{
		{ // TEST START
			name:    "Test.01",
			args:    targs{yamlString: test01DagStr},
			wantRet: &test01WantDag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dag := &DagConfigService{}
			err := dag.LoadDagConfigFromString(tt.args.yamlString)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadDagConfigFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(dag, tt.wantRet) {
				t.Errorf("LoadDagConfigFromString() =\n%#v,\nWant:\n%#v\n", dag, tt.wantRet)
			}
		})
	}
}

// Sample DAG file data
const test01DagStr = `Name: "Azure Event Hubs Sample"
Id: "d6e4a5e9-696a-4626-ba7a-534d6ff450a5"
Services:
- Id: "Event Generator"
  Type: "EventGenerator"
  Properties: {}
- Id: "Azure Event Hub"
  Type: "EventHub"
  Properties: {}
- Id: "Event Logger"
  Type: "EventLogger"
  Properties: {}
Relationships:
- Id: "Generator to Event Hubs Link"
  Description: "Event Generator to Event Hub connection"
  From: "Event Generator"
  To: "Azure Event Hub"
  Properties: {}
- Id: "Event Hubs to Event Logger Link"
  Description: "Event Hubs to Event Logger connection"
  From: "Azure Event Hub"
  To: "Event Logger"
  Properties: {}
`

var test01WantDag DagConfigService = DagConfigService{
	Name: "Azure Event Hubs Sample",
	ID:   guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
	Services: []DagService{
		{
			ID:         "Event Generator",
			Type:       "EventGenerator",
			Properties: make(map[string]DagProperty),
		},
		{
			ID:         "Azure Event Hub",
			Type:       "EventHub",
			Properties: make(map[string]DagProperty),
		},
		{
			ID:         "Event Logger",
			Type:       "EventLogger",
			Properties: make(map[string]DagProperty),
		},
	},
	Relationships: []DagRelationship{
		{
			ID:          "Generator to Event Hubs Link",
			Description: "Event Generator to Event Hub connection",
			From:        "Event Generator",
			To:          "Azure Event Hub",
			Properties:  make(map[string]DagProperty),
		},
		{
			ID:          "Event Hubs to Event Logger Link",
			Description: "Event Hubs to Event Logger connection",
			From:        "Azure Event Hub",
			To:          "Event Logger",
			Properties:  make(map[string]DagProperty),
		},
	},
}
