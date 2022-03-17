package commands

import (
	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

func showConfig(ctx *cli.Context) error {
	_, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	config.PrettyPrint()

	return nil
}
