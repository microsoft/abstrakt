package cmd

import (
	logger "github.com/microsoft/abstrakt/tools/logger"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	*baseCmd
}

var (
	version = "edge"
	commit  = "n/a"
)

func newVersionCmd() *versionCmd {
	cc := &versionCmd{}

	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "version",
		Short: "The version of Abstrakt being used",
		Long:  "The version of Abstrakt being used",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("abstrakt version %v, commit %v", Version(), Commit())
		},
	})

	return cc
}

// Version returns the version of abstrakt running.
func Version() string {
	return version
}

// Commit returns the git commit SHA for the code that abstrakt was built from.
func Commit() string {
	return commit
}
