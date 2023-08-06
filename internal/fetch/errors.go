package fetch

import "fmt"

type InvalidVersionError struct {
	Version string
}

func (e *InvalidVersionError) Error() string {
	return fmt.Sprintf("the requested version \"%s\" does not exist.", e.Version)
}

type CloneError struct {
	GitError error
}

func (e *CloneError) Error() string {
	return fmt.Sprintf("something went wrong while cloning: \"%s\"", e.GitError.Error())
}
