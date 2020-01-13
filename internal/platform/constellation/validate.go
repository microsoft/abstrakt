package constellation

import "gopkg.in/dealancer/validate.v2"

// CheckDuplicates checks for duplicate Relationship and Service IDs in a constellation file.
func (m *Config) CheckDuplicates() (duplicates []string) {
	IDs := []string{string(m.ID)}

	for _, i := range m.Services {
		_, exists := find(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	for _, i := range m.Relationships {
		_, exists := find(IDs, i.ID)
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
		_, exists := find(IDs, i.To)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.To)
		}

		_, exists = find(IDs, i.From)
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

// find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
