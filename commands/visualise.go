package commands

// visualise is a subcommand that constructs a graph representation of the yaml
// input file and renders this into GraphViz 'dot' notation.
// Initial version renders to dot syntax only, to graphically depict this the output
// has to be run through a graphviz visualisation tool/utiliyy

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
	"gopkg.in/yaml.v2"
)

// var constellationFilePath string
var verbose string

//Config is a struct used to parse the yaml from the constellation definition
// It was created using Yaml to Go
type Config struct {
	Name     string `yaml:"Name"`
	ID       string `yaml:"Id"`
	Services []struct {
		Name       string `yaml:"Name"`
		ID         string `yaml:"Id"`
		Type       string `yaml:"Type"`
		Properties struct {
		} `yaml:"Properties"`
	} `yaml:"Services"`
	Relationships []struct {
		Name        string `yaml:"Name"`
		ID          string `yaml:"Id"`
		Description string `yaml:"Description"`
		From        string `yaml:"From"`
		To          string `yaml:"To"`
		Properties  struct {
		} `yaml:"Properties"`
	} `yaml:"Relationships"`
}

var visualiseCmd = &cobra.Command{
	Use:   "visualise",
	Short: "format a constellation configuration as Graphviz dot notation",
	Long: "visualise is for producing Graphviz dot notation code of a constellation configuration\n" +
		"abstrakt visualise -f [constellationFilePath]",

	Run: func(cmd *cobra.Command, args []string) {
		if verbose == "true" {
			fmt.Println("args: " + strings.Join(args, " "))
			fmt.Println("constellationFilePath: " + constellationFilePath)
		}

		yamlString := readYaml(constellationFilePath)
		// fmt.Println("Finished reading %s", yamlString)
		result := parseYaml(yamlString)
		generateGraph(result)

	},
}

func init() {
	visualiseCmd.Flags().StringVarP(&constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	visualiseCmd.MarkFlagRequired("constellationFilePath")
	// visualiseCmd.Flags().Bool("verbose", true, "verbose - show logging  information")
	visualiseCmd.Flags().StringVarP(&verbose, "verbose", "v", "true", "verbose - show logging  information")
}

//fileExists - basic utility function to check the provided filename can be opened and is not a folder/directory
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//parseYaml - utility function to take in a yaml string and return a Config structure filled with the results of
//parsing the yaml file
func parseYaml(inputString string) Config {
	var config Config

	err := yaml.Unmarshal([]byte(inputString), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// fmt.Printf("--- t:\n%v\n\n", config)
	return config
}

//readYaml - read the yaml from the provided filename into a memory string variable.
func readYaml(fileName string) string {
	readData, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(readData)
}

//generateGraph - function to take a Config structure and create a graph object that contains the
//definition of the graph. Also outputs a string representation (GraphViz dot notation) of the resulting graph
//this can be passed on to GraphViz to graphically render the resulting graph
func generateGraph(readConfig Config) {

	//lookup is used to map IDs to names. Names are easier to visualise but IDs are more important to ensure the
	//presented constellation is correct and IDs are used to link nodes together
	lookup := make(map[string]string)

	g := gographviz.NewGraph()
	if err := g.SetName(strings.Replace(readConfig.Name, " ", "_", -1)); err != nil { //Replace spaces with underscores, names with spaces can break graphviz engines
		log.Fatalf("error: %v", err)
		panic(err)
	}

	if err := g.SetDir(true); err != nil { //Make the graph directed (a constellation is  DAG)
		log.Fatalf("error: %v", err)
		panic(err)
	}

	//Add all nodes to the graph storing the lookup from ID to name (for later adding relationships)
	//Replace spaces in names with underscores, names with spaces can break graphviz engines)
	for _, v := range readConfig.Services {
		if verbose == "true" {
			log.Printf("Adding node %s %s\n", v.ID, v.Name)
		}
		newName := strings.Replace(v.Name, " ", "_", -1)
		lookup[v.ID] = newName
		g.AddNode(readConfig.Name, newName, nil)
	}

	//Add relationships to the graph linking using the lookup IDs to name map
	//Replace spaces in names with underscores, names with spaces can break graphviz engines)
	for _, v := range readConfig.Relationships {
		if verbose == "true" {
			log.Printf("Adding relationship from %s ---> %s\n", v.From, v.To)
		}
		localFrom := lookup[v.From]
		localTo := lookup[v.To]
		g.AddEdge(localFrom, localTo, true, nil)
	}

	//Produce resulting graph in dot notation format
	fmt.Printf("%s", g.String())

}
