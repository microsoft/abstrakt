// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package sl

import (
	"github.com/microsoft/abstrakt/internal/sampleservice"
)

// Injectors from sl.go:

func SampleService() sampleservice.SampleService {
	sampleService := sampleservice.New()
	return sampleService
}
