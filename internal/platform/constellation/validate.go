package constellation

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"gopkg.in/dealancer/validate.v2"
)

// DuplicateIDs checks for duplicate Relationship and Service IDs in a constellation file.
func (m *Config) DuplicateIDs() (duplicates []string) {
	IDs := []string{string(m.ID)}

	for _, i := range m.Services {
		_, exists := helpers.Find(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	for _, i := range m.Relationships {
		_, exists := helpers.Find(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	return
}

// CheckServiceExists loops through each Relationship and checks if the services are declared.
func (m *Config) CheckServiceExists() (missing map[string][]string) {
	missing = make(map[string][]string)
	IDs := []string{}

	for _, i := range m.Services {
		IDs = append(IDs, i.ID)
	}

	for _, i := range m.Relationships {
		_, exists := helpers.Find(IDs, i.To)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.To)
		}

		_, exists = helpers.Find(IDs, i.From)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.From)
		}
	}

	return
}

// ValidateModel checks if constellation has all required felids
func (m *Config) ValidateModel() error {
	return validate.Validate(m)
}
