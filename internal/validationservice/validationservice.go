package validationservice

import (
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	"gopkg.in/dealancer/validate.v2"
)

// Validator is a service
type Validator struct {
	Config *dagconfigservice.DagConfigService
}

// CheckDuplicates checks for duplicate Relationship and Service IDs in a constellation file.
func (dag *Validator) CheckDuplicates() (duplicates []string) {
	IDs := []string{string(dag.Config.ID)}

	for _, i := range dag.Config.Services {
		_, exists := Find(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	for _, i := range dag.Config.Relationships {
		_, exists := Find(IDs, i.ID)
		if exists {
			duplicates = append(duplicates, i.ID)
		} else {
			IDs = append(IDs, i.ID)
		}
	}

	return
}

// CheckServiceExists loops through each Relationship and checks if the services are declared.
func (dag *Validator) CheckServiceExists() (missing map[string][]string) {
	missing = make(map[string][]string)
	IDs := []string{}

	for _, i := range dag.Config.Services {
		IDs = append(IDs, i.ID)
	}

	for _, i := range dag.Config.Relationships {
		_, exists := Find(IDs, i.To)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.To)
		}

		_, exists = Find(IDs, i.From)
		if !exists {
			missing[i.ID] = append(missing[i.ID], i.From)
		}
	}

	return
}

// ValidateModel checks if constellation has all required felids
func (dag *Validator) ValidateModel() error {
	return validate.Validate(dag.Config)
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
