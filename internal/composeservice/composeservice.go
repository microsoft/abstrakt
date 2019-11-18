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

	deps := make([]*chart.Dependency, 0)

	values := newChart.Values

	for _, n := range m.DagConfigService.Services {
		service := m.BuildMapService.FindByType(n.Type)
		if service == nil {
			return nil, fmt.Errorf("Could not find service: %v", service)
		}

		count := serviceMap[service.Type]
		alias := service.ChartName
		if count > 0 {
			alias = fmt.Sprintf("%v%v", service.ChartName, count)
		}

		fmt.Printf(alias)

		serviceMap[service.Type]++

		dep := &chart.Dependency{
			Name: service.ChartName, Version: service.Version, Repository: service.Location,
		}

		if count > 0 {
			dep.Alias = alias
		}

		deps = append(deps, dep)

		valMap := make(map[string]interface{})
		values[alias] = &valMap

		valMap["name"] = alias
		valMap["type"] = service.Type

		relationships := make(map[string]interface{})
		valMap["relationships"] = &relationships
		m.DagConfigService.FindRelationshipByID()
	}
	newChart.Values = values
	newChart.Metadata.Dependencies = deps

	return newChart, nil

}

//LoadFromFile takes a string dag and map and loads them
func (m *ComposeService) LoadFromFile(dagFile string, mapFile string) {
	m.DagConfigService.LoadDagConfigFromFile(dagFile)
	m.BuildMapService.LoadMapFromFile(mapFile)
}

//LoadFromString takes a string dag and map and loads them
func (m *ComposeService) LoadFromString(dagString string, mapString string) (err error) {
	err = m.DagConfigService.LoadDagConfigFromString(dagString)

	if err != nil {
		return err
	}

	err = m.BuildMapService.LoadMapFromString(mapString)

	return err
}

//NewComposeService constructs a new compose service
func NewComposeService() ComposeService {
	s := ComposeService{}
	s.DagConfigService = dagconfigservice.NewDagConfigService()
	s.BuildMapService = buildmapservice.NewBuildMapService()
	return s
}
