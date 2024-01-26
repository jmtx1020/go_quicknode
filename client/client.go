package client

import (
	"net/http"
)

type AuthTransport struct {
	Token string
	Next  http.RoundTripper
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-api-key", t.Token)
	return t.Next.RoundTrip(req)
}

type APIWrapper struct {
	Client  *http.Client
	BaseURL string
}

func NewAPIClient(apiToken string) *http.Client {
	return &http.Client{
		Transport: &AuthTransport{
			Token: apiToken,
			Next:  http.DefaultTransport,
		},
	}
}

func NewAPIWrapper(apiToken string, baseURL string) *APIWrapper {
	return &APIWrapper{
		Client:  NewAPIClient(apiToken),
		BaseURL: baseURL,
	}
}

func (a *APIWrapper) SetBaseURL(baseURL string) {
	a.BaseURL = baseURL
}
