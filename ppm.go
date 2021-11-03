package main

import (
	"log"
	"os"

	commands "github.com/Glow-Project/ppm/commands"
	"github.com/urfave/cli/v2"
)


func main(){
	app := &cli.App{
		Name: "ppm",
		Usage: "Pour Entertainment package manager for Godot",
		EnableBashCompletion: true,
		Commands: commands.Commands(),
		Version: "v1.0.1",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}