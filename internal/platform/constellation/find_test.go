package constellation_test

import (
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"testing"
)

func TestRelationshipFinding(t *testing.T) {
	dag := new(constellation.Config)
	_ = dag.LoadFile("testdata/valid.yaml")
	rel1 := dag.FindRelationshipByFromName("Event Generator")
	rel2 := dag.FindRelationshipByToName("Azure Event Hub")

	if rel1[0].From != rel2[0].From || rel1[0].To != rel2[0].To {
		t.Error("Relationships were not correctly resolved")
	}
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

	if len(from) != 2 {
		t.Error("Event Generator did not have the correct number of `From` relationships")
	}

	if len(to) != 2 {
		t.Error("Event Logger did not have the correct number of `To` relationships")
	}
}
