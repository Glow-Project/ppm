package commands

import (
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

// The update method is called by the cli
// It updates all dependencies
func update(ctx *cli.Context) error {
	paths, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	utility.CheckOrCreateDir(paths.Addons)

	loading := make(chan interface{}, 1)
	go utility.PlayLoadingAnim(loading)
	updateAllDependencies(config, paths)
	loading <- nil

	utility.PrintDone()

	return nil
}

func updateAllDependencies(config utility.PpmConfig, paths utility.Paths) error {
	for _, dependency := range config.Dependencies {
		_, version := utility.GetVersionOrNot(dependency)
		if len(version) > 0 {
			continue
		}

		err := utility.Update(filepath.Join(paths.Addons, dependency))
		if err != nil {
			return err
		}
	}
	return nil
}
