package chart

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/downloader"
)

const startMeta = `apiVersion: v1
name: wormhole_constellation
description: A Helm chart for Kubernetes
version: 4.3.2
home: ""`

//Create makes a new chart at the specified location
func Create(name string, dir string) (chartReturn *chart.Chart, err error) {
	tdir, err := ioutil.TempDir("./", "output-")

	if err != nil {
		return
	}

	defer func() {
		err = os.RemoveAll(tdir)
		if err != nil {
			return
		}
	}()

	err = os.Mkdir(path.Join(tdir, "templates"), 0777)

	if err != nil {
		return
	}

	files := []string{"values.yaml", path.Join("templates", "NOTES.txt")}

	for _, k := range files {
		err = ioutil.WriteFile(path.Join(tdir, k), []byte(""), 0777)
		if err != nil {
			return
		}
	}

	err = ioutil.WriteFile(path.Join(tdir, "Chart.yaml"), []byte(startMeta), 0777)

	if err != nil {
		return
	}

	cfile := &chart.Metadata{
		Name:        name,
		Description: "A Helm chart for Kubernetes",
		Type:        "application",
		Version:     "0.1.0",
		AppVersion:  "0.1.0",
		APIVersion:  chart.APIVersionV2,
	}

	err = chartutil.CreateFrom(cfile, dir, tdir)

	if err != nil {
		return
	}

	chartReturn, err = loader.LoadDir(path.Join(dir, name))

	if err != nil {
		return
	}

	return
}

// LoadFromDir loads a Helm chart from the specified director
func LoadFromDir(dir string) (*chart.Chart, error) {
	h, err := loader.LoadDir(dir)

	if err != nil {
		return nil, err
	}

	return h, nil
}

// SaveToDir takes the chart object and saves it as a set of files in the specified director
func SaveToDir(chart *chart.Chart, dir string) error {
	return chartutil.SaveDir(chart, dir)
}

// ZipToDir compresses the chart and saves it in compiled format
func ZipToDir(chart *chart.Chart, dir string) (string, error) {
	return chartutil.Save(chart, dir)
}

// Build download charts
func Build(dir string) (out *bytes.Buffer, err error) {

	out = &bytes.Buffer{}

	manager := downloader.Manager{
		Out:       out,
		ChartPath: dir,
	}

	err = manager.Build()
	return
}
