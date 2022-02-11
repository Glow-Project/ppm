package utility

import (
	"os"
	"path/filepath"
	"strings"
)

type Paths struct {
	Root   string
	Addons string
}

func CreatePaths(rootPath string) Paths {
	var addons string
	if strings.HasSuffix(filepath.Dir(rootPath), "addons") {
		addons = filepath.Dir(rootPath)
	} else {
		addons = filepath.Join(rootPath, "addons")
	}

	return Paths{
		Root:   rootPath,
		Addons: addons,
	}
}

func CreatePathsFromCwd() (Paths, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return Paths{}, err
	}

	return CreatePaths(rootPath), nil
}
