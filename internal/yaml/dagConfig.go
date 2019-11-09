package yaml

////////////////////////////////////////////////////////////
// DagConfig class - information for a deployment regarding
// the Services and the Relationships between them.
//
// Usual starting point would be to construct a DatConfig
// instance from the corresponding yaml using either:
//    dcPointer := NewDagConfigFromFile(<filename>)
// or
//    dcPointer := NewDagConfigFromString(<yamlTextString>)
//
// Parsing failures are indicated by a nil return.
////////////////////////////////////////////////////////////

import (
	yamlParser "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

// Note: the yaml mappings are necessary (despite the 1-1 name correspondence).
// The yaml parser would otherwise expect the names in the YAML file to be all
// lower-case.  e.g. ChartName would only work if "chartname" was used in the
// yaml file.

// DagProperty - an individual property in the DAG.
// For now, these are just interfaces as the value types are not firmed up
// for individual properties.  As the entire set of properties becomes
// known, each should be promoted out of the Properties collection to
// the main struct -- handling presence/absence via using pointer members,
// so as to allow for nil value == absence.
type DagProperty interface{}

// DagService -- a DAG Service description
type DagService struct {
	Name       string                 `yaml:"Name"`
	ID         GUID                   `yaml:"Id"`
	Type       string                 `yaml:"Type"`
	Properties map[string]DagProperty `yaml:"Properties"`
}

// DagRelationship -- a relationship between Services
type DagRelationship struct {
	Name        string                 `yaml:"Name"`
	ID          GUID                   `yaml:"Id"`
	Description string                 `yaml:"Description"`
	From        GUID                   `yaml:"From"`
	To          GUID                   `yaml:"To"`
	Properties  map[string]DagProperty `yaml:"Properties"`
}

// DagConfig -- The DAG config for a deployment
type DagConfig struct {
	Name          string            `yaml:"Name"`
	ID            GUID              `yaml:"Id"`
	Services      []DagService      `yaml:"Services"`
	Relationships []DagRelationship `yaml:"Relationships"`
}

// FindServiceByName -- Find a Service by name.
func (m *DagConfig) FindServiceByName(serviceName string) (res *DagService) {
	for _, val := range m.Services {
		// try first for an exact match
		if val.Name == serviceName {
			return &val
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc) ...
		if tolerateMiscasedKey && strings.EqualFold(val.Name, serviceName) {
			return &val
		}
	}
	return nil
}

// FindServiceByID -- Find a Service by id.
func (m *DagConfig) FindServiceByID(serviceID GUID) (res *DagService) {
	sid := string(serviceID) // no-op conversion, but needed for strings.* functions
	for _, val := range m.Services {
		// try first for an exact match
		if val.ID == serviceID {
			return &val
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc),
		if tolerateMiscasedKey && strings.EqualFold(string(val.ID), sid) {
			return &val
		}
	}
	return nil
}

// FindRelationshipByName -- Find a Relationship by name.
func (m *DagConfig) FindRelationshipByName(relationshipName string) (res *DagRelationship) {
	for _, val := range m.Relationships {
		// try first for an exact match
		if val.Name == relationshipName {
			return &val
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc) ...
		if tolerateMiscasedKey && strings.EqualFold(val.Name, relationshipName) {
			return &val
		}
	}
	return nil
}

// FindRelationshipByID -- Find a Relationship by id.
func (m *DagConfig) FindRelationshipByID(relationshipID GUID) (res *DagService) {
	rid := string(relationshipID) // no-op conversion, but needed for strings.* functions
	for _, val := range m.Services {
		// try first for an exact match
		if val.ID == relationshipID {
			return &val
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc),
		if tolerateMiscasedKey && strings.EqualFold(string(val.ID), rid) {
			return &val
		}
	}
	return nil
}

// NewDagConfigFromFile -- New DAG info instance from the named file.
func NewDagConfigFromFile(fileName string) (ret *DagConfig, err error) {
	err = nil
	contentBytes, err := ioutil.ReadFile(fileName)
	if nil != err {
		return nil, err
	}

	return NewDagConfigFromString(string(contentBytes))
}

// NewDagConfigFromString -- New DAG info instance from the given yaml string.
func NewDagConfigFromString(yamlString string) (ret *DagConfig, err error) {
	err = nil
	tp := &DagConfig{}
	err = yamlParser.Unmarshal([]byte(yamlString), tp)
	if err != nil {
		tp = nil
	}
	return tp, err
}
