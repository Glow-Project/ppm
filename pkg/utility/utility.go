package utility

import (
	"errors"
	"fmt"
	"os"
	"regexp"
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

// get the index of a certain item in a string slice
func IndexOf[t comparable](target t, data []t) int {
	for index, value := range data {
		if target == value {
			return index
		}
	}
	return -1
}

// check wether a certain item exists in a string slice
func SliceContains[t comparable](target t, data []t) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}

	return false
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

func GetPluginName(name string) string {
	if strings.HasPrefix(name, "https://") {
		urlParts := strings.Split(name, "/")
		return urlParts[len(urlParts)-1]
	} else {
		return name
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

// get the first predicate match of the slice
//
// returns nil if none of the items match
func GetFirstMatch[t any](slice []t, predicate func(t, int) bool) *t {
	for i, value := range slice {
		if predicate(value, i) {
			return &value
		}
	}

	return nil
}

// returns a new slice containing all the items on which the predicate returns true
func Filter[t any](slice []t, predicate func(t, int) bool) []t {
	newSlice := []t{}
	for i, item := range slice {
		if predicate(item, i) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}

func IsUrl(str string) bool {
	a, _ := regexp.Match(`https?:\/\/[a-zA-Z0-9_\-\.]+\.[a-zA-Z]{1,5}([a-zA-Z0-9_\/\-\=\&\?\:]+)*`, []byte(str))
	return a
}

func IsGithubRepoUrl(str string) bool {
	a, _ := regexp.Match(`https?:\/\/github\.com(\/[a-zA-Z0-9_\-\=\&\?\:]+){2}`, []byte(str))
	return a
}

func IsUserAndRepo(str string) bool {
	a, _ := regexp.Match(`[a-zA-Z0-9_\-]+\/[a-zA-Z0-9_\-]+`, []byte(str))
	return a
}
