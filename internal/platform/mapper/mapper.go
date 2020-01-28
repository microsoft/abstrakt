package mapper

//////////////////////////////////////////////////////
// BuildMapService:  Process map files relating services
// to helm chart files.  For example, see accompanying
// test file.
//////////////////////////////////////////////////////

import (
	"github.com/microsoft/abstrakt/tools/guid"
	"gopkg.in/dealancer/validate.v2"
	yamlParser "gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

// Note: the yaml mapping attributes are necessary (despite the nearly
// uniform 1-1 name correspondence). The yaml parser would otherwise
// expect the names in the YAML file to be all lower-case.
// e.g. ChartName would only work if "chartname" was used in the yaml file.

// Info -- info about an individual component
type Info struct {
	ChartName string `yaml:"ChartName" validate:"empty=false"`
	Type      string `yaml:"Type" validate:"empty=false"`
	Location  string `yaml:"Location" validate:"empty=false"`
	Version   string `yaml:"Version" validate:"empty=false"`
}

// Config -- data from the entire build map.
type Config struct {
	Name string    `yaml:"Name"`
	ID   guid.GUID `yaml:"Id"`
	Maps []Info    `yaml:"Maps" validate:"empty=false"`
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

//IsEmpty checks if config is empty.
func (m *Config) IsEmpty() bool {
	return reflect.DeepEqual(Config{}, *m)
}

// ValidateModel checks if mapper has all required felids
func (m *Config) ValidateModel() error {
	return validate.Validate(m)
}
