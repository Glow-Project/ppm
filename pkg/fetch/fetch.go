package fetch

import (
	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/go-git/go-git/v5"
)

func InstallDependency(dep *utility.Dependency, paths *utility.Paths) {
	if dep.Type == utility.GithubAsset {
		git.PlainClone(paths.Addons, false, &git.CloneOptions{
			URL: dep.Url,
		})
	} else {
		// TODO: Install Godot Assets Addon
	}
}
