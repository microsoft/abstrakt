//+build wireinject

package sl

import (
	"github.com/google/wire"
	"github.com/microsoft/abstrakt/services/sampleservice"
)

func SampleService() sampleservice.SampleService {
	wire.Build(sampleservice.New)
	return sampleservice.SampleService{}
}
