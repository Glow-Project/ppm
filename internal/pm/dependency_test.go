package pm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDepndencyCreation(t *testing.T) {
	d1 := DependencyFromString("dude/project1")
	require.Equal(t, "project1", d1.Identifier)
	require.Equal(t, GithubAsset, d1.Type)
	require.Equal(t, "https://github.com/dude/project1", d1.Url)

	d2 := DependencyFromString("project2")
	require.Equal(t, "project2", d2.Identifier)
	require.Equal(t, GDAsset, d2.Type)
	require.Equal(t, "https://godotengine.org/asset-library/api/asset?filter=project2", d2.Url)

	d3 := DependencyFromString("https://github.com/other/project3")
	require.Equal(t, "project3", d3.Identifier)
	require.Equal(t, GithubAsset, d3.Type)
	require.Equal(t, "https://github.com/other/project3", d3.Url)
}
