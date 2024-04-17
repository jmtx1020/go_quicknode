package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jmtx1020/go_quicknode/client"
)

type Gateway struct {
	ID        string    `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"`
	IsPrivate bool      `json:"isPrivate"`
	IsEnabled bool      `json:"isEnabled"`
	CreatedAT time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GatewayPayload struct {
	Name      string `json:"name"`
	IsPrivate bool   `json:"isPrivate"`
	IsEnabled bool   `json:"isEnabled"`
}

type GatewayAPI struct {
	API *client.APIWrapper
}

func NewGatewayAPI(apiToken, baseURL string) *GatewayAPI {
	apiWrapper := client.NewAPIWrapper(apiToken, baseURL)
	return &GatewayAPI{API: apiWrapper}
}

func (g *GatewayAPI) CreateGateway(name string, isPrivate, isEnabled bool) (*Gateway, error) {
	g.API.SetBaseURL("https://api.quicknode.com/ipfs/rest/v1/gateway")
	endpoint := g.API.BaseURL

	payload := GatewayPayload{
		Name:      name,
		IsPrivate: isPrivate,
		IsEnabled: isEnabled,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.API.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("error: %s", body)
	}

	var gateway Gateway
	err = json.Unmarshal(body, &gateway)
	if err != nil {
		return nil, err
	}
	return &gateway, nil
}
