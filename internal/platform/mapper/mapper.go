package mapper

//////////////////////////////////////////////////////
// BuildMapService:  Process map files relating services
// to helm chart files.  For example, see accompanying
// test file.
//////////////////////////////////////////////////////

import (
	"github.com/microsoft/abstrakt/internal/tools/guid"
	yamlParser "gopkg.in/yaml.v2"
	"io/ioutil"
)

// Note: the yaml mapping attributes are necessary (despite the nearly
// uniform 1-1 name correspondence). The yaml parser would otherwise
// expect the names in the YAML file to be all lower-case.
// e.g. ChartName would only work if "chartname" was used in the yaml file.

// Info -- info about an individual component
type Info struct {
	ChartName string `yaml:"ChartName"`
	Type      string `yaml:"Type"`
	Location  string `yaml:"Location"`
	Version   string `yaml:"Version"`
}

// Config -- data from the entire build map.
type Config struct {
	Name string    `yaml:"Name"`
	ID   guid.GUID `yaml:"Id"`
	Maps []Info    `yaml:"Maps"`
}

// LoadFile -- New Map info instance from the named file.
func (m *Config) LoadFile(fileName string) (err error) {
	contentBytes, err := ioutil.ReadFile(fileName)
	if nil != err {
		return
	}
	return m.LoadString(string(contentBytes))
}

// LoadString -- New Map info instance from the given yaml string.
func (m *Config) LoadString(yamlString string) error {
	return yamlParser.Unmarshal([]byte(yamlString), m)
}
