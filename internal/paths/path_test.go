package paths

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var separator string

func init() {
	if runtime.GOOS == "windows" {
		separator = "\\"
	} else {
		separator = "/"
	}
}

func cln(s string) string {
	return strings.ReplaceAll(s, "/", separator)
}

func TestPathsCreation(t *testing.T) {
	p1 := CreatePaths("/test")
	require.Equal(t, "/test", p1.Root)
	require.Equal(t, cln("/test/addons"), p1.Addons)
	require.Equal(t, cln("/test/ppm.json"), p1.ConfigFile)

	p2 := CreatePaths("/test/addons/x")
	require.Equal(t, "/test/addons/x", p2.Root)
	require.Equal(t, cln("/test/addons"), p2.Addons)
	require.Equal(t, cln("/test/addons/x/ppm.json"), p2.ConfigFile)

	p3, err := CreatePathsFromCwd()
	require.Nil(t, err)
	require.NotNil(t, p3.Addons)
	require.NotNil(t, p3.ConfigFile)
	require.NotNil(t, p3.Root)
}
