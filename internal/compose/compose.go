package compose

import (
	"fmt"
	"github.com/microsoft/abstrakt/internal/platform/chart"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	helm "helm.sh/helm/v3/pkg/chart"
)

//Composer takes maps and configs and builds out the helm chart
type Composer struct {
	Constellation constellation.Config
	Mapper        mapper.Config
}

//Build takes the loaded DAG and maps and builds the Helm values and requirements documents
func (m *Composer) Build(name string, dir string) (*helm.Chart, error) {
	if m.Constellation.Name == "" || m.Mapper.Name == "" {
		return nil, fmt.Errorf("Please initialise with LoadFromFile or LoadFromString")
	}

	newChart, err := chart.Create(name, dir)

	if err != nil {
		return nil, err
	}

	serviceMap := make(map[string]int)
	aliasMap := make(map[string]string)
	deps := make([]*helm.Dependency, 0)

	values := newChart.Values

	for _, n := range m.Constellation.Services {
		service := m.Mapper.FindByType(n.Type)
		if service == nil {
			return nil, fmt.Errorf("Could not find service %v", service)
		}

		count := serviceMap[service.Type]
		alias := service.ChartName
		if count > 0 {
			alias = fmt.Sprintf("%v%v", alias, count)
		}

		serviceMap[service.Type]++

		dep := &helm.Dependency{
			Name: service.ChartName, Version: service.Version, Repository: service.Location,
		}

		if count > 0 {
			dep.Alias = alias
		}

		aliasMap[string(n.ID)] = alias

		deps = append(deps, dep)

		valMap := make(map[string]interface{})
		values[alias] = &valMap

		valMap["name"] = alias
		valMap["type"] = service.Type

		relationships := make(map[string][]interface{})
		valMap["relationships"] = &relationships
		toRels := m.Constellation.FindRelationshipByToName(n.ID)
		fromRels := m.Constellation.FindRelationshipByFromName(n.ID)

		for _, i := range toRels {
			toRelations := make(map[string]string)
			relationships["input"] = append(relationships["input"], &toRelations)

			//find the target service
			foundService := m.Constellation.FindService(i.From)

			if foundService == nil {
				return nil, fmt.Errorf("Service '%v' referenced in relationship '%v' not found", i.From, i.ID)
			}

			toRelations["service"] = string(i.ID)
			toRelations["type"] = foundService.Type

			//ensure this only runs once all the counting is done
			closure := func() {
				relAlias := aliasMap[string(foundService.ID)]
				toRelations["name"] = relAlias
			}
			defer closure()
		}

		for _, i := range fromRels {
			fromRelations := make(map[string]string)
			relationships["output"] = append(relationships["output"], &fromRelations)

			//find the target service
			foundService := m.Constellation.FindService(i.To)

			if foundService == nil {
				return nil, fmt.Errorf("Service '%v' referenced in relationship '%v' not found", i.To, i.ID)
			}

			fromRelations["service"] = string(i.ID)
			fromRelations["type"] = foundService.Type

			//ensure this only runs once all the counting is done
			closure := func() {
				relAlias := aliasMap[string(foundService.ID)]
				fromRelations["name"] = relAlias
			}
			defer closure()
		}
	}

	newChart.Values = values
	newChart.Metadata.Dependencies = deps

	return newChart, nil
}

//LoadFile takes a string dag and map and loads them
func (m *Composer) LoadFile(dagFile string, mapFile string) (err error) {
	err = m.Constellation.LoadFile(dagFile)
	if err != nil {
		return err
	}
	return m.Mapper.LoadFile(mapFile)
}

//LoadString takes a string dag and map and loads them
func (m *Composer) LoadString(dagString string, mapString string) (err error) {
	err = m.Constellation.LoadString(dagString)
	if err != nil {
		return
	}
	return m.Mapper.LoadString(mapString)
}
