package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// about the naming: the "Ppm" prefix for the config was chosen,
// because godot project config files may be supported in the future

// representing a ppm.json configuration file
type PpmConfig struct {
	IsPlugin        bool         `json:"plugin"`
	Dependencies    []Dependency `json:"dependencies"`
	SubDependencies []Dependency `json:"sub-dependencies"`
	filePath        string
}

// add an item safely to the Dependencies property
func (ppm *PpmConfig) AddDependency(dependency Dependency) {
	ppm.Dependencies = append(ppm.Dependencies, dependency)
	ppm.Write()
}

// add an item safely to the sub-dependencies property
func (ppm *PpmConfig) AddSubDependency(dependency Dependency) {
	ppm.SubDependencies = append(ppm.SubDependencies, dependency)
	ppm.Write()
}

// remove ALL (sub)dependencies
func (ppm *PpmConfig) RemoveAllDependencies() {
	ppm.Dependencies = []Dependency{}
	ppm.SubDependencies = []Dependency{}
	ppm.Write()
}

// remove an item safely from the Dependencies property by its name
func (ppm *PpmConfig) RemoveSubDependency(dependency Dependency) {
	ppm.SubDependencies = Filter(ppm.SubDependencies, func(item Dependency, _ int) bool {
		return item.Identifier != dependency.Identifier
	})
	ppm.Write()
}

// remove an item safely from the sub-dependencies property by its name
func (ppm *PpmConfig) RemoveDependency(dependency Dependency) {
	ppm.Dependencies = Filter(ppm.Dependencies, func(item Dependency, _ int) bool {
		return item.Identifier != dependency.Identifier
	})
	ppm.Write()
}

// check wether the config file has a certain dependency
func (ppm PpmConfig) HasDependency(dependency Dependency) bool {
	return Some(ppm.Dependencies, func(dep Dependency, _ int) bool {
		return dep.Identifier == dependency.Identifier
	})
}

// check wether the config file has a certain sub-dependency
func (ppm PpmConfig) HasSubDependency(dependency Dependency) bool {
	return Some(ppm.SubDependencies, func(dep Dependency, _ int) bool {
		return dep.Identifier == dependency.Identifier
	})
}

func (ppm PpmConfig) PrettyPrint() {
	ppmType := "game"
	if ppm.IsPlugin {
		ppmType = "plugin"
	}

	depsToIdentifiers := func(deps []Dependency) []string {
		return Map(deps, func(item Dependency, _ int) string {
			return item.Identifier
		})
	}
	dependencies := SliceToString(depsToIdentifiers(ppm.Dependencies), ", ")
	subDependencies := SliceToString(depsToIdentifiers(ppm.SubDependencies), ", ")

	fmt.Printf("this project is a %s\ndependencies: %v\nsubdependencies: %v\n", ppmType, dependencies, subDependencies)
}

// write the current state of the configuartion to the config file
func (ppm PpmConfig) Write() error {
	content, err := json.MarshalIndent(ppm, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(ppm.filePath, content, 0755)
}

// parse the ppm.json file to an object
func ParsePpmConfig(filePath string) (PpmConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return PpmConfig{}, err
	}
	defer file.Close()

	content, err := os.ReadFile(file.Name())
	if err != nil {
		return PpmConfig{}, err
	}

	config := PpmConfig{}
	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		return PpmConfig{}, err
	}
	config.filePath = filePath

	return config, nil
}

// create a ppm.json file
func CreatePpmConfig(path string) (PpmConfig, error) {
	configPath := filepath.Join(path, "ppm.json")

	if fileExists, _ := DoesPathExist(configPath); fileExists {
		return PpmConfig{}, errors.New("file already exists")
	}

	config := PpmConfig{
		IsPlugin:        strings.HasSuffix(filepath.Dir(path), "addons"),
		Dependencies:    []Dependency{},
		SubDependencies: []Dependency{},
	}

	content, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return config, err
	}

	file, err := os.Create(configPath)
	if err != nil {
		return config, err
	}

	_, err = file.Write(content)

	return config, err
}

// get the config of a certain plugin
func GetPluginConfig(p Paths, dep Dependency) (PpmConfig, error) {
	config, err := ParsePpmConfig(path.Join(p.Addons, dep.Identifier, "ppm.json"))

	if err != nil {
		return PpmConfig{}, err
	}

	return config, nil
}
