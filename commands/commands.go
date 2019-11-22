package commands

import (
	"github.com/microsoft/abstrakt/internal/tools/logger"
	cobra "github.com/spf13/cobra"
)

// DefaultRootCommand returns the default (aka root) command for  command.
func DefaultRootCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "abstrakt",
		Short: "Scalable, config driven data pipelines for Kubernetes.",
		Long:  "Scalable, config driven data pipelines for Kubernetes.",
	}

	c.PersistentPreRunE = func(cmd *cobra.Command, args []string) (err error) {
		verbose := cmd.Flag("verbose").Value.String()

		if verbose == "true" {
			logger.SetLevelDebug()
		} else {
			logger.SetLevelInfo()
		}

		return nil
	}

	c.AddCommand(
		composeCmd, versionCmd, visualiseCmd,
	)

	return c
}
