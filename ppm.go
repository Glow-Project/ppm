package main

import (
	"log"
	"os"

	"github.com/Glow-Project/ppm/pkg/commands"
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
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
