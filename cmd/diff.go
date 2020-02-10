package cmd

import (
	"bytes"
	"fmt"
	"github.com/microsoft/abstrakt/internal/diff"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/tools/logger"
	"github.com/spf13/cobra"
	"strings"
)

type diffCmd struct {
	constellationFilePathOrg string
	constellationFilePathNew string
	showOriginal             *bool
	showNew                  *bool
	*baseCmd
}

func newDiffCmd() *diffCmd {
	cc := &diffCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "diff",
		Short: "Graphviz dot notation comparing two constellations",
		Long: `Diff is for producing a Graphviz dot notation representation of the difference between two constellations (line an old and new version)
	
Example: abstrakt diff -o [constellationFilePathOriginal] -n [constellationFilePathNew]`,
		SilenceUsage:  true,
		SilenceErrors: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Debug("args: " + strings.Join(args, " "))
			logger.Debugf("constellationFilePathOrg: %v", cc.constellationFilePathOrg)
			logger.Debugf("constellationFilePathNew: %v", cc.constellationFilePathNew)

			logger.Debugf("showOriginalOutput: %t", *cc.showOriginal)
			logger.Debugf("showNewOutput: %t", *cc.showNew)

			dsGraphOrg := new(constellation.Config)
			err := dsGraphOrg.LoadFile(cc.constellationFilePathOrg)
			if err != nil {
				return fmt.Errorf("Constellation config failed to load file %q: %s", cc.constellationFilePathOrg, err)
			}

			if *cc.showOriginal {
				out := &bytes.Buffer{}
				var resStringOrg string
				resStringOrg, err = dsGraphOrg.GenerateGraph(out)
				if err != nil {
					return err
				}
				logger.Output(resStringOrg)
			}

			dsGraphNew := new(constellation.Config)
			err = dsGraphNew.LoadFile(cc.constellationFilePathNew)
			if err != nil {
				return fmt.Errorf("Constellation config failed to load file %q: %s", cc.constellationFilePathNew, err)
			}

			if *cc.showNew {
				out := &bytes.Buffer{}
				var resStringNew string
				resStringNew, err = dsGraphNew.GenerateGraph(out)
				if err != nil {
					return err
				}
				logger.Output(resStringNew)
			}

			constellationSets := diff.Compare{Original: dsGraphOrg, New: dsGraphNew}
			resStringDiff, err := constellationSets.CompareConstellations()

			if err != nil {
				return err
			}

			logger.Output(resStringDiff)

			return nil
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePathOrg, "constellationFilePathOriginal", "o", "", "original or base constellation file path")
	cc.cmd.Flags().StringVarP(&cc.constellationFilePathNew, "constellationFilePathNew", "n", "", "new or changed constellation file path")
	cc.showOriginal = cc.cmd.Flags().Bool("showOriginalOutput", false, "will additionally produce dot notation for original constellation")
	cc.showNew = cc.cmd.Flags().Bool("showNewOutput", false, "will additionally produce dot notation for new constellation")
	_ = cc.cmd.MarkFlagRequired("constellationFilePathOriginal")
	_ = cc.cmd.MarkFlagRequired("constellationFilePathNew")

	return cc
}
