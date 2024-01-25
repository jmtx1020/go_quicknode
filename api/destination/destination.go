package destinations

import (
	"net/http"

	"github.com/jmtx1020/go_quicknode/client"
)

type Item struct {
	// Item fields
}

type APIWrapper struct {
	Client *http.Client
}

func NewAPIWrapper(apiToken string) *APIWrapper {
	return &APIWrapper{
		Client: client.NewAPIClient(apiToken),
	}
}
