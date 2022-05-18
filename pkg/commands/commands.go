package commands

import "github.com/urfave/cli/v2"

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install a certain plugin or dependencies",
			Action:  install,
		},
		{
			Name:   "uninstall",
			Usage:  "uninstall a certain plugin or dependencies",
			Action: uninstall,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "hard",
					Value: false,
					Usage: "remove dependencies from config - Can't be un-done",
				},
			},
		},
		{
			Name:   "init",
			Usage:  "initialize a ppm.json file",
			Action: initialize,
		},
		{
			Name:   "update",
			Usage:  "update all dependencies",
			Action: update,
		},
		{
			Name:   "show",
			Usage:  "show the config",
			Action: showConfig,
		},
	}
}
