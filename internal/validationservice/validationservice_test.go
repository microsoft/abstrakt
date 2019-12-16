package validationservice

import (
	config "github.com/microsoft/abstrakt/internal/dagconfigservice"
	"testing"
)

func TestForDuplicatIDsInServices(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/valid.yaml")

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates != nil {
		t.Error("No duplicates should be found.")
	}
}

func TestServicesExists(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/valid.yaml")

	service := Validator{Config: testData}
	missing := service.CheckServiceExists()

	if len(missing) != 0 {
		t.Error("No missing services should be found.")
	}
}

func TestSchemaChecks(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/valid.yaml")

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err != nil {
		t.Errorf("Model validation should not return errors: %v", err.Error())
	}
}

func TestForDuplicatIDsInServicesFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/duplicate/servIds.yaml")

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates == nil {
		t.Errorf("There should be %v duplicate IDs found", len(testData.Services))
	}
}

func TestForDuplicatIDsInRelationshipsFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/duplicate/relIds.yaml")

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates == nil {
		t.Errorf("There should be %v duplicate IDs found", len(testData.Relationships))
	}
}

func TestForDuplicatIDsFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/duplicate/servRelIds.yaml")

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates == nil {
		t.Error("There should be 2 duplicate IDs found")
	}
}

func TestServicesExistsFail(t *testing.T) {
	expected := "Azure Event Hub"
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/missing/relServRefId.yaml")

	service := Validator{Config: testData}
	missing := service.CheckServiceExists()

	foundID := missing["Event Hubs to Event Logger Link"][0]

	if len(missing) != 1 {
		t.Error("There should be only 1 missing services found")
	} else if foundID != expected {
		t.Errorf("Incorrect reference found. Expected: %v \nGot: %v", expected, foundID)
	}
}

func TestSchemaMissingDagName(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/missingName.yaml")

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingDagID(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/missing/id.yaml")

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingService(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/missing/serv.yaml")

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingServiceID(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/missing/servId.yaml")

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingRelationshipID(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromFile("testdata/missing/relId.yaml")

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}
