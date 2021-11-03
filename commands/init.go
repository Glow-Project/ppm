package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/Glow-Project/ppm/pkg"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func initialize(ctx *cli.Context) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	err = pkg.CreateNewPpmConfig(currentPath)
	if err != nil {
		return errors.New("ppm.json config-file already exists in this directory")
	}

	fmt.Println(color.GreenString("New ppm.json config-file generated"))

	return nil
}
