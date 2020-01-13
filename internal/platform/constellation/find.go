package constellation

import (
	"github.com/microsoft/abstrakt/internal/tools/guid"
	"strings"
)

// FindService -- Find a Service by id.
func (m *Config) FindService(serviceID string) *Service {
	for _, val := range m.Services {
		if val.ID == serviceID {
			return &val
		}
		if guid.TolerateMiscasedKey && strings.EqualFold(val.ID, serviceID) {
			return &val
		}
	}
	return nil
}

// FindRelationship -- Find a Relationship by id.
func (m *Config) FindRelationship(relationshipID string) *Relationship {
	for _, val := range m.Relationships {
		if val.ID == relationshipID {
			return &val
		} else if guid.TolerateMiscasedKey && strings.EqualFold(val.ID, relationshipID) {
			return &val
		}
	}
	return nil
}

// FindRelationshipByToName -- Find a Relationship by the name that is the target of the rel.
func (m *Config) FindRelationshipByToName(relationshipToName string) (res []Relationship) {
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
func (m *Config) FindRelationshipByFromName(relationshipFromName string) (res []Relationship) {
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
