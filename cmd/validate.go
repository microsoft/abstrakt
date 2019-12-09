package cmd

import (
	"github.com/microsoft/abstrakt/internal/tools/logger"
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

		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info(cc.constellationFilePath)

			return nil
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	_ = cc.cmd.MarkFlagRequired("constellationFilePath")

	return cc
}
