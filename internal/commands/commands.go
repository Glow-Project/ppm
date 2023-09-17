package commands

import "github.com/urfave/cli/v2"

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install one or more dependencies",
			Action:  install,
		},
		{
			Name:    "uninstall",
			Aliases: []string{"u"},
			Usage:   "uninstall one or more dependencies",
			Action:  uninstall,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "hard",
					Value: false,
					Usage: "remove dependencies from config - Can't be un-done",
				},
			},
		},
		{
			Name:   "reinstall",
			Usage:  "delete and reinstall all dependencies",
			Action: reinstall,
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
		{
			Name:   "tidy",
			Usage:  "clean up the config",
			Action: tidy,
		},
	}
}
