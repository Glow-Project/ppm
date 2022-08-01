package tests

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/Glow-Project/ppm/pkg/utility"
)

const (
	pathNotEqualMessage string = "Path \"%s\" is not equal to \"%s\""
)

func TestPaths(t *testing.T) {
	t.Run("Create Paths instance as Game", func(t *testing.T) {
		path := "/this/is/a/test"
		addonsPath := filepath.Join(path + "/addons")
		paths := utility.CreatePaths(path)
		fmt.Println(path)
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
		paths := utility.CreatePaths(path)

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

		if utility.IndexOf("Test 1", arr) != 0 ||
			utility.IndexOf("Test 2", arr) != 1 ||
			utility.IndexOf("Test 3", arr) != 2 {
			t.Fail()
		}
	})

	t.Run("Test StringSliceContains function", func(t *testing.T) {
		if !utility.SliceContains("Test", []string{"tEst", "setT", "Test"}) || utility.SliceContains("Test", []string{"idka", "rhiugh", "zuroghz"}) {
			t.Fail()
		}
	})

	t.Run("Test URL recognition functions", func(t *testing.T) {
		if !utility.IsUrl("https://something.my-url.com") ||
			utility.IsUrl("IAmNotAnURL") {
			t.Error("URL recognition broken")
		}
		if !utility.IsGithubRepoUrl("https://github.com/User/Repository") ||
			utility.IsGithubRepoUrl("https://github.com/User") {
			t.Error("Github Repository URL recognition broken")
		}
		if !utility.IsUserAndRepo("User/Repository") ||
			utility.IsUserAndRepo("UserRepository") {
			t.Error("User and Repository recognition broken")
		}
	})
}
