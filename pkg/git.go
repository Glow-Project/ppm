package pkg

import (
	"fmt"
	"net/url"
	otherPath "path"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
)

// Clone a certain github repository into a certain folder
func Clone(path string, repoName string, version string) error {
	var repository string
	if strings.HasPrefix(repoName, "https://") {
		repository = repoName
	} else {
		repository = fmt.Sprintf("https://github.com/Glow-Project/%s", repoName)
	}

	if len(version) != 0 {
		// repository = fmt.Sprintf("%s.git", repository)

		//! Add versions
		//! Clone certain tag
		uri, err := url.Parse(repository)
		if err != nil {
			return nil
		}

		uri.Path = otherPath.Join(uri.Path, "tree", version)
		fmt.Println(uri.String())
		_, err = git.PlainClone(filepath.Join(path, repoName), false, &git.CloneOptions{
			URL: uri.String(),
		})

		return err
	}

	_, err := git.PlainClone(filepath.Join(path, repoName), false, &git.CloneOptions{
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
