package main

import (
	"log"
	"os"

	"github.com/Glow-Project/ppm/pkg/commands"
	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

var BuildVersion string

func main() {
	app := &cli.App{
		Name:                 "ppm",
		Version:              BuildVersion,
		Usage:                "Pour Entertainment package manager for Godot",
		EnableBashCompletion: true,
		Commands:             commands.Commands(),
		HideVersion:          false,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "no-color",
				Required: false,
				Action: func(ctx *cli.Context, b bool) error {
					if b {
						utility.UseColor = false
					}

					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
