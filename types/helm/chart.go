package helm;

import "text/template";
import "bytes"

// ChartYaml Go template for the Chart.yaml file
var ChartYaml string  = `
apiVersion: v2
name: {{ .Name }}
description: {{ .Description }}
type: application
version: {{ .Version }}
appVersion: {{ .Version }}
`

// Config Input to template for generating the main chart
type Config interface {
	Name() string; 
	Version() string; 
	Description() string;
}

// GenerateChart Outputs a string with go template filled
func GenerateChart(config Config) (*string, error) {
	template, err := template.New("GenerateChart").Parse(ChartYaml)

	if err != nil {
		return nil, err;
	}
	
	var buf bytes.Buffer;	
	
	err = template.Execute(&buf, config);	
	if err != nil {
		return nil, err;
	}

	var formatted *string =  new(string)
	*formatted = buf.String();
	return formatted, nil;
}
