package utility

import (
	"fmt"
	"strings"
)

// asset-type enum
const (
	GithubAsset = "GITHUB_ASSET"
	GDAsset     = "GODOT_ASSET"
)

type Dependency struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
	Type       string `json:"type"`
}

// create a new dependency struct from a string
func DependencyFromString(str string) Dependency {
	// the (ideally) unique identifier of the dependency
	var identifier string
	// the url used for accessing the dependency
	var url string
	// the type of the dependency, GithubAsset | GDAsset
	var t string

	if isGithubUrl := IsGithubRepoUrl(str); isGithubUrl || IsUserAndRepo(str) {
		if isGithubUrl {
			url = str
		} else {
			url = fmt.Sprintf("https://github.com/%s", str)
		}
		tmp := strings.Split(str, "/")
		identifier = tmp[len(tmp)-1]
		t = GithubAsset
	} else {
		identifier = strings.Replace(str, " ", "-", -1)
		url = fmt.Sprintf("https://godotengine.org/asset-library/api/asset?filter=%s", str)
		t = GDAsset
	}

	return Dependency{
		Identifier: identifier,
		Url:        url,
		Type:       t,
	}
}
