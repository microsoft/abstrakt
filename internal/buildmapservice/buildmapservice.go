package buildmapservice

//////////////////////////////////////////////////////
// BuildMapService:  Process map files relating services
// to helm chart files.  For example, see accompanying
// test file.
//////////////////////////////////////////////////////

import (
	"io/ioutil"
	"strings"

	"github.com/microsoft/abstrakt/internal/tools/guid"
	yamlParser "gopkg.in/yaml.v2"
)

// Note: the yaml mapping attributes are necessary (despite the nearly
// uniform 1-1 name correspondence). The yaml parser would otherwise
// expect the names in the YAML file to be all lower-case.
// e.g. ChartName would only work if "chartname" was used in the yaml file.

// BuildMapInfo -- info about an individual component
type BuildMapInfo struct {
	ChartName string `yaml:"ChartName"`
	Type      string `yaml:"Type"`
	Location  string `yaml:"Location"`
	Version   string `yaml:"Version"`
}

// BuildMapService -- data from the entire build map.
type BuildMapService struct {
	Name string         `yaml:"Name"`
	ID   guid.GUID      `yaml:"Id"`
	Maps []BuildMapInfo `yaml:"Maps"`
}

// FindByName -- Look up a map by chart name.
func (m *BuildMapService) FindByName(chartName string) (res *BuildMapInfo) {
	for _, wmi := range m.Maps {
		// try first for an exact match
		if chartName == wmi.ChartName {
			return &wmi
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc),
		if guid.TolerateMiscasedKey && strings.EqualFold(string(wmi.ChartName), chartName) {
			return &wmi
		}
	}
	return nil
}

// FindByType -- Look up a map by the "Type" value.
func (m *BuildMapService) FindByType(typeName string) (res *BuildMapInfo) {
	for _, wmi := range m.Maps {
		// try first for an exact match
		if typeName == wmi.Type {
			return &wmi
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc),
		if guid.TolerateMiscasedKey && strings.EqualFold(string(wmi.Type), typeName) {
			return &wmi
		}
	}
	return nil
}

// LoadMapFromFile -- New Map info instance from the named file.
func (m *BuildMapService) LoadMapFromFile(fileName string) (err error) {
	err = nil
	contentBytes, err := ioutil.ReadFile(fileName)
	if nil != err {
		return err
	}
	err = m.LoadMapFromString(string(contentBytes))
	return err
}

// LoadMapFromString -- New Map info instance from the given yaml string.
func (m *BuildMapService) LoadMapFromString(yamlString string) (err error) {
	err = nil

	err = yamlParser.Unmarshal([]byte(yamlString), m)

	return err
}
