package compose_test

import (
	"github.com/microsoft/abstrakt/internal/compose"
	helper "github.com/microsoft/abstrakt/tools/test"
	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestComposeService(t *testing.T) {
	_, _, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	comp := new(compose.Composer)
	_, err := comp.Build("test", tdir)

	assert.Error(t, err, "Compose should fail if not yet loaded")

	dag := "testdata/constellation.yaml"
	mapper := "testdata/mapper.yaml"

	_ = comp.LoadFile(dag, mapper)

	h, err := comp.Build("test", tdir)

	assert.NoError(t, err, "Compose should have loaded")

	_ = chartutil.SaveDir(h, tdir)
	h, _ = loader.LoadDir(tdir)

	contentBytes, err := ioutil.ReadFile("testdata/values.yaml")
	assert.NoError(t, err)

	for _, raw := range h.Raw {
		if raw.Name == "test/values.yaml" {
			assert.Equal(t, string(contentBytes), string(raw.Data))
		}
	}
}

func TestHelmLibCompose(t *testing.T) {
	_, _, tdir := helper.PrepareRealFilesForTest(t)

	defer helper.CleanTempTestFiles(t, tdir)

	c, err := chartutil.Create("foo", tdir)
	assert.NoError(t, err)

	dir := filepath.Join(tdir, "foo")

	mychart, err := loader.LoadDir(c)
	assert.NoErrorf(t, err, "Failed to load newly created chart %q: %s", c, err)

	assert.Equalf(t, "foo", mychart.Name(), "Expected name to be 'foo', got %q", mychart.Name())

	for _, f := range []string{
		chartutil.ChartfileName,
		chartutil.DeploymentName,
		chartutil.HelpersName,
		chartutil.IgnorefileName,
		chartutil.NotesName,
		chartutil.ServiceAccountName,
		chartutil.ServiceName,
		chartutil.TemplatesDir,
		chartutil.TemplatesTestsDir,
		chartutil.TestConnectionName,
		chartutil.ValuesfileName,
	} {
		_, err := os.Stat(filepath.Join(dir, f))
		assert.NoError(t, err)
	}

	mychart.Values["Jordan"] = "testing123"

	deps := []*chart.Dependency{
		{Name: "alpine", Version: "0.1.0", Repository: "https://example.com/charts"},
		{Name: "mariner", Version: "4.3.2", Repository: "https://example.com/charts"},
	}

	t.Logf("Directory: %v", tdir)

	mychart.Metadata.Dependencies = deps

	_ = chartutil.SaveDir(mychart, filepath.Join(tdir, "anotheretst"))

}

func TestLoadFromString(t *testing.T) {
	comp := new(compose.Composer)

	dag := "testdata/constellation.yaml"
	mapper := "testdata/mapper.yaml"

	err := comp.LoadFile(dag, mapper)
	assert.NoErrorf(t, err, "Error: %v", err)

	err = comp.LoadFile("sfdsd", mapper)
	assert.Error(t, err, "Didn't get error when should")

	err = comp.LoadFile(dag, "sdfsdf")
	assert.Error(t, err, "Didn't get error when should")
}
