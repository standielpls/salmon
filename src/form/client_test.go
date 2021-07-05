package form_test

import (
	"net/http"
	"testing"

	"salmon/src/form"
	"salmon/src/test"

	"github.com/google/go-cmp/cmp"
)

func TestSubmit(t *testing.T) {
	table := []struct {
		label     string
		fn        func(*http.Request) (*http.Response, error)
		formData  map[string]string
		expSubmit bool
	}{
		{
			label: "can submit",
			fn: func(req *http.Request) (*http.Response, error) {
				return nil, nil
			},
			formData:  map[string]string{},
			expSubmit: false,
		},
	}

	for i := 0; i < len(table); i++ {
		tc := table[i]
		t.Run(tc.label, func(t *testing.T) {

			mockHttpClient := test.NewHttpClient(tc.fn)
			fc := form.FormClient{Client: *mockHttpClient, BaseUrl: "http.com"}

			submitted, err := fc.Submit(tc.formData)
			if err != nil {
				t.Fatalf("Unable to submit: %s", err.Error())
			}
			diff := cmp.Diff(submitted, tc.expSubmit)
			if diff != "" {
				t.Fatalf("Unexpected result returned: %s", diff)
			}
		})
	}
}
