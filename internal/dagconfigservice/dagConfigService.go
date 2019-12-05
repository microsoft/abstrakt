package dagconfigservice

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
	"github.com/microsoft/abstrakt/internal/tools/guid"
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
	ID         string                 `yaml:"Id"`
	Type       string                 `yaml:"Type"`
	Properties map[string]DagProperty `yaml:"Properties"`
}

// DagRelationship -- a relationship between Services
type DagRelationship struct {
	ID          string                 `yaml:"Id"`
	Description string                 `yaml:"Description"`
	From        string                 `yaml:"From"`
	To          string                 `yaml:"To"`
	Properties  map[string]DagProperty `yaml:"Properties"`
}

// DagConfigService -- The DAG config for a deployment
type DagConfigService struct {
	Name          string            `yaml:"Name"`
	ID            guid.GUID         `yaml:"Id"`
	Services      []DagService      `yaml:"Services"`
	Relationships []DagRelationship `yaml:"Relationships"`
}

// NewDagConfigService -- Create a new DagConfigService instance
func NewDagConfigService() DagConfigService {
	return DagConfigService{}
}

// FindService -- Find a Service by id.
func (m *DagConfigService) FindService(serviceID string) (res *DagService) {
	for _, val := range m.Services {
		// try first for an exact match
		if val.ID == serviceID {
			return &val
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc) ...
		if guid.TolerateMiscasedKey && strings.EqualFold(val.ID, serviceID) {
			return &val
		}
	}
	return nil
}

// FindRelationship -- Find a Relationship by id.
func (m *DagConfigService) FindRelationship(relationshipID string) (res *DagRelationship) {
	for _, val := range m.Relationships {
		// try first for an exact match
		if val.ID == relationshipID {
			return &val
		} else if guid.TolerateMiscasedKey && strings.EqualFold(val.ID, relationshipID) {
			return &val
		}
	}
	return nil
}

// FindRelationshipByToName -- Find a Relationship by the name that is the target of the rel.
func (m *DagConfigService) FindRelationshipByToName(relationshipToName string) (res []DagRelationship) {
	for _, val := range m.Relationships {
		// try first for an exact match
		if val.To == relationshipToName {
			res = append(res, val)
		} else if guid.TolerateMiscasedKey && strings.EqualFold(string(val.To), relationshipToName) {
			res = append(res, val)
		}
	}
	return
}

// FindRelationshipByFromName -- Find a Relationship by the name that is the source of the rel.
func (m *DagConfigService) FindRelationshipByFromName(relationshipFromName string) (res []DagRelationship) {
	for _, val := range m.Relationships {
		// try first for an exact match
		if val.From == relationshipFromName {
			res = append(res, val)
		} else if guid.TolerateMiscasedKey && strings.EqualFold(string(val.From), relationshipFromName) {
			res = append(res, val)
		}
	}
	return
}

// LoadDagConfigFromFile -- New DAG info instance from the named file.
func (m *DagConfigService) LoadDagConfigFromFile(fileName string) (err error) {
	err = nil
	contentBytes, err := ioutil.ReadFile(fileName)
	if nil != err {
		return err
	}
	err = m.LoadDagConfigFromString(string(contentBytes))
	return err
}

// LoadDagConfigFromString -- New DAG info instance from the given yaml string.
func (m *DagConfigService) LoadDagConfigFromString(yamlString string) (err error) {
	err = nil
	//tp := &DagConfigService{}
	err = yamlParser.Unmarshal([]byte(yamlString), m)

	return err
}
