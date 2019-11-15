package composeservice

import (
	"errors"
	"github.com/microsoft/abstrakt/internal/buildmapservice"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
)

//ComposeService takes maps and configs and builds out the helm chart
type ComposeService struct {
	DagConfigService dagconfigservice.DagConfigService
	BuildMapService  buildmapservice.BuildMapService
}

//Compose takes the loaded DAG and maps and builds the Helm values and requirements documents
func (m *ComposeService) Compose() error {
	if m.DagConfigService.Name == "" || m.BuildMapService.Name == "" {
		return errors.New("Please initialise with LoadFromFile or LoadFromString")
	}

	return nil
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
