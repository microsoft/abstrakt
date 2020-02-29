package constellation_test

import (
	"bytes"
	"testing"

	"github.com/microsoft/abstrakt/internal/platform/constellation"
	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestGenerateGraph(t *testing.T) {

	retConfig := new(constellation.Config)
	err := retConfig.LoadFile("testdata/valid.yaml")

	assert.NoError(t, err)

	if err != nil {
		assert.FailNow(t, err.Error())
	}

	out := &bytes.Buffer{}

	cmpString := test02ConstGraphString
	retString, err := retConfig.GenerateGraph(out)

	assert.NoErrorf(t, err, "Should not receive error: %v", err)
	assert.True(t, helper.CompareGraphOutputAsSets(cmpString, retString), "Input graph did not generate expected output graphviz representation\nExpected:\n%v \nGot:\n%v", cmpString, retString)
}

const test02ConstGraphString = `digraph Azure_Event_Hubs_Sample {
	rankdir=LR;
	"Event_Generator"->"Azure_Event_Hub";
	"Azure_Event_Hub"->"Event_Logger";
	"Azure_Event_Hub";
	"Event_Generator";
	"Event_Logger";

}
`
