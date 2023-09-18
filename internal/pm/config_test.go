package pm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ppmConfigStr = `
{
	"plugin": false,
	"dependencies": [
	 {
	  "identifier": "additional-audio-players",
	  "version": null,
	  "url": "https://github.com/Glow-Project/additional-audio-players",
	  "type": "GITHUB_ASSET"
	 }
	],
	"sub-dependencies": []
   }
`

func withTempFile(content []byte, fn func(path *os.File)) {
	f, err := os.CreateTemp("", "tmpfile-")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	defer os.Remove(f.Name())

	if _, err := f.Write(content); err != nil {
		panic(err)
	}

	fn(f)
}

func withTempDir(fn func(path string)) {
	tmpDir := os.TempDir()

	dir, err := os.MkdirTemp(tmpDir, "*")
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir)

	fn(dir)
}

func TestConfigCreation(t *testing.T) {
	withTempDir((func(path string) {
		ppm, err := CreateConfig(path)
		assert.Nil(t, err)
		assert.False(t, ppm.IsPlugin)

		assert.Equal(t, len(ppm.Dependencies), 0)
		assert.Equal(t, len(ppm.SubDependencies), 0)
	}))

	withTempFile([]byte(ppmConfigStr), func(v *os.File) {
		ppm, err := ParseConfig(v.Name())
		assert.Nil(t, err)
		assert.False(t, ppm.IsPlugin)

		assert.Equal(t, len(ppm.Dependencies), 1)
		assert.Equal(t, len(ppm.SubDependencies), 0)

		dep := ppm.Dependencies[0]
		assert.Equal(t, dep.Identifier, "additional-audio-players")
		assert.Nil(t, dep.Version)
		assert.Equal(t, dep.Url, "https://github.com/Glow-Project/additional-audio-players")
		assert.Equal(t, dep.Type, "GITHUB_ASSET")
	})
}
