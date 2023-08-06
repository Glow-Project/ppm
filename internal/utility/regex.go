package utility

import "regexp"

var (
	urlRegex             = regexp.MustCompile(`^https?:\/\/[a-zA-Z0-9_\-\.]+\.[a-zA-Z]{1,5}([a-zA-Z0-9_\/\-\=\&\?\:]+)*$`)
	githubUrlRegex       = regexp.MustCompile(`^https?:\/\/github\.com(\/[a-zA-Z0-9_\-\=\&\?\:]+){2}$`)
	userAndRepoRegex     = regexp.MustCompile(`^[a-zA-Z0-9_\-]+\/[a-zA-Z0-9_\-]+$`)
	assetDependencyRegex = regexp.MustCompile(`^[a-zA-Z0-9\-\_]+$`)
)

func IsUrl(str string) bool {
	return urlRegex.MatchString(str)
}

func IsGithubRepoUrl(str string) bool {
	return githubUrlRegex.MatchString(str)
}

func IsUserAndRepo(str string) bool {
	return userAndRepoRegex.MatchString(str)
}

func IsAssetDependency(str string) bool {
	return assetDependencyRegex.MatchString(str)
}
