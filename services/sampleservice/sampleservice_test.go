package sampleservice_test

import (
	sl "github.com/microsoft/abstrakt/internal/servicelocator"
	"testing"
)

func TestResolveService(t *testing.T) {
	s := sl.SampleService
	if s == nil {
		t.Fail()
	}
}
