package dagconfigservice

import (
	"reflect"
	"testing"

	"github.com/microsoft/abstrakt/internal/tools/guid"
)

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
- Name: "Event Generator"
  Id: "9e1bcb3d-ff58-41d4-8779-f71e7b8800f8"
  Type: "EventGenerator"
  Properties: {}
- Name: "Azure Event Hub"
  Id: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  Type: "EventHub"
  Properties: {}
- Name: "Event Logger"
  Id: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Type: "EventLogger"
  Properties: {}
Relationships:
- Name: "Generator to Event Hubs Link"
  Id: "211a55bd-5d92-446c-8be8-190f8f0e623e"
  Description: "Event Generator to Event Hub connection"
  From: "e1bcb3d-ff58-41d4-8779-f71e7b8800f8"
  To: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  Properties: {}
- Name: "Event Hubs to Event Logger Link"
  Id: "08ccbd67-456f-4349-854a-4e6959e5017b"
  Description: "Event Hubs to Event Logger connection"
  From: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  To: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Properties: {}
`

var test01WantDag DagConfigService = DagConfigService{
	Name: "Azure Event Hubs Sample",
	ID:   guid.GUID("d6e4a5e9-696a-4626-ba7a-534d6ff450a5"),
	Services: []DagService{
		DagService{
			Name:       "Event Generator",
			ID:         guid.GUID("9e1bcb3d-ff58-41d4-8779-f71e7b8800f8"),
			Type:       "EventGenerator",
			Properties: make(map[string]DagProperty),
		},
		DagService{
			Name:       "Azure Event Hub",
			ID:         guid.GUID("3aa1e546-1ed5-4d67-a59c-be0d5905b490"),
			Type:       "EventHub",
			Properties: make(map[string]DagProperty),
		},
		DagService{
			Name:       "Event Logger",
			ID:         guid.GUID("a268fae5-2a82-4a3e-ada7-a52eeb7019ac"),
			Type:       "EventLogger",
			Properties: make(map[string]DagProperty),
		},
	},
	Relationships: []DagRelationship{
		DagRelationship{
			Name:        "Generator to Event Hubs Link",
			ID:          guid.GUID("211a55bd-5d92-446c-8be8-190f8f0e623e"),
			Description: "Event Generator to Event Hub connection",
			From:        guid.GUID("e1bcb3d-ff58-41d4-8779-f71e7b8800f8"),
			To:          guid.GUID("3aa1e546-1ed5-4d67-a59c-be0d5905b490"),
			Properties:  make(map[string]DagProperty),
		},
		DagRelationship{
			Name:        "Event Hubs to Event Logger Link",
			ID:          guid.GUID("08ccbd67-456f-4349-854a-4e6959e5017b"),
			Description: "Event Hubs to Event Logger connection",
			From:        guid.GUID("3aa1e546-1ed5-4d67-a59c-be0d5905b490"),
			To:          guid.GUID("a268fae5-2a82-4a3e-ada7-a52eeb7019ac"),
			Properties:  make(map[string]DagProperty),
		},
	},
}
