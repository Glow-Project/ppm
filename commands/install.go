package commands

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg"
	"github.com/urfave/cli/v2"
)

func install(ctx *cli.Context) error {
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

	repo := ctx.Args().Get(0)
	if len(repo) > 0 {
		installDependency(config, newPath, repo)
	} else {
		installAllDependencies(config, newPath)
	}

	return nil
}

func installAllDependencies(config pkg.PpmConfig, currentPath string) error {
	for _, dependency := range config.Dependencies {
		dependency, version := pkg.GetVersionOrNot(dependency)
		err := pkg.Clone(currentPath, dependency, version)
		if err != nil {
			return err
		}
	}
	return nil
}

func installDependency(config pkg.PpmConfig, currentPath string, dependency string) error {
	dependency, version := pkg.GetVersionOrNot(dependency)
	
	err := pkg.Clone(currentPath, dependency, version)
	if err != nil {
		return err
	}
	
	config.AddDependency(dependency)
	
	return nil
}