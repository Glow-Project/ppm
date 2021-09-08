package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Clone a certain github repository into a certain folder
func Clone(path string, repoName string) error {
	var repository string
	if strings.HasPrefix(repoName, "http") {
		repository = repoName
	} else {
		repository = fmt.Sprintf("https://github.com/Glow-Project/%s", repoName)
	}

	return runGitCommand(exec.Command("git", "clone", repository), path)
}

// Update a certain repository
func Update(path string) error {
	return runGitCommand(exec.Command("git", "pull") ,path)
}

func runGitCommand(command *exec.Cmd, path string) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(path)
	if err != nil {
		return err
	}

	if err := command.Run(); err != nil {
		return err
	}

	err = os.Chdir(currentPath)
	if err != nil {
		return err
	}

	return nil
}