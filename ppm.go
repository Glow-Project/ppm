package main

import (
	"log"
	"os"

	commands "github.com/Glow-Project/ppm/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "ppm",
		Usage:                "Pour Entertainment package manager for Godot",
		EnableBashCompletion: true,
		Commands:             commands.Commands(),
		Version:              "v1.0.1",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
