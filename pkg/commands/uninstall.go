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
		uninstallAllDependencies(&config, &paths, ctx.Bool("hard"))
		return nil
	}

	for i := 0; i < dependencies.Len(); i++ {
		dep := utility.DependencyFromString(dependencies.Get(i))
		if config.HasDependency(dep) && !config.HasSubDependency(dep) {
			uninstallDependency(&config, &paths, dep, false)
		} else if config.HasSubDependency(dep) {
			fmt.Println(color.RedString("the plugin"), color.YellowString(dep.Identifier), color.RedString("is a sub dependency and can only be uninstalled by uninstalling its parent"))
		} else {
			fmt.Println(color.RedString("the plugin"), color.YellowString(dep.Identifier), color.RedString("is not installed"))
		}
	}

	return nil
}

func uninstallAllDependencies(config *utility.PpmConfig, paths *utility.Paths, hard bool) error {
	loadAnim := utility.StartLoading()

	if err := os.RemoveAll(paths.Addons); err != nil {
		return err
	}

	if hard {
		config.RemoveAllDependencies()
	}

	loadAnim.Stop()
	utility.PrintDone()
	return nil
}

func uninstallDependency(config *utility.PpmConfig, paths *utility.Paths, dependency *utility.Dependency, isSubDependency bool) error {
	if !isSubDependency {
		fmt.Println("\runinstalling", color.YellowString(dependency.Identifier))
	} else {
		fmt.Println("\t -> uninstalling", color.YellowString(dependency.Identifier))
	}
	loadAnim := utility.StartLoading()

	subConfig, err := utility.GetPluginConfig(paths, dependency)
	for i := 0; err == nil && i < len(subConfig.Dependencies); i++ {
		dep := subConfig.Dependencies[i]
		if !config.HasDependency(dep) {
			uninstallDependency(config, paths, dep, true)
		}
	}

	// path: root/addons/dependency
	err = os.RemoveAll(path.Join(paths.Addons, dependency.Identifier))
	loadAnim.Stop()
	if err != nil {
		return err
	}

	if isSubDependency {
		config.RemoveSubDependency(dependency)
		return nil
	}

	config.RemoveDependency(dependency)
	utility.PrintDone()
	return nil
}
