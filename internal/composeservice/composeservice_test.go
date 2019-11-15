package composeservice

import (
	"testing"
)

func TestCompose(t *testing.T) {
	comp := NewComposeService()
	err := comp.Compose()

	if err == nil {
		t.Errorf("Compose should fail if not yet loaded")
	}

	err = comp.LoadFromString(test01DagStr, configMapTest01String)
	if err != nil {
		t.Errorf("Compose should have loaded")
	}

}

func TestLoadFromString(t *testing.T) {
	comp := NewComposeService()
	err := comp.LoadFromString(test01DagStr, configMapTest01String)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = comp.LoadFromString("sfdsd", configMapTest01String)
	if err == nil {
		t.Errorf("Didn't get error when should")
	}

	err = comp.LoadFromString(test01DagStr, "sdfsdf")
	if err == nil {
		t.Errorf("Didn't get error when should")
	}
}

const test01DagStr = `Name: "Azure Event Hubs Sample"
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

const configMapTest01String = `
Name: "Basic Azure Event Hubs maps"
Id: "a5a7c413-a020-44a2-bd23-1941adb7ad58"
Maps:
- ChartName: "event_hub_sample_event_generator"
  Type: "EventGenerator"
  Location: "../../helm"
  Version: "1.0.0"
- ChartName: "event_hub_sample_event_logger"
  Type: "EventLogger"
  Location: "../../helm"
  Version: "1.0.0"
- ChartName: "event_hub_sample_event_hub"
  Type: "EventHub"
  Location: "../../helm"
  Version: "1.0.0"
`
