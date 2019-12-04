package cmd

// visualise is a subcommand that constructs a graph representation of the yaml
// input file and renders this into GraphViz 'dot' notation.
// Initial version renders to dot syntax only, to graphically depict this the output
// has to be run through a graphviz visualisation tool/utiliyy

import (
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	"github.com/microsoft/abstrakt/internal/tools/logger"
)

var visualiseCmd = &cobra.Command{
	Use:   "visualise",
	Short: "Format a constellation configuration as Graphviz dot notation",
	Long: `Visualise is for producing Graphviz dot notation code of a constellation configuration

Example: abstrakt visualise -f [constellationFilePath]`,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("args: " + strings.Join(args, " "))
		logger.Debug("constellationFilePath: " + constellationFilePath)

		if !fileExists(constellationFilePath) {
			logger.Error("Could not open YAML input file for reading")
			os.Exit(-1)
		}

		dsGraph := dagconfigservice.NewDagConfigService()
		err := dsGraph.LoadDagConfigFromFile(constellationFilePath)
		if err != nil {
			logger.Fatalf("dagConfigService failed to load file %q: %s", constellationFilePath, err)
		}

		resString := generateGraph(dsGraph)
		logger.Output(resString)
	},
}

func init() {
	visualiseCmd.Flags().StringVarP(&constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	err := visualiseCmd.MarkFlagRequired("constellationFilePath")
	if err != nil {
		logger.Panic(err)
	}
}

// fileExists - basic utility function to check the provided filename can be opened and is not a folder/directory
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// generateGraph - function to take a dagconfigService structure and create a graph object that contains the
// representation of the graph. Also outputs a string representation (GraphViz dot notation) of the resulting graph
// this can be passed on to GraphViz to graphically render the resulting graph
func generateGraph(readGraph dagconfigservice.DagConfigService) string {

	// Lookup is used to map IDs to names. Names are easier to visualise but IDs are more important to ensure the
	// presented constellation is correct and IDs are used to link nodes together
	lookup := make(map[string]string)

	g := gographviz.NewGraph()

	// Replace spaces with underscores, names with spaces can break graphviz engines
	if err := g.SetName(strings.Replace(readGraph.Name, " ", "_", -1)); err != nil {
		logger.Fatalf("error: %v", err)
		logger.Panic(err)
	}

	// Make the graph directed (a constellation is  DAG)
	if err := g.SetDir(true); err != nil {
		logger.Fatalf("error: %v", err)
		logger.Panic(err)
	}

	// Add all nodes to the graph storing the lookup from ID to name (for later adding relationships)
	// Replace spaces in names with underscores, names with spaces can break graphviz engines)
	for _, v := range readGraph.Services {
		logger.Debugf("Adding node %s %s\n", v.ID, v.Name)
		newName := strings.Replace(v.Name, " ", "_", -1)
		lookup[v.Name] = newName
		err := g.AddNode(readGraph.Name, newName, nil)
		if err != nil {
			logger.Panic(err)
		}
	}

	// Add relationships to the graph linking using the lookup IDs to name map
	// Replace spaces in names with underscores, names with spaces can break graphviz engines)
	for _, v := range readGraph.Relationships {
		logger.Debugf("Adding relationship from %s ---> %s\n", v.From, v.To)
		localFrom := lookup[v.From]
		localTo := lookup[v.To]
		err := g.AddEdge(localFrom, localTo, true, nil)
		if err != nil {
			logger.Panic(err)
		}
	}

	// Produce resulting graph in dot notation format
	return g.String()
}
