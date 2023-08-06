package fetch

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Glow-Project/ppm/internal/utility"
	"github.com/go-git/go-git/v5"
)

// install a dependency `dep` into its directory inside the `addons` directory
func InstallDependency(dep utility.Dependency, paths utility.Paths) error {
	if dep.Type == utility.GithubAsset {
		return installGithubRepo(dep, paths)
	}

	return installGodotAsset(dep, paths)
}

// install a plugin from github
func installGithubRepo(dep utility.Dependency, paths utility.Paths) error {
	fullPath := path.Join(paths.Addons, dep.Identifier)
	repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL: dep.Url,
	})
	if err != nil {
		return &CloneError{err}
	}

	if dep.Version != nil {
		r, err := repo.Tag(*dep.Version)
		if err == nil {
			wt, _ := repo.Worktree()
			wt.Checkout(&git.CheckoutOptions{Hash: r.Hash()})
		} else {
			return &InvalidVersionError{Version: *dep.Version}
		}
	}

	return nil
}

// install a plugin from the godot asset store
func installGodotAsset(dep utility.Dependency, paths utility.Paths) error {
	r := Requester{}
	data, err := r.Get(dep.Url)
	if err != nil {
		return err
	}

	/* structure of data:
		{
	    	"result": [
	  			{
	        		"asset_id": "<id>"
	    		}
			]
		}
	*/
	results := data["result"].([]interface{})
	if len(results) == 0 {
		return fmt.Errorf("no results for dependency \"%s\"", dep.Identifier)
	}

	id := results[0].(map[string]interface{})["asset_id"].(string)

	data, err = r.Get(fmt.Sprintf("https://godotengine.org/asset-library/api/asset/%s", id))
	if err != nil {
		return err
	}

	dwdUrl := data["download_url"].(string)
	f, err := os.CreateTemp("", "tempfile")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())

	r.Download(dwdUrl, f)
	f.Close()
	return unzip(f.Name(), paths.Addons)
}

// unzip a .zip file from src into dest
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	shaRegex := regexp.MustCompile("-[0-9a-f]{40}")

	writeFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		fileName := string(shaRegex.ReplaceAll([]byte(f.Name), []byte("")))
		path := filepath.Join(dest, fileName)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := writeFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
