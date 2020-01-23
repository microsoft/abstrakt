package cmd

import (
	"fmt"
	"github.com/microsoft/abstrakt/internal/compose"
	"github.com/microsoft/abstrakt/internal/platform/chart"
	"github.com/microsoft/abstrakt/tools/logger"
	"github.com/spf13/cobra"
	"path"
	"strings"
)

type composeCmd struct {
	templateType          string
	constellationFilePath string
	mapsFilePath          string
	outputPath            string
	zipChart              *bool
	noChecks              *bool
	*baseCmd
}

func newComposeCmd() *composeCmd {
	cc := &composeCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "compose [chart name]",
		Short: "Compose a package into requested template type",
		Long: `Compose is for composing a package based on mapsFilePath and constellationFilePath and template (default value is helm).
	
Example: abstrakt compose [chart name] -t [templateType] -f [constellationFilePath] -m [mapsFilePath] -o [outputPath] -z --noChecks`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,

		RunE: func(cmd *cobra.Command, args []string) (err error) {

			chartName := args[0]

			logger.Info(chartName)

			if cc.templateType != "helm" && cc.templateType != "" {
				return fmt.Errorf("Template type: %v is not known", cc.templateType)
			}

			service := new(compose.Composer)
			_ = service.LoadFile(cc.constellationFilePath, cc.mapsFilePath)

			logger.Debugf("noChecks is set to %t", *cc.noChecks)

			if !*cc.noChecks {
				logger.Debug("Starting validating constellation")

				err = validateDag(&service.Constellation)

				if err != nil {
					return
				}

				logger.Debug("Finished validating constellation")
			}

			helm, err := service.Build(chartName, cc.outputPath)
			if err != nil {
				return fmt.Errorf("Could not compose: %v", err)
			}

			err = chart.SaveToDir(helm, cc.outputPath)

			if err != nil {
				return fmt.Errorf("There was an error saving the chart: %v", err)
			}

			logger.Infof("Chart was saved to: %v", cc.outputPath)

			out, err := chart.Build(path.Join(cc.outputPath, chartName))

			if err != nil {
				return fmt.Errorf("There was an error saving the chart: %v", err)
			}

			if *cc.zipChart {
				_, err = chart.ZipToDir(helm, cc.outputPath)
				if err != nil {
					return fmt.Errorf("There was an error zipping the chart: %v", err)
				}
			}

			logger.PrintBuffer(out, true)

			logger.Debugf("args: %v", strings.Join(args, " "))
			logger.Debugf("template: %v", cc.templateType)
			logger.Debugf("constellationFilePath: %v", cc.constellationFilePath)
			logger.Debugf("mapsFilePath: %v", cc.mapsFilePath)
			logger.Debugf("outputPath: %v", cc.outputPath)
			return nil
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	_ = cc.cmd.MarkFlagRequired("constellationFilePath")
	cc.cmd.Flags().StringVarP(&cc.mapsFilePath, "mapsFilePath", "m", "", "maps file path")
	_ = cc.cmd.MarkFlagRequired("mapsFilePath")
	cc.cmd.Flags().StringVarP(&cc.outputPath, "outputPath", "o", "", "destination directory")
	_ = cc.cmd.MarkFlagRequired("outputPath")
	cc.cmd.Flags().StringVarP(&cc.templateType, "template type", "t", "helm", "output template type")
	cc.zipChart = cc.cmd.Flags().BoolP("zipChart", "z", false, "zips the chart")
	cc.noChecks = cc.cmd.Flags().Bool("noChecks", false, "turn off validation checks of constellation file before composing")

	return cc
}
