package commands

import (
	"errors"
	"fmt"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func initialize(ctx *cli.Context) error {
	paths, err := utility.CreatePathsFromCwd()
	if err != nil {
		return err
	}

	err = utility.CreateNewPpmConfig(paths.Root)
	if err != nil {
		return errors.New("ppm.json config-file already exists in this directory")
	}

	fmt.Println(color.GreenString("New ppm.json config-file generated"))

	return nil
}
