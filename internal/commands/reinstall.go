package commands

import (
	"github.com/Glow-Project/ppm/internal/pm"
	"github.com/urfave/cli/v2"
)

func reinstall(ctx *cli.Context) error {
	paths, config, err := pm.GetPathsAndConfig()
	if err != nil {
		return err
	}

	if err := uninstallAllDependencies(&config, paths, false); err != nil {
		return err
	}

	return installAllDependencies(&config, paths)
}
