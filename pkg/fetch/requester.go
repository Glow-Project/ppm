package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Requester struct {
	BaseUrl string
}

func (r *Requester) Get(url string) (map[string]any, error) {
	rawResp, err := http.Get(r.BaseUrl + url)
	if err != nil {
		return map[string]any{}, err
	}

	var result map[string]any
	json.NewDecoder(rawResp.Body).Decode(&result)

	return result, err
}

func (r *Requester) Download(url string, w io.Writer) error {
	response, err := http.Get(r.BaseUrl + url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("request to %s failed", url)
	}

	_, err = io.Copy(w, response.Body)
	if err != nil {
		return err
	}

	return nil
}
