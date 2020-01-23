package diff

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	set "github.com/deckarep/golang-set"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"strings"
)

// ComparisonSet is a combination of two graphviz graphs.
type ComparisonSet struct {
	SetCommonSvcs set.Set
	SetCommonRels set.Set
	SetAddedSvcs  set.Set
	SetAddedRels  set.Set
	SetDelSvcs    set.Set
	SetDelRels    set.Set
}

// Set contains two constellation configurations to compare against one another.
type Set struct {
	Original *constellation.Config
	New      *constellation.Config
}

//CompareConstellations takes two constellation configurations, compares and returns the differences.
func (d *Set) CompareConstellations() (string, error) {
	// populate comparison sets with changes between original and new graph
	sets := d.FillComparisonSets()

	// build the graphviz output from new graph and comparison sets
	return CreateGraphWithChanges(d.New, &sets)
}

// FillComparisonSets - loads provided set struct with data from the constellations and then determines the various differences between the
// sets (original constellation and new) to help detemine what has been added, removed or changed.
func (d *Set) FillComparisonSets() (sets ComparisonSet) {
	setOrgSvcs, setOrgRel := createSet(d.Original)
	setNewSvcs, setNewRel := createSet(d.New)

	// "Added" means in the new constellation but not the original one
	// "Deleted" means something is present in the original constellation but not in the new constellation
	// Services
	sets.SetCommonSvcs = setOrgSvcs.Intersect(setNewSvcs) //present in both
	sets.SetAddedSvcs = setNewSvcs.Difference(setOrgSvcs) //pressent in new but not in original
	sets.SetDelSvcs = setOrgSvcs.Difference(setNewSvcs)   //present in original but not in new
	// Relationships
	sets.SetCommonRels = setOrgRel.Intersect(setNewRel) //present in both
	sets.SetAddedRels = setNewRel.Difference(setOrgRel) //pressent in new but not in original
	sets.SetDelRels = setOrgRel.Difference(setNewRel)   //present in original but not in new

	return
}

// createSet - utility function used to create a pair of result sets (services + relationships) based on an input constellation DAG
func createSet(dsGraph *constellation.Config) (set.Set, set.Set) {

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

// CreateGraphWithChanges - use both input constellations (new and original) as well as the comparison sets to create
// a dag that can be visualized. It uses the comparison sets to identify additions, deletions and changes between the original
// and new constellations.
func CreateGraphWithChanges(newGraph *constellation.Config, sets *ComparisonSet) (string, error) {
	// Lookup is used to map IDs to names. Names are easier to visualize but IDs are more important to ensure the
	// presented constellation is correct and IDs are used to link nodes together
	lookup := make(map[string]string)
	g := gographviz.NewGraph()

	// Replace spaces with underscores, names with spaces can break graphviz engines
	if err := g.SetName(strings.Replace(newGraph.Name, " ", "_", -1) + "_diff"); err != nil {
		return "", fmt.Errorf("error setting graph name: %v", err)
	}
	// Attribute in graphviz to change graph orientation - LR indicates Left to Right. Default is top to bottom
	if err := g.AddAttr(g.Name, "rankdir", "LR"); err != nil {
		return "", fmt.Errorf("error adding node: %v", err)
	}

	// Make the graph directed (a constellation is  DAG)
	if err := g.SetDir(true); err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	// Add all services from the new constellation
	// - New services - highlight with color (i.e in setAddedSvcs)
	// - Deleted services (i.e. in setDelSvcs) - include and format appropriately
	for _, v := range newGraph.Services {
		newName := strings.Replace(v.ID, " ", "_", -1) // Replace spaces in names with underscores, names with spaces can break graphviz engines)

		// attributes are graphviz specific that control formatting. The linrary used requires then in the form of a map passed as an argument so it needs
		// to be built before adding the item
		attrs := make(map[string]string)

		// Check if the service is new to this constellation
		if sets.SetAddedSvcs.Contains(newName) {
			attrs["color"] = "\"#d8ffa8\""
		}
		attrs["label"] = "\"" + v.Type + "\n" + v.ID + "\""
		fillShapeAndStyleForNodeType(v.Type, attrs)

		lookup[v.ID] = newName
		err := g.AddNode(newGraph.Name, "\""+newName+"\"", attrs) //Surround names/labels with quotes, stops graphviz seeing special characters and breaking
		if err != nil {
			return "", fmt.Errorf("error: %v", err)
		}
	}

	//======================================== process deleted services ==========================================
	//Process services that have been removed from the original constellation
	for v := range sets.SetDelSvcs.Iter() {
		vString, _ := v.(string)
		newName := strings.Replace(vString, " ", "_", -1)
		attrs := make(map[string]string)
		attrs["color"] = "\"#ff9494\""
		fillShapeAndStyleForNodeType("", attrs)
		if err := g.AddNode(newGraph.Name, "\""+newName+"\"", attrs); err != nil {
			return "", fmt.Errorf("error adding node: %v", err)
		}
	}

	//======================================== process  relationships ==========================================
	// Add Relationships from the new constellation
	//  - New relationships - highlight with color (i.e. in setAddedSvcs)
	// - Deleted relationships (i.e. in setDelRels) - include and format appropriately
	for _, v := range newGraph.Relationships {
		//Surround names/labels with quotes, stops graphviz seeing special characters and breaking
		localFrom := "\"" + lookup[v.From] + "\""
		localTo := "\"" + lookup[v.To] + "\""

		attrs := make(map[string]string)
		// Relationship is stored in the map as source|destination - both are needed to tell if a relationship is new (this is how they are stored in the sets created earlier)
		relLookupName := lookup[v.From] + "|" + lookup[v.To]

		// Check if the relationship is new (it will then not be in the common list of relationships between old and new constellation)
		if sets.SetAddedRels.Contains(relLookupName) {
			attrs["color"] = "\"#d8ffa8\""
		}

		err := g.AddEdge(localFrom, localTo, true, attrs)
		if err != nil {
			return "", fmt.Errorf("error: %v", err)
		}
	}

	//======================================== process deleted relationships =====================================
	// Deleted relationships
	for v := range sets.SetDelRels.Iter() {
		vString := strings.Replace(v.(string), " ", "_", -1)
		if len(strings.Split(vString, "|")) != 2 {
			return "", fmt.Errorf("Relationships string should be two items separated by | but got %v", vString)
		}
		newFrom := "\"" + strings.Split(vString, "|")[0] + "\""
		newTo := "\"" + strings.Split(vString, "|")[1] + "\""

		attrs := make(map[string]string)
		attrs["color"] = "\"#ff9494\""
		err := g.AddEdge(newFrom, newTo, true, attrs)
		if err != nil {
			return "", fmt.Errorf("error: %v", err)
		}
	}

	// Produce resulting graph in dot notation format
	return g.String(), nil
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
