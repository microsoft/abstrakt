package constellation

////////////////////////////////////////////////////////////
// Constellation class - information for a deployment
// regarding the Services and the Relationships between
// them.
//
// Usual starting point would be to construct a
// constellation.Config instance from the corresponding
// yaml using either:
//    dcPointer := constellation.LoadFile(<filename>)
// or
//    dcPointer := constellation.LoadString(<yamlTextString>)
//
// Parsing failures are indicated by a nil return.
////////////////////////////////////////////////////////////

import (
	"io/ioutil"
	"reflect"

	"github.com/microsoft/abstrakt/tools/guid"
	"gopkg.in/dealancer/validate.v2"
	yamlParser "gopkg.in/yaml.v2"
)

// Note: the yaml mappings are necessary (despite the 1-1 name correspondence).
// The yaml parser would otherwise expect the names in the YAML file to be all
// lower-case.  e.g. ChartName would only work if "chartname" was used in the
// yaml file.

// Property - an individual property in the DAG.
// For now, these are just interfaces as the value types are not firmed up
// for individual properties.  As the entire set of properties becomes
// known, each should be promoted out of the Properties collection to
// the main struct -- handling presence/absence via using pointer members,
// so as to allow for nil value == absence.
type Property interface{}

// Service -- a DAG Service description
type Service struct {
	ID         string              `yaml:"Id" validate:"empty=false"`
	Type       string              `yaml:"Type" validate:"empty=false"`
	Properties map[string]Property `yaml:"Properties"`
}

// Relationship -- a relationship between Services
type Relationship struct {
	ID          string              `yaml:"Id" validate:"empty=false"`
	Description string              `yaml:"Description"`
	From        string              `yaml:"From" validate:"empty=false"`
	To          string              `yaml:"To" validate:"empty=false"`
	Properties  map[string]Property `yaml:"Properties"`
}

// Config -- The DAG config for a deployment
type Config struct {
	Name          string         `yaml:"Name" validate:"empty=false"`
	ID            guid.GUID      `yaml:"Id" validate:"empty=false"`
	Services      []Service      `yaml:"Services" validate:"empty=false"`
	Relationships []Relationship `yaml:"Relationships"`
}

// LoadFile -- New DAG info instance from the named file.
func (m *Config) LoadFile(fileName string) (err error) {
	contentBytes, err := ioutil.ReadFile(fileName)
	if nil != err {
		return
	}
	return m.LoadString(string(contentBytes))
}

// LoadString -- New DAG info instance from the given yaml string.
func (m *Config) LoadString(yamlString string) error {
	return yamlParser.Unmarshal([]byte(yamlString), m)
}

//IsEmpty checks if config is empty.
func (m *Config) IsEmpty() bool {
	return reflect.DeepEqual(Config{}, *m)
}

// ValidateModel checks if constellation has all required felids
func (m *Config) ValidateModel() error {
	return validate.Validate(m)
}
