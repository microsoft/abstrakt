package chartservice

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestChartSavesAndLoads(t *testing.T) {
	tdir, err := ioutil.TempDir("", "helm-")
	tdir2, err2 := ioutil.TempDir("", "helm-")

	if err != nil {
		t.Fatal(err)
	}

	if err2 != nil {
		t.Fatal(err2)
	}
	defer os.RemoveAll(tdir)
	defer os.RemoveAll(tdir2)

	c, err := CreateChart("foo", tdir)

	if err != nil {
		t.Fatal(err)
	}

	err = SaveChartToDir(c, tdir2)
	if err != nil {
		t.Fatalf("Failed to save newly created chart %q: %s", tdir2, err)
	}

	newPath := filepath.Join(tdir2, "foo")

	_, err = LoadChartFromDir(newPath)

	if err != nil {
		t.Fatalf("Failed to load newly created chart %q: %s", newPath, err)
	}

}
