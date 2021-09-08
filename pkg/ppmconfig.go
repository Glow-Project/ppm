package pkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Representing a ppm.json configuration file
type PpmConfig struct {
	Plugin       bool     `json:"plugin"`
	Dependencies []string `json:"dependencies"`
	filePath string
}

// Add an item safely to the Dependencies property
func (ppm *PpmConfig) AddDependency(dependency string) {
	ppm.Dependencies = append(ppm.Dependencies, dependency)
	ppm.write()
}

// Remove an item safely from the Dependencies property by its name
func (ppm *PpmConfig) RemoveDependency(dependency string) {
	index := IndexOf(dependency, ppm.Dependencies)
	ppm.Dependencies = append(ppm.Dependencies[:index], ppm.Dependencies[index+1:]...)
	ppm.write()
}

// Check wether the config file has a certain dependency
func (ppm PpmConfig) HasDependency(dependency string) bool {
	return StringSliceContains(dependency, ppm.Dependencies)
}

// Write the current state of the configuartion to the config file
func (ppm PpmConfig) write() error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	content, err := json.MarshalIndent(ppm, "", " ")
	if err != nil {
		return err
	}

	os.Chdir(filepath.Dir(ppm.filePath))
	err = ioutil.WriteFile("ppm.json", content, 0755)
	if err != nil {
		os.Chdir(currentPath)
		return err
	}
	
	os.Chdir(currentPath)

	return nil
}

// Parse the ppm.json file to an object
func ParsePpmConfig(filePath string) (PpmConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return PpmConfig{}, err
	}
	defer file.Close()

	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return PpmConfig{}, err
	}

	config := PpmConfig {}
	json.Unmarshal([]byte(content), &config)
	config.filePath = filePath

	return config, nil
}

// Create a new ppm.json file
func CreateNewPpmConfig(path string) error {
	configPath := filepath.Join(path, "ppm.json")

	fileExists, _ := DoesPathExist(configPath)
	
	if fileExists {
		return errors.New("file already exists")
	}
	
	var plugin bool

	if strings.HasSuffix(filepath.Dir(path), "addons") {
		plugin = true
	} else {
		plugin = false
	}

	config := PpmConfig{
		Plugin: plugin,
		Dependencies: []string{},
	}

	content, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		
		return err
	}
	
	file, err := os.Create(filepath.Join(path, "ppm.json"))
	if err != nil {
		return err
	}

	
	_, err = file.Write(content)
	
	return err
}