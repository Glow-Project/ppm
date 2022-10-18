package fetch

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Glow-Project/ppm/pkg/utility"
	"github.com/go-git/go-git/v5"
)

// install a dependency `dep` into its directory inside the `addons` directory
func InstallDependency(dep *utility.Dependency, paths *utility.Paths) error {
	var err error
	if dep.Type == utility.GithubAsset {
		err = installGithubRepo(dep, paths)
	} else {
		err = installGodotAsset(dep, paths)
	}

	return err
}

// install a plugin from github
func installGithubRepo(dep *utility.Dependency, paths *utility.Paths) error {
	fullPath := path.Join(paths.Addons, dep.Identifier)
	_, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL: dep.Url,
	})
	if err != nil {
		return err
	}

	return nil
}

// install a plugin from the godot asset store
func installGodotAsset(dep *utility.Dependency, paths *utility.Paths) error {
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
	id := data["result"].([]interface{})[0].(map[string]interface{})["asset_id"]

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
	err = unzip(f.Name(), paths.Addons)
	if err != nil {
		return err
	}

	return nil
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

		path := filepath.Join(dest, f.Name)

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
