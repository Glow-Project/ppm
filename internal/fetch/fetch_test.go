package fetch

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBadUrl(t *testing.T) {
	req := Requester{
		BaseUrl: "<not exist>",
	}

	_, err := req.Get("")
	require.NotNil(t, err)

	err = req.Download("", io.Discard)
	require.NotNil(t, err)
}
