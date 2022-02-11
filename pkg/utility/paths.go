package utility

import (
	"os"
	"path/filepath"
	"strings"
)

// Access different important paths in the GD project
type Paths struct {
	/*
		The root of the project.

		If the project is a game, the root is on the same level
		as the godot config files.

		If the project is a plugin, the root is somewhere under /addons
	*/
	Root string

	// The path to the addons folder
	Addons string

	// The path to the ppm.json config file
	ConfigFile string
}

// Create a new paths instance from the root path
func CreatePaths(rootPath string) Paths {
	var addons string
	if strings.HasSuffix(filepath.Dir(rootPath), "addons") {
		addons = filepath.Dir(rootPath)
	} else {
		addons = filepath.Join(rootPath, "addons")
	}

	return Paths{
		Root:       rootPath,
		Addons:     addons,
		ConfigFile: filepath.Join(rootPath, "ppm.json"),
	}
}

// Create a new paths instance from the current working directory
func CreatePathsFromCwd() (Paths, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return Paths{}, err
	}

	return CreatePaths(rootPath), nil
}
