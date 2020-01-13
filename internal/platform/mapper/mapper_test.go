package mapper_test

import (
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/microsoft/abstrakt/internal/tools/guid"
	"reflect"
	"testing"
)

func TestMapFromString(t *testing.T) {
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
			args:    args{yamlString: configMapTest01String},
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
  Location: "../../helm/basictest"
  Version: "1.0.0"
- ChartName: "event_hub_sample_event_logger"
  Type: "EventLogger"
  Location: "../../helm/basictest"
  Version: "1.0.0"
- ChartName: "event_hub_sample_event_hub"
  Type: "EventHub"
  Location: "../../helm/basictest"
  Version: "1.0.0"
`

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
			Location:  "../../helm/basictest",
			Version:   "1.0.0",
		},
		{
			ChartName: "event_hub_sample_event_hub",
			Type:      "EventHub",
			Location:  "../../helm/basictest",
			Version:   "1.0.0",
		},
	},
}
