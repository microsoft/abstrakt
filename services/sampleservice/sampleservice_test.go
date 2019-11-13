//so named because you cannot import the same module again in a circular fashion
package sampleservice_test

import (
	sl "github.com/microsoft/abstrakt/internal/servicelocator"
	"testing"
)

func TestResolveService(t *testing.T) {
	s := sl.SampleService()

	s.Print()

}
