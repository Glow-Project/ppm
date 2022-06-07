package commands

import (
	"fmt"
	"strings"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/urfave/cli/v2"
)

func tidy(ctx *cli.Context) error {
	_, config, err := utility.GetPathsAndConfig()
	if err != nil {
		return err
	}

	for i, dep := range config.Dependencies {
		if utility.IsGithubRepoUrl(dep) {
			tmp := strings.Split(dep, "/")
			config.Dependencies[i] = fmt.Sprintf("%s/%s", tmp[len(tmp)-2], tmp[len(tmp)-1])
		}
	}

	return config.Write()
}
