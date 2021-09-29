package form

import (
	"errors"
	"fmt"
	"net/http"
)

type StanForm struct{}

func (sf StanForm) SubmitForm() (bool, error) {
	client := FormClient{
		Client: *http.DefaultClient,
		// TODO: Also belongs in a config
		BaseUrl: "https://docs.google.com/forms/u/0/d/e/1FAIpQLSf4hSYBd3o3MCEIA-jq1jl4RWqMULQQbwvVLqXR3wYWmHimow/formResponse",
	}

	m := map[string]string{
		StanFieldMapping[StanName]:    "Stan Test",
		StanFieldMapping[StanEmail]:   "stan@stanmail.com",
		StanFieldMapping[StanSubject]: "Flying animals",
		StanFieldMapping[StanBody]:    "Flying animal friend",
	}
	successful, err := client.Submit(m)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Unable to submit StanForm: %s", err.Error()))
	}
	return successful, nil
}
