package commands

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg"
	"github.com/urfave/cli/v2"
)

// The update method is called by the cli
// It updates all dependencies
func update(ctx *cli.Context) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	config, err := pkg.ParsePpmConfig(filepath.Join(currentPath, "ppm.json"))
	if err != nil {
		return errors.New("could not find ppm.json file - try to run: ppm init")
	}

	if config.Plugin {
		os.Chdir(filepath.Dir(currentPath))
	} else {
		addonPath := filepath.Join(currentPath, "addons")
		pathExists, _ := pkg.DoesPathExist(addonPath)
		if !pathExists {
			err := os.Mkdir("addons", 0755)
			if err != nil {
				return err
			}
		}
		os.Chdir(addonPath)
	}

	newPath, err := os.Getwd()
	if err != nil {
		return err
	}

	updateAllDependencies(config, newPath)

	return nil
}

func updateAllDependencies(config pkg.PpmConfig, currentPath string) error {
	for _, dependency := range config.Dependencies {
		err := pkg.Update(filepath.Join(currentPath, dependency))
		if err != nil {
			return err
		}
	}
	return nil
}
