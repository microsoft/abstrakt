//For examples on how to use the wire libarary
// see: https://github.com/google/wire/blob/master/_tutorial/README.md

//+build wireinject



package sl

import (
	"github.com/google/wire"
	"github.com/microsoft/abstrakt/internal/sampleservice"
)

func SampleService() sampleservice.SampleService {
	wire.Build(sampleservice.New)
	return sampleservice.SampleService{}
}
