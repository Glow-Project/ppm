package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// a struct used for simple requests
type Requester struct {
	BaseUrl string
}

func (r Requester) Get(url string) (map[string]any, error) {
	rawResp, err := http.Get(r.BaseUrl + url)
	if err != nil {
		return map[string]any{}, err
	}

	var result map[string]any
	err = json.NewDecoder(rawResp.Body).Decode(&result)

	return result, err
}

func (r Requester) Download(url string, w io.Writer) error {
	response, err := http.Get(r.BaseUrl + url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"request to %s failed with status code %d",
			url,
			response.StatusCode,
		)
	}

	_, err = io.Copy(w, response.Body)
	return err
}
