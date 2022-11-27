package utility

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// check wether an absolute path exists
func DoesPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// create a certain directory if it doesn't exist
func CheckOrCreateDir(path string) error {
	pathExisting, _ := DoesPathExist(path)
	if pathExisting {
		return nil
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func GetVersionOrNot(dependency string) (string, string) {
	version := ""

	splitDependency := strings.Split(dependency, "@")

	if len(splitDependency) > 1 {
		version = splitDependency[1]
	}
	dependencyName := splitDependency[0]

	return dependencyName, version
}

func PrintDone() {
	fmt.Print(color.GreenString("\rdone\n"))
}

func GetPluginIdentifier(name string) string {
	if strings.HasPrefix(name, "https://") {
		urlParts := strings.Split(name, "/")
		return urlParts[len(urlParts)-1]
	} else {
		return name
	}
}

func GetPluginName(name string) string {
	s := GetPluginIdentifier(name)
	if strings.Contains(s, "/") {
		return strings.Split(s, "/")[1]
	} else {
		return s
	}
}

func GetPathsAndConfig() (Paths, PpmConfig, error) {
	paths, err := CreatePathsFromCwd()
	if err != nil {
		return Paths{}, PpmConfig{}, err
	}

	config, err := ParsePpmConfig(paths.ConfigFile)
	if err != nil {
		return Paths{}, PpmConfig{}, errors.New("could not find ppm.json file - try to run: ppm init")
	}

	return paths, config, nil
}

func SliceToString(slice []string, seperator string) string {
	if len(slice) == 0 {
		return "none"
	}

	str := ""
	for _, item := range slice {
		if len(str) != 0 {
			str += seperator
		}

		str += item
	}

	return str
}
