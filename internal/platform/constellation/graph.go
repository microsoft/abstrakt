package constellation

import (
	"fmt"
	"io"
	"strings"

	"github.com/awalterschulze/gographviz"
)

// GenerateGraph - function to take a dagconfigService structure and create a graph object that contains the
// representation of the graph. Also outputs a string representation (GraphViz dot notation) of the resulting graph
// this can be passed on to GraphViz to graphically render the resulting graph
func (readGraph *Config) GenerateGraph(out io.Writer) (string, error) {

	// Lookup is used to map IDs to names. Names are easier to visualise but IDs are more important to ensure the
	// presented constellation is correct and IDs are used to link nodes together
	lookup := make(map[string]string)

	g := gographviz.NewGraph()

	// Replace spaces with underscores, names with spaces can break graphviz engines
	if err := g.SetName(strings.Replace(readGraph.Name, " ", "_", -1)); err != nil {
		return "", err
	}
	if err := g.AddAttr(g.Name, "rankdir", "LR"); err != nil {
		return "", err
	}

	// Make the graph directed (a constellation is  DAG)
	if err := g.SetDir(true); err != nil {
		return "", err
	}

	// Add all nodes to the graph storing the lookup from ID to name (for later adding relationships)
	// Replace spaces in names with underscores, names with spaces can break graphviz engines)
	for _, v := range readGraph.Services {
		fmt.Fprintf(out, "Adding node %s\n", v.ID)
		newName := strings.Replace(v.ID, " ", "_", -1)

		if strings.Compare(newName, v.ID) != 0 {
			fmt.Fprintf(out, "Changing %s to %s\n", v.ID, newName)
		}

		lookup[v.ID] = newName
		err := g.AddNode(readGraph.Name, "\""+newName+"\"", nil)
		if err != nil {
			return "", err
		}
	}

	// Add relationships to the graph linking using the lookup IDs to name map
	// Replace spaces in names with underscores, names with spaces can break graphviz engines)
	for _, v := range readGraph.Relationships {
		fmt.Fprintf(out, "Adding relationship from %s ---> %s\n", v.From, v.To)
		localFrom := "\"" + lookup[v.From] + "\""
		localTo := "\"" + lookup[v.To] + "\""
		err := g.AddEdge(localFrom, localTo, true, nil)
		if err != nil {
			return "", err
		}
	}

	// Produce resulting graph in dot notation format
	return g.String(), nil
}
