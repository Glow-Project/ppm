package commands

import (
	"path/filepath"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

// the update method is called by the cli
// it updates all dependencies
func update(ctx *cli.Context) error {
	paths, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	utility.CheckOrCreateDir(paths.Addons)

	loadAnim := utility.StartLoading()
	updateAllDependencies(config, paths)
	loadAnim.Stop()

	utility.PrintDone()

	return nil
}

func updateAllDependencies(config utility.PpmConfig, paths utility.Paths) error {
	for _, dependency := range config.Dependencies {
		_, version := utility.GetVersionOrNot(dependency.Identifier)
		if len(version) > 0 {
			continue
		}

		err := utility.Update(filepath.Join(paths.Addons, dependency.Identifier))
		if err != nil {
			return err
		}
	}
	return nil
}
