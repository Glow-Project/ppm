package commands

import (
	"errors"
	"fmt"
	"path"

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

	fmt.Println(color.GreenString("new ppm.json config-file generated"))

	if ok, _ := utility.DoesPathExist(path.Join(paths.Root, ".git")); ok {
		fmt.Println(color.YellowString("when using ppm it is recommended to add the addons directory to your .gitignore file"))
	}

	return nil
}
