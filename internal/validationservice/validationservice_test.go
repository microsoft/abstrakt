package validationservice

import (
	config "github.com/microsoft/abstrakt/internal/dagconfigservice"
	"testing"
)

func TestForDuplicatIDsInServices(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates != nil {
		t.Error("No duplicates should be found.")
	}
}

func TestForDuplicatIDsInServicesFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	for i := range testData.Services {
		testData.Services[i].ID = "Test"
	}

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates == nil {
		t.Errorf("There should be %v duplicate IDs found", len(testData.Services))
	}
}

func TestForDuplicatIDsInRelationships(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates != nil {
		t.Error("No duplicates should be found.")
	}
}

func TestForDuplicatIDsInRelationshipsFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	for i := range testData.Relationships {
		testData.Relationships[i].ID = "Test"
	}

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates == nil {
		t.Errorf("There should be %v duplicate IDs found", len(testData.Relationships))
	}
}

func TestForDuplicatIDs(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates != nil {
		t.Error("No duplicates should be found.")
	}
}

func TestForDuplicatIDsFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.Services[0].ID = "Test"
	testData.Relationships[0].ID = "Test"

	service := Validator{Config: testData}
	duplicates := service.CheckDuplicates()

	if duplicates == nil {
		t.Error("There should be 2 duplicate IDs found")
	}
}

func TestServicesExists(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	service := Validator{Config: testData}
	missing := service.CheckServiceExists()

	if len(missing) != 0 {
		t.Error("No missing services should be found.")
	}
}

func TestServicesExistsFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	for i := range testData.Services {
		testData.Services[i].ID = "Test"
	}

	service := Validator{Config: testData}
	missing := service.CheckServiceExists()

	if len(missing) == 0 {
		t.Errorf("There should be %v missing services found", len(testData.Services))
	}
}

func TestSchemaChecks(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err != nil {
		t.Errorf("Model validation should not return errors: %v", err.Error())
	}
}

func TestSchemaChecksFail(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.Name = ""

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingService(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.Services = nil

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingServiceID(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.Services[2].ID = ""

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingDagID(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.ID = ""

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingDagName(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.Name = ""

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

func TestSchemaMissingRelationshipID(t *testing.T) {
	testData := &config.DagConfigService{}
	_ = testData.LoadDagConfigFromString(test01DagStr)

	testData.Relationships[1].ID = ""

	service := Validator{Config: testData}
	err := service.ValidateModel()

	if err == nil {
		t.Error("Model validation should be invalid")
	}
}

const test01DagStr = `Name: "Azure Event Hubs Sample"
Id: "d6e4a5e9-696a-4626-ba7a-534d6ff450a5"
Services:
- Id: "Event Generator"
  Type: "EventGenerator"
  Properties: {}
- Id: "Azure Event Hub"
  Type: "EventHub"
  Properties: {}
- Id: "Event Logger"
  Type: "EventLogger"
  Properties: {}
Relationships:
- Id: "Generator to Event Hubs Link"
  Description: "Event Generator to Event Hub connection"
  From: "Event Generator"
  To: "Azure Event Hub"
  Properties: {}
- Id: "Event Hubs to Event Logger Link"
  Description: "Event Hubs to Event Logger connection"
  From: "Azure Event Hub"
  To: "Event Logger"
  Properties: {}
`
