package test

import "net/http"

type transport struct {
	f func(*http.Request) (*http.Response, error)
}

// RoundTrip calls the transport's RoundTrip function.
func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	return t.f(r)
}

func NewHttpClient(fn func(req *http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: &transport{fn},
	}
}
