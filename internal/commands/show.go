package commands

import (
	"github.com/Glow-Project/ppm/internal/pm"
	"github.com/urfave/cli/v2"
)

func showConfig(ctx *cli.Context) error {
	_, config, err := pm.GetPathsAndConfig()
	if err != nil {
		return err
	}

	config.PrettyPrint()

	return nil
}
