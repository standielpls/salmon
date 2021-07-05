package form

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type FormClient struct {
	Client  http.Client
	BaseUrl string
}

func (h *FormClient) Submit(data map[string]string) (bool, error) {
	urlData := url.Values{}
	for k, v := range data {
		urlData.Set("entry."+k, v)
	}

	req, err := http.NewRequest("POST", h.BaseUrl, strings.NewReader(urlData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := h.Client.Do(req)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Unable submit form: %s", err.Error()))
	}

	if res.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		return false, errors.New(buf.String())
	}

	return true, nil
}
