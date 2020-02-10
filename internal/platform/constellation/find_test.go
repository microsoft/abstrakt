package constellation_test

import (
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationshipFinding(t *testing.T) {
	dag := new(constellation.Config)
	_ = dag.LoadFile("testdata/valid.yaml")
	rel1 := dag.FindRelationshipByFromName("Event Generator")
	rel2 := dag.FindRelationshipByToName("Azure Event Hub")

	assert.Condition(t, func() bool { return !(rel1[0].From != rel2[0].From || rel1[0].To != rel2[0].To) }, "Relationships were not correctly resolved")
}

func TestMultipleInstanceInRelationships(t *testing.T) {
	newRelationship := constellation.Relationship{
		ID:          "Event Generator to Event Logger Link",
		Description: "Event Hubs to Event Logger connection",
		From:        "Event Generator",
		To:          "Event Logger",
	}

	dag := new(constellation.Config)
	_ = dag.LoadFile("testdata/valid.yaml")

	dag.Relationships = append(dag.Relationships, newRelationship)

	from := dag.FindRelationshipByFromName("Event Generator")
	to := dag.FindRelationshipByToName("Event Logger")

	assert.EqualValues(t, 2, len(from), "Event Generator did not have the correct number of `From` relationships")
	assert.EqualValues(t, 2, len(to), "Event Logger did not have the correct number of `To` relationships")
}

func TestForDuplicatIDsInServices(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/valid.yaml")
	assert.NoError(t, err)

	duplicates := testData.FindDuplicateIDs()
	assert.Nil(t, duplicates, "No duplicates should be found.")
}

func TestServicesExists(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/valid.yaml")
	assert.NoError(t, err)

	missing := testData.ServiceExists()
	assert.Empty(t, missing, "No missing services should be found.")
}

func TestForDuplicatIDsInServicesFail(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/duplicate/servIds.yaml")
	assert.NoError(t, err)

	duplicates := testData.FindDuplicateIDs()

	assert.NotNilf(t, duplicates, "There should be %v duplicate IDs found", 2)
	assert.Equal(t, 2, len(duplicates))
}

func TestForDuplicatIDsInRelationshipsFail(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/duplicate/relIds.yaml")
	assert.NoError(t, err)

	duplicates := testData.FindDuplicateIDs()

	assert.NotNilf(t, duplicates, "There should be %v duplicate IDs found", 1)
	assert.Equal(t, 1, len(duplicates))
}

func TestForDuplicatIDsFail(t *testing.T) {
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/duplicate/servRelIds.yaml")
	assert.NoError(t, err)

	duplicates := testData.FindDuplicateIDs()

	assert.NotNilf(t, duplicates, "There should be %v duplicate IDs found", 1)
	assert.Equal(t, 1, len(duplicates))
}

func TestServicesExistsFail(t *testing.T) {
	expected := "Azure Event Hub"
	testData := new(constellation.Config)

	err := testData.LoadFile("testdata/missing/relServRefId.yaml")
	assert.NoError(t, err)

	missing := testData.ServiceExists()
	foundID := missing["Event Hubs to Event Logger Link"][0]

	assert.Equal(t, 1, len(missing), "There should be only 1 missing services found")
	assert.Equalf(t, expected, foundID, "Incorrect reference found\nExpected: %v \nGot: %v", expected, foundID)
}

func TestSchemaMissingDagName(t *testing.T) {
	testData := new(constellation.Config)
	err := testData.LoadFile("testdata/missing/name.yaml")
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
