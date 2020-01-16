package mapper

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"gopkg.in/dealancer/validate.v2"
)

// CheckDuplicates checks for duplicate chart names in a mapper file.
func (m *Config) CheckDuplicates() (duplicates []string) {
	chartNames := []string{}

	for _, i := range m.Maps {
		_, exists := helpers.Find(chartNames, i.ChartName)
		if exists {
			duplicates = append(duplicates, i.ChartName)
		} else {
			chartNames = append(chartNames, i.ChartName)
		}
	}

	return
}

// ValidateModel checks if mapper has all required felids
func (m *Config) ValidateModel() error {
	return validate.Validate(m)
}
