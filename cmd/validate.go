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
				return err
			}

			service := validationservice.Validator{Config: &d}

			valid := service.CheckDuplicates()

			if len(valid) > 0 {
				logger.Error("Duplicate IDs found:")
				for _, i := range valid {
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

			if err == nil {
				logger.Info("Constellation is valid.")
			}

			return err
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	_ = cc.cmd.MarkFlagRequired("constellationFilePath")

	return cc
}
