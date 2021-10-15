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
		if config.HasDependency(repo) {
			alreadyInstalled(repo)
			return nil
		}
		installDependency(config, newPath, repo, false)
	} else {
		installAllDependencies(config, newPath)
	}

	return nil
}

func installAllDependencies(config pkg.PpmConfig, currentPath string) error {
	loading := true
	go pkg.PlayLoadingAnim(&loading)
	for _, dependency := range config.Dependencies {
		fmt.Printf("\rInstalling %s\n", color.YellowString(dependency))
		dependency, version := pkg.GetVersionOrNot(dependency)
		err := pkg.Clone(currentPath, dependency, version)
		if err != nil {
			installError(dependency)
			return err
		}
	}
	pkg.PrintDone()
	loading = false
	return nil
}

func installDependency(config pkg.PpmConfig, currentPath string, dependency string, isSubdependency bool) error {
	dependency, version := pkg.GetVersionOrNot(dependency)
	fmt.Printf("Installing %s\n", color.YellowString(dependency))
	loading := true

	go pkg.PlayLoadingAnim(&loading)
	err := pkg.Clone(currentPath, dependency, version)
	loading = false

	var addDependency bool

	if err != nil {
		installError(dependency)
		return err
	} else {
		pkg.PrintDone()
		addDependency = true
	}
	
	
	if addDependency && isSubdependency {
		config.AddSubDependency(dependency)
	} else if addDependency {
		config.AddDependency(dependency)
	}
	
	
	subConfig, err := pkg.GetPluginConfig(currentPath, dependency)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	// Iterate over dependencies and install them if needed
	for _, dep := range subConfig.Dependencies {
		if !config.HasDependency(dep) && !config.HasSubDependency(dep) {
			installDependency(config, currentPath, dep, true)
		}
	}

	return nil
}

func alreadyInstalled(dependency string) {
	fmt.Println(color.GreenString("The Plugin"), color.YellowString(dependency), color.GreenString("is already installed"))
}

func installError(dependency string) {
	fmt.Printf(color.RedString("\rSome issues occured while trying to install %s"), color.YellowString(dependency))
}