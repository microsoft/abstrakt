package yaml

//////////////////////////////////////////////////////
// WormholeMap:  Process map files relating services
// to helm chart files.  For example, see accompanying
// test file.
//////////////////////////////////////////////////////

import (
	yamlParser "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

// Note: the yaml mapping attributes are necessary (despite the nearly
// uniform 1-1 name correspondence). The yaml parser would otherwise
// expect the names in the YAML file to be all lower-case.
// e.g. ChartName would only work if "chartname" was used in the yaml file.

// WormholeMapInfo -- info about an individual component
type WormholeMapInfo struct {
	ChartName string `yaml:"ChartName"`
	Type      string `yaml:"Type"`
	Location  string `yaml:"Location"`
	Version   string `yaml:"Version"`
}

// WormholeMap -- data from the entire build map.
type WormholeMap struct {
	Name string            `yaml:"Name"`
	ID   GUID              `yaml:"Id"`
	Maps []WormholeMapInfo `yaml:"Maps"`
}

// FindByName -- Look up a map by chart name.
func (m *WormholeMap) FindByName(chartName string) (res *WormholeMapInfo) {
	for _, wmi := range m.Maps {
		// try first for an exact match
		if chartName == wmi.ChartName {
			return &wmi
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc),
		if tolerateMiscasedKey && strings.EqualFold(string(wmi.ChartName), chartName) {
			return &wmi
		}
	}
	return nil
}

// FindByType -- Look up a map by the "Type" value.
func (m *WormholeMap) FindByType(typeName string) (res *WormholeMapInfo) {
	for _, wmi := range m.Maps {
		// try first for an exact match
		if typeName == wmi.Type {
			return &wmi
		}
		// if we want to tolerate case being incorrect (e.g., ABC vs. abc),
		if tolerateMiscasedKey && strings.EqualFold(string(wmi.Type), typeName) {
			return &wmi
		}
	}
	return nil
}

// NewWormholeMapFromFile -- New Map info instance from the named file.
func NewWormholeMapFromFile(fileName string) (ret *WormholeMap, err error) {
	err = nil
	contentBytes, err := ioutil.ReadFile(fileName)
	if nil != err {
		return nil, err
	}

	return NewWormholeMapFromString(string(contentBytes))
}

// NewWormholeMapFromString -- New Map info instance from the given yaml string.
func NewWormholeMapFromString(yamlString string) (ret *WormholeMap, err error) {
	err = nil
	tp := &WormholeMap{}
	err = yamlParser.Unmarshal([]byte(yamlString), tp)
	if err != nil {
		tp = nil
	}
	return tp, err
}
