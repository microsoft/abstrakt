package cmd

// visualize is a subcommand that constructs a graph representation of the yaml
// input file and renders this into GraphViz 'dot' notation.
// Initial version renders to dot syntax only, to graphically depict this the output
// has to be run through a graphviz visualization tool/utiliyy

import (
	"bytes"
	"fmt"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/tools/file"
	"github.com/microsoft/abstrakt/internal/tools/logger"
	"github.com/spf13/cobra"
	"strings"
)

type visualizeCmd struct {
	constellationFilePath string
	*baseCmd
}

func newVisualizeCmd() *visualizeCmd {
	cc := &visualizeCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "visualize",
		Short: "Format a constellation configuration as Graphviz dot notation",
		Long: `Visualize is for producing Graphviz dot notation code of a constellation configuration
	
Example: abstrakt visualize -f [constellationFilePath]`,

		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Debug("args: " + strings.Join(args, " "))
			logger.Debug("constellationFilePath: " + cc.constellationFilePath)

			if !file.Exists(cc.constellationFilePath) {
				return fmt.Errorf("Could not open YAML input file for reading %v", cc.constellationFilePath)
			}

			dsGraph := new(constellation.Config)
			err := dsGraph.LoadFile(cc.constellationFilePath)
			if err != nil {
				return fmt.Errorf("dagConfigService failed to load file %q: %s", cc.constellationFilePath, err)
			}

			out := &bytes.Buffer{}
			resString, err := dsGraph.GenerateGraph(out)
			if err != nil {
				return err
			}

			logger.PrintBuffer(out, true)
			logger.Output(resString)

			return nil
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	_ = cc.cmd.MarkFlagRequired("constellationFilePath")

	return cc
}
