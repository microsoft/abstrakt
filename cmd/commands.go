package cmd

import (
	"github.com/microsoft/abstrakt/internal/tools/logger"
	cobra "github.com/spf13/cobra"
)

type baseCmd struct {
	cmd *cobra.Command
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

type cmder interface {
	getCommand() *cobra.Command
}

func addCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}

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

	addCommands(c, newComposeCmd(), newVersionCmd(), newVisualiseCmd())

	return c
}
