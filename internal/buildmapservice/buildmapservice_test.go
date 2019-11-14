package buildmapservice

import (
	"reflect"
	"testing"

	"github.com/microsoft/abstrakt/internal/tools/guid"
)

func TestMapFromString(t *testing.T) {
	type args struct {
		yamlString string
	}
	tests := []struct {
		name    string
		args    args
		wantRet *BuildMapService
		wantErr bool
	}{
		{
			name:    "Test.01",
			args:    args{yamlString: configMapTest01String},
			wantRet: &buildMap01,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := &BuildMapService{}
			err := mapper.LoadMapFromString(tt.args.yamlString)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadMapFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(mapper, tt.wantRet) {
				t.Errorf("LoadMapFromString() = %v, want %v", mapper, tt.wantRet)
			}
		})
	}
}

const configMapTest01String = `
Name: "Basic Azure Event Hubs maps"
Id: "a5a7c413-a020-44a2-bd23-1941adb7ad58"
Maps:
- ChartName: "event_hub_sample_event_generator"
  Type: "EventGenerator"
  Location: "../../helm"
  Version: "1.0.0"
- ChartName: "event_hub_sample_event_logger"
  Type: "EventLogger"
  Location: "../../helm"
  Version: "1.0.0"
- ChartName: "event_hub_sample_event_hub"
  Type: "EventHub"
  Location: "../../helm"
  Version: "1.0.0"
`

var buildMap01 = BuildMapService{
	Name: "Basic Azure Event Hubs maps",
	ID:   guid.GUID("a5a7c413-a020-44a2-bd23-1941adb7ad58"),
	Maps: []BuildMapInfo{
		BuildMapInfo{
			ChartName: "event_hub_sample_event_generator",
			Type:      "EventGenerator",
			Location:  "../../helm",
			Version:   "1.0.0",
		},
		BuildMapInfo{
			ChartName: "event_hub_sample_event_logger",
			Type:      "EventLogger",
			Location:  "../../helm",
			Version:   "1.0.0",
		},
		BuildMapInfo{
			ChartName: "event_hub_sample_event_hub",
			Type:      "EventHub",
			Location:  "../../helm",
			Version:   "1.0.0",
		},
	},
}
