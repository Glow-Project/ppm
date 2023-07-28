package commands

import (
	"fmt"
	"path"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

func initialize(ctx *cli.Context) error {
	paths, err := utility.CreatePathsFromCwd()
	if err != nil {
		return err
	}

	if _, err := utility.CreatePpmConfig(paths.Root); err != nil {
		return fmt.Errorf("error creating ppm.json config-file: %w", err)
	}

	utility.ColorPrintln("{GRN}new ppm.json config-file generated")

	if ok, _ := utility.DoesPathExist(path.Join(paths.Root, ".git")); ok {
		utility.ColorPrintln("{YLW}when using ppm it is recommended to add the addons direcotry to your .gitignore file")
	}

	return nil
}
