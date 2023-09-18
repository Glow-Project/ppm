package commands

import (
	"os"
	"path"

	"github.com/Glow-Project/ppm/internal/paths"
	"github.com/Glow-Project/ppm/internal/pm"
	"github.com/Glow-Project/ppm/internal/utility"
	"github.com/urfave/cli/v2"
)

func uninstall(ctx *cli.Context) error {
	pth, config, err := pm.GetPathsAndConfig()
	if err != nil {
		return err
	}

	dependencies := ctx.Args()
	if len(dependencies.First()) == 0 {
		uninstallAllDependencies(&config, pth, ctx.Bool("hard"))
		return nil
	}

	for i := 0; i < dependencies.Len(); i++ {
		dep := pm.DependencyFromString(dependencies.Get(i))
		if config.HasDependency(dep) && !config.HasSubDependency(dep) {
			uninstallDependency(&config, pth, dep, false)
		} else if config.HasSubDependency(dep) {
			utility.ColorPrintln("{RED}the plugin {YLW}%s {RED}is a sub-dependency and can only be uninstalled by uninstalling its parent", dep.Identifier)
		} else {
			utility.ColorPrintln("{RED}the plugin %s is not installed", dep.Identifier)
		}
	}

	return nil
}

func uninstallAllDependencies(config *pm.Config, pth paths.Paths, hard bool) error {
	loadAnim := utility.StartLoading()

	if err := os.RemoveAll(pth.Addons); err != nil {
		return err
	}

	if hard {
		config.RemoveAllDependencies()
	}

	loadAnim.Stop()
	utility.PrintDone()
	return nil
}

func uninstallDependency(config *pm.Config, pth paths.Paths, dependency pm.Dependency, isSubDependency bool) error {
	if !isSubDependency {
		utility.ColorPrintln("\runinstalling {YLW}%s", dependency.Identifier)
	} else {
		utility.ColorPrintln("\t -> uninstalling {YLW}%s", dependency.Identifier)
	}
	loadAnim := utility.StartLoading()

	subConfig, err := pm.GetPluginConfig(pth, dependency)
	for i := 0; err == nil && i < len(subConfig.Dependencies); i++ {
		dep := subConfig.Dependencies[i]
		if !config.HasDependency(dep) {
			uninstallDependency(config, pth, dep, true)
		}
	}

	// path: root/addons/dependency
	err = os.RemoveAll(path.Join(pth.Addons, dependency.Identifier))
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
