package commands

import (
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
		fmt.Println(color.RedString("ppm.json config-file already exists in this directory"))
		return nil
	}

	fmt.Println(color.GreenString("New ppm.json config-file generated"))

	return nil
}