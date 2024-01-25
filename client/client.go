package client

import (
	"net/http"
)

func NewAPIClient(apiToken string) *http.Client {
	client := &http.Client{}

	client.Transport = &AuthTransport{
		Token: apiToken,
		Next:  client.Transport,
	}

	return client
}

// AuthTransport is a custom transport to add authentication headers
type AuthTransport struct {
	Token string
	Next  http.RoundTripper
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-api-key", t.Token)
	return t.Next.RoundTrip(req)
}
