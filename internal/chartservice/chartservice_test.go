package chartservice

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"	
	"os/exec"
	"strings"
	"io"
	"crypto/md5"
	"encoding/hex"
)

func TestChartSavesAndLoads(t *testing.T) {
	tdir, err := ioutil.TempDir("./", "output-")
	tdir2, err2 := ioutil.TempDir("./", "output-")

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

func TestChartBuildChart(t *testing.T) {
	tdir, err := ioutil.TempDir("./", "output-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)

	err = exec.Command("cp", "-r", "../../sample/helm", tdir + "/").Run()
	if err != nil {
		t.Fatal(err)
	}

	err = BuildChart(tdir + "/helm");
	if err != nil {
		t.Fatalf("Failed to BuildChart(): %s", err)
	}	
	
	chartsDir := tdir + "/helm/charts/";

	checkMd5(t, "1a59f6425dddb08a44a6e959ce324593", chartsDir + "event_hub_sample_event_generator-1.0.0.tgz");
	checkMd5(t, "06380fd94eee71844f32843dc5f723be", chartsDir + "event_hub_sample_event_hub-1.0.0.tgz");
	checkMd5(t, "30768e1a94471b3faa210d36ce58fc23", chartsDir + "event_hub_sample_event_logger-1.0.0.tgz");
}


func checkMd5(t *testing.T, expected, file string) {
	f, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		t.Fatal(err);
	}

	actual := hex.EncodeToString(h.Sum(nil));
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("checksum failed for %s, expected: %s actual: %s", file, expected, actual);
	}	
}
