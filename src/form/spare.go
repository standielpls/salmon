package form

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type SpareForm struct {
	Date time.Time
	Body string
	Name string
}

func (sf *SpareForm) SubmitForm() (bool, error) {
	client := FormClient{
		Client:  *http.DefaultClient,
		BaseUrl: "https://docs.google.com/forms/u/0/d/e/1FAIpQLSf51R_x-8GB22KHRyG1-azHoJaXu16CUjOnt2pHvrvMeCHyYw/formResponse",
	}

	m := map[string]string{
		SpareFieldMapping[SpareName]:        sf.Name,
		SpareFieldMapping[SpareDay]:         fmt.Sprintf("%d", sf.Date.Day()),
		SpareFieldMapping[SpareMonth]:       fmt.Sprintf("%d", sf.Date.Month()),
		SpareFieldMapping[SpareYear]:        fmt.Sprintf("%d", sf.Date.Year()),
		SpareFieldMapping[SpareDescription]: sf.Body,
		SpareFieldMapping[SpareHours]:       "40",
	}

	fmt.Println(m)
	successful, err := client.Submit(m)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Unable to submit SpareForm: %s", err.Error()))
	}
	return successful, nil
}
