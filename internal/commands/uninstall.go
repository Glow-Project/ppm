package commands

import (
	"os"
	"path"

	"github.com/Glow-Project/ppm/internal/utility"
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
		dep := utility.DependencyFromString(dependencies.Get(i))
		if config.HasDependency(dep) && !config.HasSubDependency(dep) {
			uninstallDependency(&config, paths, dep, false)
		} else if config.HasSubDependency(dep) {
			utility.ColorPrintln("{RED}the plugin {YLW}%s {RED}is a sub-dependency and can only be uninstalled by uninstalling its parent", dep.Identifier)
		} else {
			utility.ColorPrintln("{RED}the plugin %s is not installed", dep.Identifier)
		}
	}

	return nil
}

func uninstallAllDependencies(config *utility.PpmConfig, paths utility.Paths, hard bool) error {
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

func uninstallDependency(config *utility.PpmConfig, paths utility.Paths, dependency utility.Dependency, isSubDependency bool) error {
	if !isSubDependency {
		utility.ColorPrintln("\runinstalling {YLW}%s", dependency.Identifier)
	} else {
		utility.ColorPrintln("\t -> uninstalling {YLW}%s", dependency.Identifier)
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
