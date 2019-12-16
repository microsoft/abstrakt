package composeservice

import (
	"errors"
	"fmt"

	"github.com/microsoft/abstrakt/internal/buildmapservice"
	"github.com/microsoft/abstrakt/internal/chartservice"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"

	"helm.sh/helm/v3/pkg/chart"
)

//ComposeService takes maps and configs and builds out the helm chart
type ComposeService struct {
	DagConfigService dagconfigservice.DagConfigService
	BuildMapService  buildmapservice.BuildMapService
}

//Compose takes the loaded DAG and maps and builds the Helm values and requirements documents
func (m *ComposeService) Compose(name string, dir string) (*chart.Chart, error) {
	if m.DagConfigService.Name == "" || m.BuildMapService.Name == "" {
		return nil, errors.New("Please initialise with LoadFromFile or LoadFromString")
	}

	newChart, err := chartservice.CreateChart(name, dir)

	if err != nil {
		return nil, err
	}

	serviceMap := make(map[string]int)
	aliasMap := make(map[string]string)
	deps := make([]*chart.Dependency, 0)

	values := newChart.Values

	for _, n := range m.DagConfigService.Services {
		service := m.BuildMapService.FindByType(n.Type)
		if service == nil {
			return nil, fmt.Errorf("Could not find service %v", service)
		}

		count := serviceMap[service.Type]
		alias := service.ChartName
		if count > 0 {
			alias = fmt.Sprintf("%v%v", alias, count)
		}

		serviceMap[service.Type]++

		dep := &chart.Dependency{
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
		toRels := m.DagConfigService.FindRelationshipByToName(n.ID)
		fromRels := m.DagConfigService.FindRelationshipByFromName(n.ID)

		for _, i := range toRels {
			toRelations := make(map[string]string)
			relationships["input"] = append(relationships["input"], &toRelations)

			//find the target service
			foundService := m.DagConfigService.FindService(i.From)

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
			foundService := m.DagConfigService.FindService(i.To)

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

//LoadFromFile takes a string dag and map and loads them
func (m *ComposeService) LoadFromFile(dagFile string, mapFile string) (err error) {
	err = m.DagConfigService.LoadDagConfigFromFile(dagFile)
	if err != nil {
		return err
	}
	err = m.BuildMapService.LoadMapFromFile(mapFile)
	return
}

//LoadFromString takes a string dag and map and loads them
func (m *ComposeService) LoadFromString(dagString string, mapString string) (err error) {
	err = m.DagConfigService.LoadDagConfigFromString(dagString)

	if err != nil {
		return
	}

	err = m.BuildMapService.LoadMapFromString(mapString)

	return
}

//NewComposeService constructs a new compose service
func NewComposeService() ComposeService {
	s := ComposeService{}
	s.DagConfigService = dagconfigservice.NewDagConfigService()
	s.BuildMapService = buildmapservice.NewBuildMapService()
	return s
}
