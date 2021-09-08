package commands

import (
	"os"

	"github.com/Glow-Project/ppm/pkg"
	"github.com/urfave/cli/v2"
)

func initialize(ctx *cli.Context) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	innerErr := pkg.CreateNewPpmConfig(currentPath)
	if innerErr != nil {
		return err
	}

	return nil
}