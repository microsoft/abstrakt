package constellation_test

import (
	"bytes"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/tools/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateGraph(t *testing.T) {

	retConfig := new(constellation.Config)
	err := retConfig.LoadFile("testdata/valid.yaml")

	assert.NoError(t, err)

	if err != nil {
		panic(err)
	}

	out := &bytes.Buffer{}

	cmpString := test02ConstGraphString
	retString, err := retConfig.GenerateGraph(out)

	assert.NoErrorf(t, err, "Should not receive error: %v", err)
	assert.True(t, helpers.CompareGraphOutputAsSets(cmpString, retString), "Input graph did not generate expected output graphviz representation\nExpected:\n%v \nGot:\n%v", cmpString, retString)
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
