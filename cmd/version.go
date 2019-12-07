package cmd

import (
	logger "github.com/microsoft/abstrakt/internal/tools/logger"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	*baseCmd
}

func newVersionCmd() *versionCmd {
	cc := &versionCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "version",
		Short: "The version of Abstrakt being used",
		Long:  "The version of Abstrakt being used",
		Run: func(cmd *cobra.Command, args []string) {
			PrintVersion()
		},
	})

	return cc
}

// PrintVersion prints the current version of Abstrakt being used.
func PrintVersion() {
	logger.Info("abstrakt version 0.0.1")
}
