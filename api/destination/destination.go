package destinations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jmtx1020/go_quicknode/client"
)

type DestinationPayload struct {
	Name        string `json:"name"`
	ToURL       string `json:"to_url"`
	WebhookType string `json:"webhook_type"`
	Service     string `json:"service"`
	PayloadType int    `json:"payload_type"`
}

type Destination struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	To          string    `json:"to"`
	WebhookType string    `json:"webhook_type"`
	Service     string    `json:"service"`
	Token       string    `json:"token"`
	PayloadType int       `json:"payload_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DestinationAPI struct {
	API *client.APIWrapper
}

func NewDestinationsAPI(apiToken, baseURL string) *DestinationAPI {
	apiWrapper := client.NewAPIWrapper(apiToken, baseURL)
	return &DestinationAPI{API: apiWrapper}
}

func (d *DestinationAPI) GetAllDestinations() ([]Destination, error) {
	d.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/destinations")
	endpoint := fmt.Sprintf("%s", d.API.BaseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.API.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var destinations []Destination
	err = json.Unmarshal(body, &destinations)
	if err != nil {
		return nil, err
	}
	return destinations, nil
}

func (d *DestinationAPI) GetDestinationByID(destinationID string) (*Destination, error) {
	d.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/destinations")
	endpoint := fmt.Sprintf("%s/%s", d.API.BaseURL, destinationID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.API.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var destination Destination
	err = json.Unmarshal(body, &destination)
	if err != nil {
		return nil, err
	}

	return &destination, nil
}

func (d *DestinationAPI) CreateDestination(name, toURL, webhookType, service string, payloadType int) (*Destination, error) {
	d.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/destinations")

	// create payload
	payload := DestinationPayload{
		Name:        name,
		ToURL:       toURL,
		WebhookType: webhookType,
		Service:     service,
		PayloadType: payloadType,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", d.API.BaseURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.API.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var destination Destination
	err = json.Unmarshal(body, &destination)
	if err != nil {
		return nil, err
	}
	return &destination, nil
}

func (d *DestinationAPI) DeleteDestinationByID(destinationID string) error {
	d.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/destinations")
	endpoint := fmt.Sprintf("%s/%s", d.API.BaseURL, destinationID)

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := d.API.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
