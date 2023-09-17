package commands

import (
	"path/filepath"

	"github.com/Glow-Project/ppm/internal/paths"
	"github.com/Glow-Project/ppm/internal/pm"
	"github.com/Glow-Project/ppm/internal/utility"
	"github.com/urfave/cli/v2"
)

// the update method is called by the cli
// it updates all dependencies
func update(ctx *cli.Context) error {
	paths, config, err := pm.GetPathsAndConfig()
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

func updateAllDependencies(config *pm.Config, pth *paths.Paths) error {
	for _, dependency := range config.Dependencies {
		if dependency.Type != pm.GithubAsset {
			continue
		}

		if err := utility.UpdateGithubRepo(filepath.Join(pth.Addons, dependency.Identifier)); err != nil {
			return err
		}
	}
	return nil
}
