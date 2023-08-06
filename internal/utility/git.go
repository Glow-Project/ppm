package utility

import (
	git "github.com/go-git/go-git/v5"
)

// update a certain repository
func UpdateGithubRepo(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	return w.Pull(&git.PullOptions{})
}
