package commands

import (
	"encoding/json"
	"os"

	"github.com/Glow-Project/ppm/internal/utility"
	"github.com/urfave/cli/v2"
)

func tidy(ctx *cli.Context) error {
	paths, err := utility.CreatePathsFromCwd()
	if err != nil {
		return err
	}

	content, err := os.ReadFile(paths.ConfigFile)
	if err != nil {
		return err
	}

	var jsonContent map[string]interface{}
	if err := json.Unmarshal(content, &jsonContent); err != nil {
		return err
	}

	strDeps, ok := jsonContent["dependencies"].([]interface{})
	if !ok {
		return nil
	}

	deps := []utility.Dependency{}
	for _, dep := range strDeps {
		str, ok := dep.(string)
		if !ok {
			return nil
		}
		deps = append(deps, utility.DependencyFromString(str))
	}

	jsonContent["dependencies"] = deps
	jsonData, err := json.MarshalIndent(jsonContent, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(paths.ConfigFile, jsonData, 0644)

	config, err := utility.ParsePpmConfig(paths.ConfigFile)
	if err != nil {
		return err
	}

	return config.Write()
}
