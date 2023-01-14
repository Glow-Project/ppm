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
	updateAllDependencies(&config, &paths)
	loadAnim.Stop()

	utility.PrintDone()

	return nil
}

func updateAllDependencies(config *utility.PpmConfig, paths *utility.Paths) error {
	for _, dependency := range config.Dependencies {
		if dependency.Type != utility.GithubAsset {
			continue
		}

		err := utility.UpdateGithubRepo(filepath.Join(paths.Addons, dependency.Identifier))
		if err != nil {
			return err
		}
	}
	return nil
}
