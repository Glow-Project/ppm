package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func install(ctx *cli.Context) error {
	// currentPath is the root directory of the project
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

	dependencies := ctx.Args()
	if dependencies.Len() > 0 {
		for i := 0; i < dependencies.Len(); i++ {
			repo := dependencies.Get(i)

			if config.HasDependency(repo) {
				alreadyInstalled(repo)
			} else {
				installDependency(config, newPath, repo, false)
			}
		}
	} else {
		installAllDependencies(config, newPath)
	}

	return nil
}

func installAllDependencies(config pkg.PpmConfig, currentPath string) error {
	for _, dependency := range config.Dependencies {
		installDependency(config, currentPath, dependency, false)
	}
	pkg.PrintDone()
	return nil
}

func installDependency(config pkg.PpmConfig, currentPath string, dependency string, isSubDependency bool) error {
	dependency, version := pkg.GetVersionOrNot(dependency)
	if !isSubDependency {
		fmt.Printf("\rinstalling %s\n", color.YellowString(pkg.GetPluginName(dependency)))
	} else {
		fmt.Printf("\r\t -> installing %s\n", color.YellowString(pkg.GetPluginName(dependency)))
	}
	loading := make(chan interface{}, 1)

	go pkg.PlayLoadingAnim(loading)
	err := pkg.Clone(currentPath, dependency, version)
	loading <- nil

	var addDependency bool

	if err != nil {
		if err.Error() == "exit status 128" {
			alreadyInstalled(dependency)
			return nil
		} else {
			installError(dependency)
			return err
		}
	} else if !config.HasDependency(dependency) && !config.HasSubDependency(dependency) {
		addDependency = true
	}

	if addDependency && isSubDependency {
		config.AddSubDependency(dependency)
	} else if addDependency {
		config.AddDependency(dependency)
	}

	subConfig, err := pkg.GetPluginConfig(currentPath, dependency)
	if err != nil {
		if !isSubDependency {
			pkg.PrintDone()
		}
		return nil
	}

	// Iterate over dependencies and install them if needed
	for _, dep := range subConfig.Dependencies {
		if !config.HasDependency(dep) && !config.HasSubDependency(dep) {
			installDependency(config, currentPath, dep, true)
		}
	}

	if !isSubDependency {
		pkg.PrintDone()
	}

	return nil
}

func alreadyInstalled(dependency string) {
	fmt.Println(color.GreenString("\rthe plugin"), color.YellowString(dependency), color.GreenString("is already installed"))
}

func installError(dependency string) {
	fmt.Printf(color.RedString("\rsome issues occured while trying to install %s\n"), color.YellowString(dependency))
}
