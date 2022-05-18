package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func uninstall(ctx *cli.Context) error {
	paths, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	dependencies := ctx.Args()
	if len(dependencies.First()) > 0 {
		for i := 0; i < dependencies.Len(); i++ {
			dep := dependencies.Get(i)
			if !config.HasDependency(dep) && !config.HasSubDependency(dep) {
				fmt.Println(color.RedString("the plugin"), color.YellowString(dep), color.RedString("is not installed"))
			} else {
				uninstallDependency(&config, paths, dep, false)
			}

		}

	} else {
		uninstallAllDependencies(config, paths, ctx.Bool("hard"))
	}

	return nil
}

func uninstallAllDependencies(config utility.PpmConfig, paths utility.Paths, hard bool) error {
	loading := make(chan interface{}, 1)
	go utility.PlayLoadingAnim(loading)

	// path: root/addons
	err := os.RemoveAll(paths.Addons)
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

func uninstallDependency(config *utility.PpmConfig, paths utility.Paths, dependency string, isSubDependency bool) error {
	dep := utility.GetPluginName(dependency)
	if !isSubDependency {
		fmt.Println("\runinstalling", color.YellowString(dep))
	} else {
		fmt.Println("\t -> uninstalling", color.YellowString(dep))
	}
	loading := make(chan interface{}, 1)
	go utility.PlayLoadingAnim(loading)

	subConfig, err := utility.GetPluginConfig(paths.Addons, dep)
	if err == nil {
		for i := 0; i < len(subConfig.Dependencies); i++ {
			subDep := subConfig.Dependencies[i]
			if !config.HasDependency(subDep) {
				uninstallDependency(config, paths, subDep, true)
			}
		}
	}

	// path: root/addons/dependency
	err = os.RemoveAll(path.Join(paths.Addons, dep))
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