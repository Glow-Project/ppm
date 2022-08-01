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
	if len(dependencies.First()) == 0 {
		uninstallAllDependencies(&config, paths, ctx.Bool("hard"))
		return nil
	}

	for i := 0; i < dependencies.Len(); i++ {
		dep := dependencies.Get(i)
		if config.HasDependency(dep) && !config.HasSubDependency(dep) {
			uninstallDependency(&config, paths, dep, false)
		} else if config.HasSubDependency(dep) {
			fmt.Println(color.RedString("the plugin"), color.YellowString(dep), color.RedString("is a sub dependency and can only be uninstalled by uninstalling its parent"))
		} else {
			fmt.Println(color.RedString("the plugin"), color.YellowString(dep), color.RedString("is not installed"))
		}
	}

	return nil
}

func uninstallAllDependencies(config *utility.PpmConfig, paths utility.Paths, hard bool) error {
	loadAnim := utility.StartLoading()

	err := os.RemoveAll(paths.Addons)
	if err != nil {
		return err
	}

	if hard {
		config.RemoveAllDependencies()
	}

	loadAnim.Stop()
	utility.PrintDone()
	return nil
}

func uninstallDependency(config *utility.PpmConfig, paths utility.Paths, dependency string, isSubDependency bool) error {
	dep := utility.GetPluginIdentifier(dependency)
	if !isSubDependency {
		fmt.Println("\runinstalling", color.YellowString(dep))
	} else {
		fmt.Println("\t -> uninstalling", color.YellowString(dep))
	}
	loadAnim := utility.StartLoading()

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
	fmt.Println(path.Join(paths.Addons, utility.GetPluginName(dep)))
	err = os.RemoveAll(path.Join(paths.Addons, utility.GetPluginName(dep)))
	loadAnim.Stop()
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
