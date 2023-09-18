package pm

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Glow-Project/ppm/internal/paths"
	"github.com/Glow-Project/ppm/internal/utility"
)

// representing a ppm.json configuration file
type Config struct {
	IsPlugin        bool         `json:"plugin"`
	Dependencies    []Dependency `json:"dependencies"`
	SubDependencies []Dependency `json:"sub-dependencies"`
	filePath        string
}

// add an item safely to the Dependencies property
func (cfg *Config) AddDependency(dependency Dependency) {
	cfg.Dependencies = append(cfg.Dependencies, dependency)
	cfg.Write()
}

// add an item safely to the sub-dependencies property
func (cfg *Config) AddSubDependency(dependency Dependency) {
	cfg.SubDependencies = append(cfg.SubDependencies, dependency)
	cfg.Write()
}

// remove ALL (sub)dependencies
func (cfg *Config) RemoveAllDependencies() {
	cfg.Dependencies = []Dependency{}
	cfg.SubDependencies = []Dependency{}
	cfg.Write()
}

// remove an item safely from the Dependencies property by its name
func (cfg *Config) RemoveSubDependency(dependency Dependency) {
	cfg.SubDependencies = utility.Filter(cfg.SubDependencies, func(item Dependency, _ int) bool {
		return item.Identifier != dependency.Identifier
	})
	cfg.Write()
}

// remove an item safely from the sub-dependencies property by its name
func (cfg *Config) RemoveDependency(dependency Dependency) {
	cfg.Dependencies = utility.Filter(cfg.Dependencies, func(item Dependency, _ int) bool {
		return item.Identifier != dependency.Identifier
	})
	cfg.Write()
}

// check wether the config file has a certain dependency
func (cfg Config) HasDependency(dependency Dependency) bool {
	return utility.Some(cfg.Dependencies, func(dep Dependency, _ int) bool {
		return dep.Identifier == dependency.Identifier
	})
}

// check wether the config file has a certain sub-dependency
func (cfg Config) HasSubDependency(dependency Dependency) bool {
	return utility.Some(cfg.SubDependencies, func(dep Dependency, _ int) bool {
		return dep.Identifier == dependency.Identifier
	})
}

func (cfg Config) PrettyPrint() {
	cfgType := "game"
	if cfg.IsPlugin {
		cfgType = "plugin"
	}

	depsToIdentifiers := func(deps []Dependency) []string {
		return utility.Map(deps, func(item Dependency, _ int) string {
			return item.Identifier
		})
	}
	dependencies := utility.SliceToString(depsToIdentifiers(cfg.Dependencies), ", ")
	subDependencies := utility.SliceToString(depsToIdentifiers(cfg.SubDependencies), ", ")

	fmt.Printf("this project is a %s\ndependencies: %v\nsubdependencies: %v\n", cfgType, dependencies, subDependencies)
}

// write the current state of the configuartion to the config file
func (ppm Config) Write() error {
	content, err := json.MarshalIndent(ppm, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(ppm.filePath, content, 0755)
}

// parse the ppm.json file to an object
func ParseConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	content, err := os.ReadFile(file.Name())
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		return Config{}, err
	}
	config.filePath = filePath

	return config, nil
}

// create a ppm.json file
func CreateConfig(path string) (Config, error) {
	configPath := filepath.Join(path, "ppm.json")

	if fileExists, _ := utility.DoesPathExist(configPath); fileExists {
		return Config{}, errors.New("file already exists")
	}

	config := Config{
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
func GetPluginConfig(p paths.Paths, dep Dependency) (Config, error) {
	config, err := ParseConfig(path.Join(p.Addons, dep.Identifier, "ppm.json"))

	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetPathsAndConfig() (paths.Paths, Config, error) {
	pth, err := paths.CreatePathsFromCwd()
	if err != nil {
		return paths.Paths{}, Config{}, err
	}

	config, err := ParseConfig(pth.ConfigFile)
	if err != nil {
		return paths.Paths{}, Config{}, errors.New("could not find ppm.json file - try to run: ppm init")
	}

	return pth, config, nil
}
