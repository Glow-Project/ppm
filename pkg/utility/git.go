package utility

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
)

// Clone a certain github repository into a certain folder
func Clone(path string, repoName string, version string) error {
	// The clonable name of the repository
	var repository string
	// The name of the repository
	var name string

	if IsUrl(repoName) {
		repository = repoName
		tmp := strings.Split(repository, "/")
		name = tmp[len(tmp)-1]
	} else if IsUserAndRepo(repoName) {
		repository = fmt.Sprintf("https://github.com/%s", repoName)
		tmp := strings.Split(repoName, "/")
		name = tmp[1]
	} else {
		// TODO: add godot asset library support
		// name = repoName
		return errors.New("godot asset library not supported yet")
	}

	_, err := git.PlainClone(filepath.Join(path, name), false, &git.CloneOptions{
		URL: repository,
	})

	return err
}

// Update a certain repository
func Update(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&git.PullOptions{})
	if err != nil {
		return err
	}

	return nil
}
