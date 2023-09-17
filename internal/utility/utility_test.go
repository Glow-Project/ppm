package utility

import (
	"testing"

	"path/filepath"

	"github.com/Glow-Project/ppm/internal/paths"
)

const (
	pathNotEqualMessage string = "Path \"%s\" is not equal to \"%s\""
)

func TestPaths(t *testing.T) {
	t.Run("Create Paths instance as Game", func(t *testing.T) {
		path := "/this/is/a/test"
		addonsPath := filepath.Join(path + "/addons")
		paths := paths.CreatePaths(path)
		if paths.Root != path {
			t.Errorf(pathNotEqualMessage, paths.Root, path)
		}
		if paths.Addons != addonsPath {
			t.Errorf(pathNotEqualMessage, paths.Addons, addonsPath)
		}
	})

	t.Run("Create Paths instance as Plugin", func(t *testing.T) {
		path := "/this/is/a/test/addons/myplugin"
		addonsPath := filepath.Join(path, "..")
		paths := paths.CreatePaths(path)

		if paths.Root != path {
			t.Errorf(pathNotEqualMessage, paths.Root, path)
		}
		if paths.Addons != addonsPath {
			t.Errorf(pathNotEqualMessage, paths.Addons, addonsPath)
		}
	})
}

func TestUtility(t *testing.T) {
	t.Run("Test IndexOf function", func(t *testing.T) {
		arr := []string{
			"Test 1",
			"Test 2",
			"Test 3",
		}

		if IndexOf("Test 1", arr) != 0 ||
			IndexOf("Test 2", arr) != 1 ||
			IndexOf("Test 3", arr) != 2 {
			t.Fail()
		}
	})

	t.Run("Test StringSliceContains function", func(t *testing.T) {
		if !SliceContains("Test", []string{"tEst", "setT", "Test"}) || SliceContains("Test", []string{"idka", "rhiugh", "zuroghz"}) {
			t.Fail()
		}
	})

	t.Run("Test URL recognition functions", func(t *testing.T) {
		if !IsUrl("https://something.my-url.com") ||
			IsUrl("IAmNotAnURL") {
			t.Error("URL recognition broken")
		}
		if !IsGithubRepoUrl("https://github.com/User/Repository") ||
			IsGithubRepoUrl("https://github.com/User") {
			t.Error("Github Repository URL recognition broken")
		}
		if !IsUserAndRepo("User/Repository") ||
			IsUserAndRepo("UserRepository") {
			t.Error("User and Repository recognition broken")
		}
	})
}
