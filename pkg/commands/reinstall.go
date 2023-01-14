package commands

import (
	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

func reinstall(ctx *cli.Context) error {
	paths, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	err = uninstallAllDependencies(&config, &paths, false)
	if err != nil {
		return err
	}

	err = installAllDependencies(&config, &paths)
	if err != nil {
		return err
	}

	return nil
}
