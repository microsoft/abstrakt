package mapper

import (
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"gopkg.in/dealancer/validate.v2"
)

// DuplicateChartName checks for duplicate chart names in a mapper file.
func (m *Config) DuplicateChartName() (duplicates []string) {
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

// DuplicateType checks for duplicate types in a mapper file.
func (m *Config) DuplicateType() (duplicates []string) {
	types := []string{}

	for _, i := range m.Maps {
		_, exists := helpers.Find(types, i.Type)
		if exists {
			duplicates = append(duplicates, i.Type)
		} else {
			types = append(types, i.Type)
		}
	}

	return
}

// DuplicateLocation checks for duplicate location in a mapper file.
func (m *Config) DuplicateLocation() (duplicates []string) {
	location := []string{}

	for _, i := range m.Maps {
		_, exists := helpers.Find(location, i.Location)
		if exists {
			duplicates = append(duplicates, i.Location)
		} else {
			location = append(location, i.Location)
		}
	}

	return
}

// ValidateModel checks if mapper has all required felids
func (m *Config) ValidateModel() error {
	return validate.Validate(m)
}
