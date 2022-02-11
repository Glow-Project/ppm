package commands

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func uninstall(ctx *cli.Context) error {
	// currentPath is the root directory of the project
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	config, err := utility.ParsePpmConfig(filepath.Join(currentPath, "ppm.json"))
	if err != nil {
		return errors.New("could not find ppm.json file - try to run: ppm init")
	}

	dependencies := ctx.Args()
	if len(dependencies.First()) > 0 {
		for i := 0; i < dependencies.Len(); i++ {
			dep := dependencies.Get(i)
			if !config.HasDependency(dep) && !config.HasSubDependency(dep) {
				fmt.Println(color.RedString("the plugin"), color.YellowString(dep), color.RedString("is not installed"))
			} else {
				uninstallDependency(&config, currentPath, dep, false)
			}

		}

	} else {
		uninstallAllDependencies(config, currentPath, ctx.Bool("hard"))
	}

	return nil
}

func uninstallAllDependencies(config utility.PpmConfig, currentPath string, hard bool) error {
	loading := make(chan interface{}, 1)
	go utility.PlayLoadingAnim(loading)

	// path: root/addons
	err := os.RemoveAll(path.Join(currentPath, "addons"))
	if err != nil {
		return err
	}

	if hard {
		config.RemoveAllDependencies()
	}

	loading <- nil
	utility.PrintDone()
	return nil
}

func uninstallDependency(config *utility.PpmConfig, currentPath string, dependency string, isSubDependency bool) error {
	dep := utility.GetPluginName(dependency)
	if !isSubDependency {
		fmt.Println("\runinstalling", color.YellowString(dep))
	} else {
		fmt.Println("\t -> uninstalling", color.YellowString(dep))
	}
	loading := make(chan interface{}, 1)
	go utility.PlayLoadingAnim(loading)

	subConfig, err := utility.GetPluginConfig(path.Join(currentPath, "addons"), dep)
	if err == nil {
		for i := 0; i < len(subConfig.Dependencies); i++ {
			subDep := subConfig.Dependencies[i]
			if !config.HasDependency(subDep) {
				uninstallDependency(config, currentPath, subDep, true)
			}
		}
	}

	// path: root/addons/dependency
	err = os.RemoveAll(path.Join(currentPath, "addons", dep))
	loading <- nil
	if err != nil {
		return err
	}

	if !isSubDependency {
		config.RemoveDependency(dep)
	} else {
		config.RemoveSubDependency(dep)
	}

	if !isSubDependency {
		utility.PrintDone()
	}
	return nil
}
