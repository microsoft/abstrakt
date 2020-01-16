package constellation_test

import (
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForDuplicatIDsInServices(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/valid.yaml")
	assert.NoError(t, err)

	duplicates := testData.CheckDuplicates()
	assert.Nil(t, duplicates, "No duplicates should be found.")
}

func TestServicesExists(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/valid.yaml")
	assert.NoError(t, err)

	missing := testData.CheckServiceExists()
	assert.Empty(t, missing, "No missing services should be found.")
}

func TestSchemaChecks(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/valid.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.NoError(t, err, "Model validation should not return errors")
}

func TestForDuplicatIDsInServicesFail(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/duplicate/servIds.yaml")
	assert.NoError(t, err)

	duplicates := testData.CheckDuplicates()

	assert.NotNilf(t, duplicates, "There should be %v duplicate IDs found", 2)
	assert.Equal(t, 2, len(duplicates))
}

func TestForDuplicatIDsInRelationshipsFail(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/duplicate/relIds.yaml")
	assert.NoError(t, err)

	duplicates := testData.CheckDuplicates()

	assert.NotNilf(t, duplicates, "There should be %v duplicate IDs found", 1)
	assert.Equal(t, 1, len(duplicates))
}

func TestForDuplicatIDsFail(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/duplicate/servRelIds.yaml")
	assert.NoError(t, err)

	duplicates := testData.CheckDuplicates()

	assert.NotNilf(t, duplicates, "There should be %v duplicate IDs found", 1)
	assert.Equal(t, 1, len(duplicates))
}

func TestServicesExistsFail(t *testing.T) {
	expected := "Azure Event Hub"
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/missing/relServRefId.yaml")
	assert.NoError(t, err)

	missing := testData.CheckServiceExists()
	foundID := missing["Event Hubs to Event Logger Link"][0]

	assert.Equal(t, 1, len(missing), "There should be only 1 missing services found")
	assert.Equalf(t, expected, foundID, "Incorrect reference found\nExpected: %v \nGot: %v", expected, foundID)
}

func TestSchemaMissingDagName(t *testing.T) {
	testData := new(constellation.Config)
	err := testData.LoadFile("testdata/missingName.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should be invalid")
}

func TestSchemaMissingDagID(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/missing/id.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should be invalid")
}

func TestSchemaMissingService(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/missing/serv.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should be invalid")
}

func TestSchemaMissingServiceID(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/missing/servId.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should be invalid")
}

func TestSchemaMissingRelationshipID(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/missing/relId.yaml")
	assert.NoError(t, err)

	err = testData.ValidateModel()
	assert.Error(t, err, "Model validation should be invalid")
}
