package mapper

import (
	"github.com/microsoft/abstrakt/tools/find"
	"github.com/microsoft/abstrakt/tools/guid"
	"strings"
)

// FindByName -- Look up a map by chart name.
func (m *Config) FindByName(chartName string) *Info {
	for _, wmi := range m.Maps {
		if chartName == wmi.ChartName {
			return &wmi
		}
		if guid.TolerateMiscasedKey && strings.EqualFold(string(wmi.ChartName), chartName) {
			return &wmi
		}
	}
	return nil
}

// FindByType -- Look up a map by the "Type" value.
func (m *Config) FindByType(typeName string) *Info {
	for _, wmi := range m.Maps {
		if typeName == wmi.Type {
			return &wmi
		}
		if guid.TolerateMiscasedKey && strings.EqualFold(string(wmi.Type), typeName) {
			return &wmi
		}
	}
	return nil
}

// FindDuplicateChartName checks for duplicate chart names in a mapper file.
func (m *Config) FindDuplicateChartName() (duplicates []string) {
	chartNames := []string{}

	for _, i := range m.Maps {
		_, exists := find.Slice(chartNames, i.ChartName)
		if exists {
			duplicates = append(duplicates, i.ChartName)
		} else {
			chartNames = append(chartNames, i.ChartName)
		}
	}

	return
}

// FindDuplicateType checks for duplicate types in a mapper file.
func (m *Config) FindDuplicateType() (duplicates []string) {
	types := []string{}

	for _, i := range m.Maps {
		_, exists := find.Slice(types, i.Type)
		if exists {
			duplicates = append(duplicates, i.Type)
		} else {
			types = append(types, i.Type)
		}
	}

	return
}

// FindDuplicateLocation checks for duplicate location in a mapper file.
func (m *Config) FindDuplicateLocation() (duplicates []string) {
	location := []string{}

	for _, i := range m.Maps {
		_, exists := find.Slice(location, i.Location)
		if exists {
			duplicates = append(duplicates, i.Location)
		} else {
			location = append(location, i.Location)
		}
	}

	return
}
