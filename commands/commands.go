package commands

import "github.com/urfave/cli/v2"

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name: "install",
			Aliases: []string{"i"},
			Usage: "Install a certain plugin or dependencies",
			Action: install,
		},
		{
			Name: "uninstall",
			Usage: "Uninstall a certain plugin or dependencies",
			Action: uninstall,
		},
		{
			Name: "init",
			Usage: "Initialize a ppm.json file",
			Action: initialize,
		},
		{
			Name: "update",
			Usage: "Update all dependencies",
			Action: update,
		},
	}
}