package sampleservice

import "fmt"

// SampleService is cool Sample service stuff
type SampleService struct {
	Something string
}

// New does nothing Sample service
func New() SampleService {
	return SampleService{}
}

// Print just prints
func (s *SampleService) Print() {
	fmt.Print("\n*************** Output!\n\n")

}
