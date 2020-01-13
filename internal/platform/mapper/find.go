package mapper

import (
	"github.com/microsoft/abstrakt/internal/tools/guid"
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
