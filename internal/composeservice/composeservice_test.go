package composeservice

import (
	"testing"
)

func TestLoadFromString(t *testing.T) {
	comp := NewComposeService()
	comp.LoadFromString(Test01DagStr, "")

}

const Test01DagStr = `Name: "Azure Event Hubs Sample"
Id: "d6e4a5e9-696a-4626-ba7a-534d6ff450a5"
Services:
- Name: "Event Generator"
  Id: "9e1bcb3d-ff58-41d4-8779-f71e7b8800f8"
  Type: "EventGenerator"
  Properties: {}
- Name: "Azure Event Hub"
  Id: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  Type: "EventHub"
  Properties: {}
- Name: "Event Logger"
  Id: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Type: "EventLogger"
  Properties: {}
Relationships:
- Name: "Generator to Event Hubs Link"
  Id: "211a55bd-5d92-446c-8be8-190f8f0e623e"
  Description: "Event Generator to Event Hub connection"
  From: "e1bcb3d-ff58-41d4-8779-f71e7b8800f8"
  To: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  Properties: {}
- Name: "Event Hubs to Event Logger Link"
  Id: "08ccbd67-456f-4349-854a-4e6959e5017b"
  Description: "Event Hubs to Event Logger connection"
  From: "3aa1e546-1ed5-4d67-a59c-be0d5905b490"
  To: "a268fae5-2a82-4a3e-ada7-a52eeb7019ac"
  Properties: {}
`
