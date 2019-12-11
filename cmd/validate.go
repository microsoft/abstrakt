package cmd

import (
	"fmt"
	"github.com/microsoft/abstrakt/internal/dagconfigservice"
	"github.com/microsoft/abstrakt/internal/tools/logger"
	"github.com/microsoft/abstrakt/internal/validationservice"
	"github.com/spf13/cobra"
)

type validateCmd struct {
	constellationFilePath string
	*baseCmd
}

func newValidateCmd() *validateCmd {
	cc := &validateCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "validate",
		Short: "Validate a constellation file for correct schema and ensure correctness.",
		Long: `Validate is used to ensure the correctness of a constellation file.
	
Example: abstrakt validate -f [constellationFilePath]`,
		SilenceUsage:  true,
		SilenceErrors: true,

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			d := dagconfigservice.NewDagConfigService()
			err = d.LoadDagConfigFromFile(cc.constellationFilePath)

			if err != nil {
				return
			}

			err = validateDag(&d)

			if err == nil {
				logger.Info("Constellation is valid.")
			}

			return
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	_ = cc.cmd.MarkFlagRequired("constellationFilePath")

	return cc
}

// validateDag takes a constellation dag and returns any errors.
func validateDag(dag *dagconfigservice.DagConfigService) (err error) {
	service := validationservice.Validator{Config: dag}

	err = service.ValidateModel()

	if err != nil {
		return
	}

	duplicates := service.CheckDuplicates()

	if duplicates != nil {
		logger.Error("Duplicate IDs found:")
		for _, i := range duplicates {
			logger.Error(i)
		}
		err = error(fmt.Errorf("Constellation is invalid"))
	}

	connections := service.CheckServiceExists()

	if len(connections) > 0 {
		for key, i := range connections {
			logger.Errorf("Relationship '%v' has missing Services:", key)
			for _, j := range i {
				logger.Error(j)
			}
		}
		err = error(fmt.Errorf("Constellation is invalid"))
	}

	return
}
