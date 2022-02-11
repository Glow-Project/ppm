package commands

import (
	"errors"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

func showConfig(ctx *cli.Context) error {
	paths, err := utility.CreatePathsFromCwd()
	if err != nil {
		return err
	}

	config, err := utility.ParsePpmConfig(paths.ConfigFile)
	if err != nil {
		return errors.New("could not find ppm.json file - try to run: ppm init")
	}

	config.PrettyPrint()

	return nil
}
