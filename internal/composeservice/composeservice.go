package composeservice

import "github.com/microsoft/abstrakt/internal/dagconfigservice"

type ComposeService struct {
	DagConfigService dagconfigservice.DagConfigService
}

func (m *ComposeService) LoadFromString(dagString string, mapString string) {
	m.DagConfigService.LoadDagConfigFromString(dagString)
}

func NewComposeService() ComposeService {
	s := ComposeService{}
	s.DagConfigService = dagconfigservice.NewDagConfigService()
	return s
}
