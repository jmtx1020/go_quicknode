package destinations

import (
	"time"

	"github.com/jmtx1020/go_quicknode/client/client.go"
)

type Destination struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	To          string    `json:"to"`
	WebhookType string    `json:"webhook_type"`
	Service     string    `json:"service"`
	Token       string    `json:"token"`
	PayloadType int       `json:"payload_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DestinationsAPI struct {
	APIWrapper *client.APIWrapper
}

func NewDestinationsAPI(apiToken, baseURL string) *DestinationsAPI {
	apiWrapper := client.NewAPIClient()
}
