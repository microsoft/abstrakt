package cmd

import (
	"fmt"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/microsoft/abstrakt/internal/tools/logger"
	"github.com/spf13/cobra"
)

type validateCmd struct {
	constellationFilePath string
	mapperFilePath        string
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
			if len(cc.constellationFilePath) == 0 && len(cc.mapperFilePath) == 0 {
				return fmt.Errorf("No arguments passed")
			}

			fail := false

			if len(cc.mapperFilePath) > 0 {
				err = loadAndTestMapper(cc.mapperFilePath)
				if err != nil {
					logger.Errorf("Mapper: %v", err)
					fail = true
				} else {
					logger.Info("Mapper: valid")
				}
			}

			if len(cc.constellationFilePath) > 0 {
				err = loadAndTestDag(cc.constellationFilePath)
				if err != nil {
					logger.Errorf("Constellation: %v", err)
					fail = true
				} else {
					logger.Info("Constellation: valid")
				}
			}

			if fail {
				err = fmt.Errorf("Invalid configuration(s)")
			}

			return
		},
	})

	cc.cmd.Flags().StringVarP(&cc.constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	cc.cmd.Flags().StringVarP(&cc.mapperFilePath, "mapperFilePath", "m", "", "mapper file path")

	return cc
}

func loadAndTestDag(path string) (err error) {
	d := new(constellation.Config)
	err = d.LoadFile(path)

	if err != nil {
		return
	}

	return validateDag(d)
}

// validateDag takes a constellation dag and returns any errors.
func validateDag(dag *constellation.Config) (err error) {
	err = dag.ValidateModel()

	if err != nil {
		return
	}

	duplicates := dag.CheckDuplicates()

	if duplicates != nil {
		logger.Error("Duplicate IDs found:")
		for _, i := range duplicates {
			logger.Errorf("'%v'", i)
		}
		err = error(fmt.Errorf("Constellation is invalid"))
	}

	connections := dag.CheckServiceExists()

	if len(connections) > 0 {
		for key, i := range connections {
			logger.Errorf("Relationship '%v' has missing Services:", key)
			for _, j := range i {
				logger.Errorf("'%v'", j)
			}
		}
		err = error(fmt.Errorf("Constellation is invalid"))
	}

	return
}

func loadAndTestMapper(path string) (err error) {
	m := new(mapper.Config)
	err = m.LoadFile(path)

	if err != nil {
		return
	}

	return validateMapper(m)
}

// validateMapper takes a constellation mapper and returns any errors.
func validateMapper(mapper *mapper.Config) (err error) {
	err = mapper.ValidateModel()

	if err != nil {
		return
	}

	duplicates := mapper.CheckDuplicates()

	if duplicates != nil {
		logger.Error("Duplicate IDs found:")
		for _, i := range duplicates {
			logger.Errorf("'%v'", i)
		}
		err = error(fmt.Errorf("Constellation is invalid"))
	}

	return
}
