package commands

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg"
	"github.com/urfave/cli/v2"
)

func showConfig(ctx *cli.Context) error {
	// currentPath is the root directory of the project
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	config, err := pkg.ParsePpmConfig(filepath.Join(currentPath, "ppm.json"))
	if err != nil {
		return errors.New("could not find ppm.json file - try to run: ppm init")
	}

	config.PrettyPrint()

	return nil
}