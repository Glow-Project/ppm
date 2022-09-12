package utility

import (
	"fmt"
	"strings"
)

const (
	GithubAsset = "GITHUB_ASSET"
	GDAsset     = "GODOT_ASSET"
)

type Dependency struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
	Type       string `json:"type"`
}

func DependencyFromString(str string) *Dependency {
	var identifier string
	var url string
	var t string

	if isGithubUrl := IsGithubRepoUrl(str); isGithubUrl || IsUserAndRepo(str) {
		if isGithubUrl {
			url = str
		} else {
			url = fmt.Sprintf("https://github.com/%s", str)
		}
		tmp := strings.Split(str, "/")
		identifier = fmt.Sprintf("%s/%s", tmp[len(tmp)-2], tmp[len(tmp)-1])
		t = GithubAsset
	} else {
		identifier = strings.Replace(str, " ", "-", -1)
		url = fmt.Sprintf("https://godotengine.org/asset-library/api/asset?filter=%s", str)
		t = GithubAsset
	}

	return &Dependency{
		Identifier: identifier,
		Url:        url,
		Type:       t,
	}
}
