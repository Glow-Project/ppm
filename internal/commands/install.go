package commands

import (
	"github.com/Glow-Project/ppm/internal/fetch"
	"github.com/Glow-Project/ppm/internal/paths"
	"github.com/Glow-Project/ppm/internal/pm"
	"github.com/Glow-Project/ppm/internal/utility"
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

func install(ctx *cli.Context) error {
	pth, config, err := pm.GetPathsAndConfig()
	if err != nil {
		return err
	}

	utility.CheckOrCreateDir(pth.Addons)

	dependencies := ctx.Args()
	if dependencies.Len() == 0 {
		installAllDependencies(&config, pth)
	}

	for _, dep := range dependencies.Slice() {
		if err = installDependency(&config, pth, pm.DependencyFromString(dep), false); err != nil {
			return err
		}
	}

	return nil
}

func installAllDependencies(config *pm.Config, pth paths.Paths) error {
	for _, dependency := range config.Dependencies {
		if err := installDependency(config, pth, dependency, false); err != nil {
			return err
		}
	}
	return nil
}

func installDependency(config *pm.Config, pth paths.Paths, dependency pm.Dependency, isSubDependency bool) error {
	if !isSubDependency {
		utility.ColorPrintln("\rinstalling {YLW}%s", dependency.Identifier)
	} else {
		utility.ColorPrintln("\t -> installing {YLW}%s", dependency.Identifier)
	}
	loadAnim := utility.StartLoading()

	err := fetch.InstallDependency(dependency, pth)
	loadAnim.Stop()

	switch err := err.(type) {
	case nil:
		break
	case *fetch.InvalidVersionError:
		dependency.Version = nil
		versionError(dependency.Identifier, err.Version)
	case *fetch.CloneError:
		gitErr := err.GitError
		if gitErr == git.ErrRepositoryAlreadyExists {
			alreadyInstalled(dependency.Identifier)
			return nil
		} else {
			installError(dependency.Identifier)
			// return err.GitError
			return nil
		}
	default:
		return err
	}

	shouldAddDep := (!isSubDependency && !config.HasDependency(dependency)) ||
		(isSubDependency && !config.HasSubDependency(dependency))

	if shouldAddDep && isSubDependency {
		config.AddSubDependency(dependency)
	} else if shouldAddDep {
		config.AddDependency(dependency)
	}

	subConfig, err := pm.GetPluginConfig(pth, dependency)
	if err != nil {
		if !isSubDependency {
			utility.PrintDone()
		}
		return nil
	}

	// iterate over dependencies and install them if needed
	for _, dep := range subConfig.Dependencies {
		if !config.HasSubDependency(dep) {
			installDependency(config, pth, dep, true)
		}
	}

	if !isSubDependency {
		utility.PrintDone()
	}

	return nil
}

func alreadyInstalled(dependency string) {
	utility.ColorPrintln("\r{GRN}the plugin {YLW}%s {GRN}is already installed", dependency)
}

func installError(dependency string) {
	utility.ColorPrintln("\r{RED}some issues occured while trying to install {YLW}%s {RED}are you sure you spelled it right?", dependency)
}

func versionError(dependency string, version string) {
	utility.ColorPrintln("\r{RED}the version {YLW}%s {RED} for the dependency {YLW}%s {RED}was not found. The default version was installed", version, dependency)
}
