package chartservice

import (
	"helm.sh/helm/v3/pkg/chartutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestChartSavesAndLoads(t *testing.T) {
	tdir, err := ioutil.TempDir("", "helm-")
	tdir2, err := ioutil.TempDir("", "helm-")

	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)
	defer os.RemoveAll(tdir2)

	c, err := chartutil.Create("foo", tdir)

	if err != nil {
		t.Fatal(err)
	}

	mychart, err := LoadChartFromDir(c)
	if err != nil {
		t.Fatalf("Failed to load newly created chart %q: %s", c, err)
	}

	err = SaveChartToDir(mychart, tdir2)
	if err != nil {
		t.Fatalf("Failed to save newly created chart %q: %s", c, err)
	}

	newPath := filepath.Join(tdir2, "foo")

	_, err = LoadChartFromDir(newPath)

	if err != nil {
		t.Fatalf("Failed to load newly created chart %q: %s", newPath, err)
	}

}
