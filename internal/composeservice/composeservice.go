package composeservice

import (
	"errors"
	"fmt"

	"github.com/microsoft/abstrakt/internal/buildmapservice"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	yamlParser "gopkg.in/yaml.v2"
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

	yamlString := "---\nversion: 1\n..."

	yaml := make(map[interface{}]interface{})

	err := yamlParser.Unmarshal([]byte(yamlString), &yaml)

	yaml["jordan"] = "jordan"

	b, err := yamlParser.Marshal(&yaml)

	fmt.Println(string(b))

	// for i, n := range m.DagConfigService.Services {

	// }

	if err != nil {
		return err
	}

	return err
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
