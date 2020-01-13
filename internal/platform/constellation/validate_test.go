package constellation_test

import (
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"testing"
)

func TestForDuplicatIDsInServices(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/valid.yaml")

	duplicates := testData.CheckDuplicates()

	if duplicates != nil {
		t.Error("No duplicates should be found.")
	}
}

func TestServicesExists(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/valid.yaml")

	missing := testData.CheckServiceExists()

	if len(missing) != 0 {
		t.Error("No missing services should be found.")
	}
}

func TestSchemaChecks(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/valid.yaml")

	err := testData.ValidateModel()

	if err != nil {
		t.Errorf("Model validation should not return errors: %v", err.Error())
	}
}

func TestForDuplicatIDsInServicesFail(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/duplicate/servIds.yaml")

	duplicates := testData.CheckDuplicates()

	if duplicates == nil {
		t.Errorf("There should be %v duplicate IDs found", len(testData.Services))
	}
}

func TestForDuplicatIDsInRelationshipsFail(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/duplicate/relIds.yaml")

	duplicates := testData.CheckDuplicates()

	if duplicates == nil {
		t.Errorf("There should be %v duplicate IDs found", len(testData.Relationships))
	}
}

func TestForDuplicatIDsFail(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/duplicate/servRelIds.yaml")

	duplicates := testData.CheckDuplicates()

	if duplicates == nil {
		t.Error("There should be 2 duplicate IDs found")
	}
}

func TestServicesExistsFail(t *testing.T) {
	expected := "Azure Event Hub"
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/missing/relServRefId.yaml")

	missing := testData.CheckServiceExists()

	foundID := missing["Event Hubs to Event Logger Link"][0]

	if len(missing) != 1 {
		t.Error("There should be only 1 missing services found")
	} else if foundID != expected {
		t.Errorf("Incorrect reference found. Expected: %v \nGot: %v", expected, foundID)
	}
}

func TestSchemaMissingDagName(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/missingName.yaml")

	err := testData.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingDagID(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/missing/id.yaml")

	err := testData.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingService(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/missing/serv.yaml")

	err := testData.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingServiceID(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/missing/servId.yaml")

	err := testData.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingRelationshipID(t *testing.T) {
	testData := new(constellation.Config)
	_ = testData.LoadFile("testdata/missing/relId.yaml")

	err := testData.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}
