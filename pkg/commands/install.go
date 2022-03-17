package commands

import (
	"fmt"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func install(ctx *cli.Context) error {
	paths, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	utility.CheckOrCreateDir(paths.Addons)

	dependencies := ctx.Args()
	if dependencies.Len() > 0 {
		for i := 0; i < dependencies.Len(); i++ {
			repo := dependencies.Get(i)

			if config.HasDependency(repo) {
				alreadyInstalled(repo)
			} else {
				if err = installDependency(config, paths, repo, false); err != nil {
					return err
				}
			}
		}
	} else {
		installAllDependencies(config, paths)
	}

	return nil
}

func installAllDependencies(config utility.PpmConfig, paths utility.Paths) error {
	for _, dependency := range config.Dependencies {
		if err := installDependency(config, paths, dependency, false); err != nil {
			return err
		}
	}
	utility.PrintDone()
	return nil
}

func installDependency(config utility.PpmConfig, paths utility.Paths, dependency string, isSubDependency bool) error {
	dependency, version := utility.GetVersionOrNot(dependency)
	if !isSubDependency {
		fmt.Printf("\rinstalling %s\n", color.YellowString(utility.GetPluginName(dependency)))
	} else {
		fmt.Printf("\r\t -> installing %s\n", color.YellowString(utility.GetPluginName(dependency)))
	}
	loading := make(chan interface{}, 1)

	go utility.PlayLoadingAnim(loading)
	err := utility.Clone(paths.Addons, dependency, version)
	loading <- nil

	var addDependency bool

	if err != nil {
		if err.Error() == "repository already exists" {
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

	subConfig, err := utility.GetPluginConfig(paths.Addons, dependency)
	if err != nil {
		if !isSubDependency {
			utility.PrintDone()
		}
		return nil
	}

	// Iterate over dependencies and install them if needed
	for _, dep := range subConfig.Dependencies {
		if !config.HasDependency(dep) && !config.HasSubDependency(dep) {
			installDependency(config, paths, dep, true)
		}
	}

	if !isSubDependency {
		utility.PrintDone()
	}

	return nil
}

func alreadyInstalled(dependency string) {
	fmt.Println(color.GreenString("\rthe plugin"), color.YellowString(dependency), color.GreenString("is already installed"))
}

func installError(dependency string) {
	fmt.Printf(color.RedString("\rsome issues occured while trying to install %s, %s"), color.YellowString(dependency), color.RedString("are you sure you spelled it right?\n"))
}
