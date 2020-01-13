package chart_test

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"github.com/microsoft/abstrakt/internal/platform/chart"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

var update = flag.Bool("update", false, "update golden dataset")

// TestUpdate Run: go test -update -run TestUpdate
func TestUpdate(t *testing.T) {
	if !*update {
		t.Skip("No golden data updated")
	}

	err := os.RemoveAll("testdata/golden")
	if err != nil {
		t.Fatal(err)
	}

	err = os.MkdirAll("testdata/golden/helm/charts", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	err = exec.Command("cp", "-r", "testdata/sample/helm", "testdata/golden/").Run()
	if err != nil {
		t.Fatal(err)
	}

	// Create zipped dependant charts
	err = exec.Command("tar", "cfz", "testdata/golden/helm/charts/event_hub_sample_event_generator-1.0.0.tgz", "testdata/sample/deps/event_hub_sample_event_generator/").Run()
	if err != nil {
		t.Fatal(err)
	}

	err = exec.Command("tar", "cfz", "testdata/golden/helm/charts/event_hub_sample_event_hub-1.0.0.tgz", "testdata/sample/deps/event_hub_sample_event_hub/").Run()
	if err != nil {
		t.Fatal(err)
	}

	err = exec.Command("tar", "cfz", "testdata/golden/helm/charts/event_hub_sample_event_logger-1.0.0.tgz", "testdata/sample/deps/event_hub_sample_event_logger/").Run()
	if err != nil {
		t.Fatal(err)
	}

	err = exec.Command("tar", "cfz", "testdata/golden/test-0.1.0.tgz", "testdata/sample/helm").Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChartSavesAndLoads(t *testing.T) {
	tdir, err := ioutil.TempDir("./", "output-")
	tdir2, err2 := ioutil.TempDir("./", "output-")

	if err != nil {
		t.Fatal(err)
	}

	if err2 != nil {
		t.Fatal(err2)
	}
	defer func() {
		err = os.RemoveAll(tdir)
		if err != nil {
			t.Error(err)
		}
		err = os.RemoveAll(tdir2)
		if err != nil {
			t.Error(err)
		}

	}()

	c, err := chart.Create("foo", tdir)

	if err != nil {
		t.Fatal(err)
	}

	err = chart.SaveToDir(c, tdir2)
	if err != nil {
		t.Fatalf("Failed to save newly created chart %q: %s", tdir2, err)
	}

	newPath := filepath.Join(tdir2, "foo")

	_, err = chart.LoadFromDir(newPath)

	if err != nil {
		t.Fatalf("Failed to load newly created chart %q: %s", newPath, err)
	}
}

func TestChartBuildChart(t *testing.T) {
	tdir, err := ioutil.TempDir("./", "output-")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err = os.RemoveAll(tdir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = exec.Command("cp", "-r", "testdata/sample/helm", tdir+"/").Run()
	if err != nil {
		t.Fatal(err)
	}

	err = exec.Command("cp", "-r", "testdata/sample/deps", tdir+"/").Run()
	if err != nil {
		t.Fatal(err)
	}

	_, err = chart.Build(tdir + "/helm")
	if err != nil {
		t.Fatalf("Failed to BuildChart(): %s", err)
	}

	chartsDir := tdir + "/helm/charts/"

	compareFiles(t, "testdata/golden/helm/charts/event_hub_sample_event_generator-1.0.0.tgz", chartsDir+"event_hub_sample_event_generator-1.0.0.tgz")
	compareFiles(t, "testdata/golden/helm/charts/event_hub_sample_event_hub-1.0.0.tgz", chartsDir+"event_hub_sample_event_hub-1.0.0.tgz")
	compareFiles(t, "testdata/golden/helm/charts/event_hub_sample_event_logger-1.0.0.tgz", chartsDir+"event_hub_sample_event_logger-1.0.0.tgz")
}

func TestZipChartToDir(t *testing.T) {
	tdir, err := ioutil.TempDir("./", "output-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.RemoveAll(tdir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	helm, err := chart.LoadFromDir("testdata/sample/helm")
	if err != nil {
		t.Fatalf("Failed on LoadChartFromDir(): %s", err)
	}

	_, err = chart.ZipToDir(helm, tdir)
	if err != nil {
		t.Fatalf("Failed on ZipChartToDir(): %s", err)
	}
	compareFiles(t, "testdata/golden/test-0.1.0.tgz", tdir+"/test-0.1.0.tgz")
}

func compareFiles(t *testing.T, expected, test string) {
	expectedHdrs := readTar(t, readGz(t, expected))
	testHdrs := readTar(t, readGz(t, test))

	for !reflect.DeepEqual(expectedHdrs, testHdrs) {
		t.Fatalf(`
		tars %[1]s and %[2]s were different
		%[1]s:
		%[3]v
		%[2]s:
		%[4]v
		`, expected, test, expectedHdrs, testHdrs)
	}
}

func readGz(t *testing.T, file string) (out bytes.Buffer) {
	f, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	zw, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = zw.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	writer := bufio.NewWriter(&out)
	_, err = io.Copy(writer, zw)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func readTar(t *testing.T, in bytes.Buffer) (out map[string]int64) {
	out = make(map[string]int64)
	tr := tar.NewReader(&in)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if hdr.Typeflag != tar.TypeDir {
			out[filepath.Base(hdr.Name)] = hdr.Size
		}
	}
	return
}
