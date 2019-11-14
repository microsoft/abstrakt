package yaml

import (
	"reflect"
	"testing"
)

func TestNewWormholeMapFromString(t *testing.T) {
	type args struct {
		yamlString string
	}
	tests := []struct {
		name    string
		args    args
		wantRet *WormholeMap
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
			gotRet, err := NewWormholeMapFromString(tt.args.yamlString)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWormholeMapFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("NewWormholeMapFromString() = %v, want %v", gotRet, tt.wantRet)
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

var buildMap01 = WormholeMap{
	Name: "Basic Azure Event Hubs maps",
	ID:   GUID("a5a7c413-a020-44a2-bd23-1941adb7ad58"),
	Maps: []WormholeMapInfo{
		WormholeMapInfo{
			ChartName: "event_hub_sample_event_generator",
			Type:      "EventGenerator",
			Location:  "../../helm",
			Version:   "1.0.0",
		},
		WormholeMapInfo{
			ChartName: "event_hub_sample_event_logger",
			Type:      "EventLogger",
			Location:  "../../helm",
			Version:   "1.0.0",
		},
		WormholeMapInfo{
			ChartName: "event_hub_sample_event_hub",
			Type:      "EventHub",
			Location:  "../../helm",
			Version:   "1.0.0",
		},
	},
}
