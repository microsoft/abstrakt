package constellation

import (
	"strings"

	"github.com/microsoft/abstrakt/tools/find"
	"github.com/microsoft/abstrakt/tools/guid"
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

// FindDuplicateIDs checks for duplicate Relationship and Service IDs in a constellation file.
func (m *Config) FindDuplicateIDs() (duplicates []string) {
	IDs := []string{string(m.ID)}

	for _, i := range m.Services {
		_, exists := find.Slice(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	for _, i := range m.Relationships {
		_, exists := find.Slice(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	return
}

// ServiceExists loops through each Relationship and checks if the services are declared.
func (m *Config) ServiceExists() (missing map[string][]string) {
	missing = make(map[string][]string)
	IDs := []string{}

	for _, i := range m.Services {
		IDs = append(IDs, i.ID)
	}

	for _, i := range m.Relationships {
		_, exists := find.Slice(IDs, i.To)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.To)
		}

		_, exists = find.Slice(IDs, i.From)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.From)
		}
	}

	return
}
