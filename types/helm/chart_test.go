package helm;

import "testing";
import "strings"

type ConfigTest struct {
	name 	string
	description string
	version string
}

func (c ConfigTest) Name() string {
	return c.name;
}

func (c ConfigTest) Description() string {
	return c.description;
}

func (c ConfigTest) Version() string {
	return c.version;
}

func TestNameSubstition(t *testing.T) {
	expected := "name: hello-world"
	config := ConfigTest { name: "hello-world" }
	testChartEntry(t, expected, config)
}

func TestDescriptionSubstition(t *testing.T) {
	expected := "description: A Helm chart for Kubernetes"
	config := ConfigTest { description: "A Helm chart for Kubernetes" }
	testChartEntry(t, expected, config)
}

func TestAppVersionSubstition(t *testing.T) {
	expected := "version: 0.1.0"
	config := ConfigTest { version: "0.1.0" }
	testChartEntry(t, expected, config)
}

func TestVersion(t *testing.T) {
	expected := "appVersion: 0.1.0"
	config := ConfigTest { version: "0.1.0" }
	testChartEntry(t, expected, config)
}

func TestApiVersion(t *testing.T) {
	expected := "apiVersion: v2"
	config := ConfigTest { }
	testChartEntry(t, expected, config)
}

func TestType(t *testing.T) {
	expected := "type: application"
	config := ConfigTest { }
	testChartEntry(t, expected, config)
}

func testChartEntry(t *testing.T, expected string, config ConfigTest) {
	test, err := GenerateChart(config)

	if(err != nil) {
		t.Errorf("GenerateChart() -- %v", err)
	}
	if !strings.Contains(*test, expected) {
		t.Errorf("GenerateChart() = %v, expected to conatain %v", *test, expected)
	}
}
