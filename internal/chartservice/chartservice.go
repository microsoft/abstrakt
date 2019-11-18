package chartservice

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
)

// LoadChartFromDir loads a Helm chart from the specified director
func LoadChartFromDir(dir string) (*chart.Chart, error) {
	h, err := loader.LoadDir(dir)

	if err != nil {
		return nil, err
	}

	return h, nil
}

// SaveChartToDir takes the chart object and saves it as a set of files in the specified director
func SaveChartToDir(chart *chart.Chart, dir string) error {
	return chartutil.SaveDir(chart, dir)
}

// ZipChartToDir compresses the chart and saves it in compiled format
func ZipChartToDir(chart *chart.Chart, dir string) (string, error) {
	return chartutil.Save(chart, dir)
}
