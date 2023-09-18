package commands

import (
	"encoding/json"
	"os"

	"github.com/Glow-Project/ppm/internal/paths"
	"github.com/Glow-Project/ppm/internal/pm"

	"github.com/urfave/cli/v2"
)

func tidy(ctx *cli.Context) error {
	pth, err := paths.CreatePathsFromCwd()
	if err != nil {
		return err
	}

	content, err := os.ReadFile(pth.ConfigFile)
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

	deps := []pm.Dependency{}
	for _, dep := range strDeps {
		str, ok := dep.(string)
		if !ok {
			return nil
		}
		deps = append(deps, pm.DependencyFromString(str))
	}

	jsonContent["dependencies"] = deps
	jsonData, err := json.MarshalIndent(jsonContent, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(pth.ConfigFile, jsonData, 0644)

	config, err := pm.ParseConfig(pth.ConfigFile)
	if err != nil {
		return err
	}

	return config.Write()
}
