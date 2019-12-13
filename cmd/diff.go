package cmd

// visualise is a subcommand that constructs a graph representation of the yaml
// input file and renders this into GraphViz 'dot' notation.
// Initial version renders to dot syntax only, to graphically depict this the output
// has to be run through a graphviz visualisation tool/utiliyy

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	set "github.com/deckarep/golang-set"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	"github.com/microsoft/abstrakt/internal/tools/logger"
	"github.com/spf13/cobra"
	// "os"
	"strings"
)

type diffCmd struct {
	constellationFilePathOrg string
	constellationFilePathNew string
	showOriginal             bool
	showNew                  bool
	*baseCmd
}

type setsForComparison struct {
	setCommonSvcs set.Set
	setCommonRels set.Set
	setAddedSvcs  set.Set
	setAddedRels  set.Set
	setDelSvcs    set.Set
	setDelRels    set.Set
}

func newDiffCmd() *diffCmd {
	cc := &diffCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "diff",
		Short: "Graphviz dot notation comparing two constellations",
		Long: `Diff is for producing a Graphviz dot notation representation of the difference between two constellations (line an old and new version)
	
Example: abstrakt diff -o [constellationFilePathOriginal] -n [constellationFilePathNew]`,

		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Debug("args: " + strings.Join(args, " "))
			logger.Debug("constellationFilePathOrg: " + cc.constellationFilePathOrg)
			logger.Debug("constellationFilePathNew: " + cc.constellationFilePathNew)

			if cmd.Flag("showOriginalOutput").Value.String() == "true" {
				logger.Debug("showOriginalOutput: true")
				cc.showOriginal = true
			} else {
				logger.Debug("showOriginalOutput: false")
				cc.showOriginal = false
			}
			if cmd.Flag("showNewOutput").Value.String() == "true" {
				logger.Debug("showNewOutput: true")
				cc.showNew = true
			} else {
				logger.Debug("showNewOutput: false")
				cc.showNew = false
			}

			if !fileExists(cc.constellationFilePathOrg) {
				return fmt.Errorf("Could not open original YAML input file for reading %v", cc.constellationFilePathOrg)
			}

			if !fileExists(cc.constellationFilePathNew) {
				return fmt.Errorf("Could not open new YAML input file for reading %v", cc.constellationFilePathNew)
			}

			dsGraphOrg := dagconfigservice.NewDagConfigService()
			err := dsGraphOrg.LoadDagConfigFromFile(cc.constellationFilePathOrg)
			if err != nil {
				logger.Fatalf("dagConfigService failed to load file %q: %s", cc.constellationFilePathOrg, err)
			}

			if cc.showOriginal {
				resStringOrg := generateGraph(dsGraphOrg)
				logger.Output(resStringOrg)
			}

			dsGraphNew := dagconfigservice.NewDagConfigService()
			err = dsGraphNew.LoadDagConfigFromFile(cc.constellationFilePathNew)
			if err != nil {
				logger.Fatalf("dagConfigService failed to load file %q: %s", cc.constellationFilePathNew, err)
			}

			if cc.showNew {
				resStringNew := generateGraph(dsGraphNew)
				logger.Output(resStringNew)
			}

			resStringDiff := compareConstellations(dsGraphOrg, dsGraphNew)
			logger.Output(resStringDiff)

			return nil
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePathOrg, "constellationFilePathOriginal", "o", "", "original or base constellation file path")
	cc.cmd.Flags().StringVarP(&cc.constellationFilePathNew, "constellationFilePathNew", "n", "", "new or changed constellation file path")
	cc.cmd.Flags().Bool("showOriginalOutput", false, "will additionally produce dot notation for original constellation")
	cc.cmd.Flags().Bool("showNewOutput", false, "will additionally produce dot notation for new constellation")
	if err := cc.cmd.MarkFlagRequired("constellationFilePathOriginal"); err != nil {
		logger.Fatalf("error adding node: %v", err)
	}
	if err := cc.cmd.MarkFlagRequired("constellationFilePathNew"); err != nil {
		logger.Fatalf("error adding node: %v", err)
	}

	return cc
}

//compareConstellations
func compareConstellations(dsOrg dagconfigservice.DagConfigService, dsNew dagconfigservice.DagConfigService) string {
	sets := &setsForComparison{}

	// populate comparison sets with changes between original and new graph
	fillComparisonSets(dsOrg, dsNew, sets)

	// build the graphviz output from new graph and comparison sets
	resString := createGraphWithChanges(dsNew, sets)

	return resString
}

// fillComparisonSets - loads provided set struct with data from the constellations and then determines the various differences between the
// sets (original constellation and new) to help detemine what has been added, removed or changed.
func fillComparisonSets(dsOrg dagconfigservice.DagConfigService, dsNew dagconfigservice.DagConfigService, sets *setsForComparison) {
	setOrgSvcs, setOrgRel := createSet(dsOrg)
	setNewSvcs, setNewRel := createSet(dsNew)

	// "Added" means in the new constellation but not the original one
	// "Deleted" means something is present in the original constellation but not in the new constellation
	// Services
	sets.setCommonSvcs = setOrgSvcs.Intersect(setNewSvcs) //present in both
	sets.setAddedSvcs = setNewSvcs.Difference(setOrgSvcs) //pressent in new but not in original
	sets.setDelSvcs = setOrgSvcs.Difference(setNewSvcs)   //present in original but not in new
	// Relationships
	sets.setCommonRels = setOrgRel.Intersect(setNewRel) //present in both
	sets.setAddedRels = setNewRel.Difference(setOrgRel) //pressent in new but not in original
	sets.setDelRels = setOrgRel.Difference(setNewRel)   //present in original but not in new
}

// createSet - utility function used to create a pair of result sets (services + relationships) based on an input constellation DAG
func createSet(dsGraph dagconfigservice.DagConfigService) (set.Set, set.Set) {

	// Create sets to hold services and relationships - used to find differences between old and new using intersection and difference operations
	retSetServices := set.NewSet()
	retSetRelationships := set.NewSet()

	//Store all services in the services set
	for _, v := range dsGraph.Services {
		retSetServices.Add(v.ID)
	}

	//Store relationships in the relationship set
	for _, v := range dsGraph.Relationships {
		retSetRelationships.Add(v.From + "|" + v.To)
	}

	return retSetServices, retSetRelationships
}

// createGraphWithChanges - use both input constellations (new and original) as well as the comparison sets to create
// a dag that can be visualised. It uses the comparison sets to identify additions, deletions and changes between the original
// and new constellations.
func createGraphWithChanges(newGraph dagconfigservice.DagConfigService, sets *setsForComparison) string {
	// Lookup is used to map IDs to names. Names are easier to visualise but IDs are more important to ensure the
	// presented constellation is correct and IDs are used to link nodes together
	lookup := make(map[string]string)
	g := gographviz.NewGraph()

	// Replace spaces with underscores, names with spaces can break graphviz engines
	if err := g.SetName(strings.Replace(newGraph.Name, " ", "_", -1) + "_diff"); err != nil {
		logger.Fatalf("error setting graph name: %v", err)
	}
	// Attribute in graphviz to change graph orientation - LR indicates Left to Right. Default is top to bottom
	if err := g.AddAttr(g.Name, "rankdir", "LR"); err != nil {
		logger.Fatalf("error adding node: %v", err)
	}

	// Make the graph directed (a constellation is  DAG)
	if err := g.SetDir(true); err != nil {
		logger.Fatalf("error: %v", err)
	}

	// Add all services from the new constellation
	// - New services - highlight with colour (i.e in setAddedSvcs)
	// - Deleted services (i.e. in setDelSvcs) - include and format appropriately
	for _, v := range newGraph.Services {
		logger.Debugf("Adding node %s", v.ID)
		newName := strings.Replace(v.ID, " ", "_", -1) // Replace spaces in names with underscores, names with spaces can break graphviz engines)

		// attributes are graphviz specific that control formatting. The linrary used requires then in the form of a map passed as an argument so it needs
		// to be built before adding the item
		attrs := make(map[string]string)

		// Check if the service is new to this constellation
		if sets.setAddedSvcs.Contains(newName) {
			attrs["color"] = "\"#d8ffa8\""
		}
		attrs["label"] = "\"" + v.Type + "\n" + v.ID + "\""
		fillShapeAndStyleForNodeType(v.Type, attrs)

		lookup[v.ID] = newName
		err := g.AddNode(newGraph.Name, "\""+newName+"\"", attrs) //Surround names/labels with quotes, stops graphviz seeing special characters and breaking
		if err != nil {
			logger.Fatalf("error: %v", err)
		}

	}

	//======================================== process deleted services ==========================================
	//Process services that have been removed from the original constellation
	for v := range sets.setDelSvcs.Iter() {
		logger.Debug(v)
		vString, _ := v.(string)
		newName := strings.Replace(vString, " ", "_", -1)
		attrs := make(map[string]string)
		attrs["color"] = "\"#ff9494\""
		fillShapeAndStyleForNodeType("", attrs)
		logger.Debug("Adding deleted service ", newName)
		if err := g.AddNode(newGraph.Name, "\""+newName+"\"", attrs); err != nil {
			logger.Fatalf("error adding node: %v", err)
		}

	}

	//======================================== process  relationships ==========================================
	// Add Relationships from the new constellation
	//  - New relationships - highlight with colour (i.e. in setAddedSvcs)
	// - Deleted relationships (i.e. in setDelRels) - include and format appropriately
	for _, v := range newGraph.Relationships {
		logger.Debugf("Adding relationship from %s ---> %s", v.From, v.To)
		//Surround names/labels with quotes, stops graphviz seeing special characters and breaking
		localFrom := "\"" + lookup[v.From] + "\""
		localTo := "\"" + lookup[v.To] + "\""

		attrs := make(map[string]string)
		// Relationship is stored in the map as source|destination - both are needed to tell if a relationship is new (this is how they are stored in the sets created earlier)
		relLookupName := lookup[v.From] + "|" + lookup[v.To]

		// Check if the relationship is new (it will then not be in the common list of relationships between old and new constellation)
		if sets.setAddedRels.Contains(relLookupName) {
			attrs["color"] = "\"#d8ffa8\""
		}

		err := g.AddEdge(localFrom, localTo, true, attrs)
		if err != nil {
			logger.Fatalf("error: %v", err)
		}
	}

	//======================================== process deleted relationships =====================================
	// Deleted relationships
	for v := range sets.setDelRels.Iter() {
		logger.Debug(v)
		vString := strings.Replace(v.(string), " ", "_", -1)
		if len(strings.Split(vString, "|")) != 2 {
			logger.Fatal("Relationships string should be two items separated by | but got ", vString)
		}
		newFrom := "\"" + strings.Split(vString, "|")[0] + "\""
		newTo := "\"" + strings.Split(vString, "|")[1] + "\""

		attrs := make(map[string]string)
		attrs["color"] = "\"#ff9494\""
		logger.Debug("Adding deleted service  relationship from ", newFrom, " to ", newTo)
		err := g.AddEdge(newFrom, newTo, true, attrs)
		if err != nil {
			logger.Fatalf("error: %v", err)
		}
	}

	// Produce resulting graph in dot notation format
	return g.String()
}

// Populate attrs map with additional attributes based on the node type
// Easier to change in a single place
func fillShapeAndStyleForNodeType(nodeType string, existAttrs map[string]string) {

	switch nodeType {
	case "EventLogger":
		existAttrs["shape"] = "rectangle"
		existAttrs["style"] = "\"rounded, filled\""
	case "EventGenerator":
		existAttrs["shape"] = "rectangle"
		existAttrs["style"] = "\"rounded, filled\""
	case "EventHub":
		existAttrs["shape"] = "rectangle"
		existAttrs["style"] = "\"rounded, filled\""
	default:
		existAttrs["shape"] = "rectangle"
		existAttrs["style"] = "\"rounded, filled\""
	}

}
