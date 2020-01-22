package cmd

import (
	"fmt"
	"github.com/microsoft/abstrakt/internal/platform/constellation"
	"github.com/microsoft/abstrakt/internal/platform/mapper"
	"github.com/microsoft/abstrakt/internal/tools/helpers"
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
		Long: `Validate is used to ensure the correctness of a constellation and mapper files.
	
Example: abstrakt validate -f [constellationFilePath] -m [mapperFilePath]
         abstrakt validate -f [constellationFilePath]
         abstrakt validate -m [mapperFilePath]`,
		SilenceUsage:  true,
		SilenceErrors: true,

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(cc.constellationFilePath) == 0 && len(cc.mapperFilePath) == 0 {
				_ = cc.baseCmd.cmd.Usage()
				return fmt.Errorf("no flags were set")
			}

			var d constellation.Config
			var m mapper.Config

			fail := false

			if len(cc.mapperFilePath) > 0 {
				m, err = loadAndValidateMapper(cc.mapperFilePath)
				if err != nil {
					logger.Errorf("Mapper: %v", err)
					fail = true
				} else {
					logger.Info("Mapper: valid")
				}
			}

			if len(cc.constellationFilePath) > 0 {
				d, err = loadAndValidateDag(cc.constellationFilePath)
				if err != nil {
					logger.Errorf("Constellation: %v", err)
					fail = true
				} else {
					logger.Info("Constellation: valid")
				}
			}

			if !d.IsEmpty() && !m.IsEmpty() {
				err = validateDagAndMapper(&d, &m)
				if err != nil {
					logger.Errorf("Deployment: %v", err)
					fail = true
				} else {
					logger.Info("Deployment: valid")
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

func validateDagAndMapper(d *constellation.Config, m *mapper.Config) (err error) {
	types := []string{}
	mapTypes := []string{}

	for _, i := range d.Services {
		_, exists := helpers.Find(types, i.Type)
		if !exists {
			types = append(types, i.Type)
		}
	}

	for _, i := range m.Maps {
		mapTypes = append(mapTypes, i.Type)
	}

	logger.Debug("deployment: checking if `Service` exists in map")
	for _, i := range types {
		_, exists := helpers.Find(mapTypes, i)
		if !exists {
			logger.Error("Missing map configuration(s)")
			logger.Errorf("Service `%v` does not exist in map", i)
			err = fmt.Errorf("invalid")
		}
	}

	return
}

func loadAndValidateDag(path string) (config constellation.Config, err error) {
	err = config.LoadFile(path)

	if err != nil {
		return
	}

	return config, validateDag(&config)
}

// validateDag takes a constellation dag and returns any errors.
func validateDag(d *constellation.Config) (err error) {
	logger.Debug("Constellation: validating schema")
	err = d.ValidateModel()

	if err != nil {
		logger.Debug(err)
		return fmt.Errorf("invalid schema")
	}

	logger.Debug("constellation: checking for duplicate `ID`")
	duplicates := d.DuplicateIDs()

	if duplicates != nil {
		logger.Error("Duplicate `ID` present in config")
		for _, i := range duplicates {
			logger.Errorf("'%v'", i)
		}
		err = fmt.Errorf("invalid")
	}

	logger.Debug("Constellation: checking if `Service` exists")
	connections := d.CheckServiceExists()

	if len(connections) > 0 {
		logger.Error("Missing relationship(s)")
		for key, i := range connections {
			logger.Errorf("Relationship '%v' has missing `Services`:", key)
			for _, j := range i {
				logger.Errorf("'%v'", j)
			}
		}
		err = fmt.Errorf("invalid")
	}

	return
}

func loadAndValidateMapper(path string) (config mapper.Config, err error) {
	err = config.LoadFile(path)

	if err != nil {
		return
	}

	return config, validateMapper(&config)
}

// validateMapper takes a constellation mapper and returns any errors.
func validateMapper(m *mapper.Config) (err error) {
	logger.Debug("Mapper: validating schema")
	err = m.ValidateModel()

	if err != nil {
		logger.Debug(err)
		return fmt.Errorf("invalid schema")
	}

	logger.Debug("Mapper: checking for duplicate `ChartName`")
	duplicates := m.DuplicateChartName()

	if duplicates != nil {
		logger.Error("Duplicate `ChartName` present in config")
		for _, i := range duplicates {
			logger.Errorf("'%v'", i)
		}
		err = fmt.Errorf("invalid")
	}

	logger.Debug("Mapper: checking for duplicate `Type`")
	duplicates = m.DuplicateType()

	if duplicates != nil {
		logger.Error("Duplicate `Type` present in config")
		for _, i := range duplicates {
			logger.Errorf("'%v'", i)
		}
		err = fmt.Errorf("invalid")
	}

	logger.Debug("Mapper: checking for duplicate `Location`")
	duplicates = m.DuplicateLocation()

	if duplicates != nil {
		logger.Error("Duplicate `Location` present in config")
		for _, i := range duplicates {
			logger.Errorf("'%v'", i)
		}
		err = fmt.Errorf("invalid")
	}

	return
}
