package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// representing a ppm.json configuration file
type PpmConfig struct {
	IsPlugin        bool     `json:"plugin"`
	Dependencies    []string `json:"dependencies"`
	SubDependencies []string `json:"sub-dependencies"`
	filePath        string
}

// add an item safely to the Dependencies property
func (ppm *PpmConfig) AddDependency(dependency string) {
	ppm.Dependencies = append(ppm.Dependencies, dependency)
	ppm.Write()
}

// add an item safely to the sub-dependencies property
func (ppm *PpmConfig) AddSubDependency(dependency string) {
	ppm.SubDependencies = append(ppm.SubDependencies, dependency)
	ppm.Write()
}

// remove ALL (sub)dependencies
func (ppm *PpmConfig) RemoveAllDependencies() {
	ppm.Dependencies = []string{}
	ppm.SubDependencies = []string{}
	ppm.Write()
}

// remove an item safely from the Dependencies property by its name
func (ppm *PpmConfig) RemoveSubDependency(dependency string) {
	index := IndexOf(dependency, ppm.SubDependencies)
	ppm.SubDependencies = append(ppm.SubDependencies[:index], ppm.SubDependencies[index+1:]...)
	ppm.Write()
}

// remove an item safely from the sub-dependencies property by its name
func (ppm *PpmConfig) RemoveDependency(dependency string) {
	if len(ppm.Dependencies) == 1 {
		ppm.Dependencies = []string{}
	} else {
		index := IndexOf(dependency, ppm.Dependencies)
		ppm.Dependencies = append(ppm.Dependencies[:index], ppm.Dependencies[index+1:]...)
	}
	ppm.Write()
}

// check wether the config file has a certain dependency
func (ppm PpmConfig) HasDependency(dependency string) bool {
	return SliceContains(dependency, ppm.Dependencies)
}

// check wether the config file has a certain sub-dependency
func (ppm PpmConfig) HasSubDependency(dependency string) bool {
	return SliceContains(dependency, ppm.SubDependencies)
}

func (ppm PpmConfig) PrettyPrint() {
	ppmType := "game"
	if ppm.IsPlugin {
		ppmType = "plugin"
	}

	dependencies := SliceToString(ppm.Dependencies, ", ")
	subDependencies := SliceToString(ppm.SubDependencies, ", ")

	fmt.Printf("this project is a %s\ndependencies: %v\nsubdependencies: %v", ppmType, dependencies, subDependencies)
}

// write the current state of the configuartion to the config file
func (ppm PpmConfig) Write() error {
	content, err := json.MarshalIndent(ppm, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ppm.filePath, content, 0755)
	if err != nil {
		return err
	}

	return nil
}

// parse the ppm.json file to an object
func ParsePpmConfig(filePath string) (PpmConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return PpmConfig{}, err
	}
	defer file.Close()

	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
		return PpmConfig{}, err
	}

	config := PpmConfig{}
	json.Unmarshal([]byte(content), &config)
	config.filePath = filePath

	return config, nil
}

// create a new ppm.json file
func CreateNewPpmConfig(path string) error {
	configPath := filepath.Join(path, "ppm.json")

	fileExists, _ := DoesPathExist(configPath)

	if fileExists {
		return errors.New("file already exists")
	}

	var isPlugin bool

	if strings.HasSuffix(filepath.Dir(path), "addons") {
		isPlugin = true
	} else {
		isPlugin = false
	}

	config := PpmConfig{
		IsPlugin:        isPlugin,
		Dependencies:    []string{},
		SubDependencies: []string{},
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

func GetPluginConfig(dirPath string, dependency string) (PpmConfig, error) {
	tmp := strings.Split(dependency, "/")
	var dependencyName string

	if tmp[len(tmp)-1] != "/" {
		dependencyName = tmp[len(tmp)-1]
	} else {
		dependencyName = tmp[len(tmp)-2]
	}

	config, err := ParsePpmConfig(path.Join(dirPath, dependencyName, "ppm.json"))

	if err != nil {
		return PpmConfig{}, err
	}

	return config, nil
}
