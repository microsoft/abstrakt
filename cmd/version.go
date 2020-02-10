package cmd

import (
	logger "github.com/microsoft/abstrakt/tools/logger"
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
			logger.Infof("abstrakt version %v", Version())
		},
	})

	return cc
}

// Version returns the version of abstrakt running.
func Version() string {
	return "0.0.1"
}
