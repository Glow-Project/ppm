package main

import (
	"log"
	"os"

	"github.com/Glow-Project/ppm/pkg/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "ppm",
		Usage:                "Pour Entertainment package manager for Godot",
		EnableBashCompletion: true,
		Commands:             commands.Commands(),
		Version:              "v1.1.0",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
