package pinning

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jmtx1020/go_quicknode/client"
)

type PinnedObjectPayload struct {
	CID     string                  `json:"cid"`
	Name    string                  `json:"name"`
	Origins []string                `json:"origins"`
	Meta    PinnedObjectPayloadMeta `json:"meta"`
}

type PinnedObjectPayloadMeta struct {
	Test      string      `json:"test"`
	MoreValue interface{} `json:"morevalue"`
}

type PinnedObjectResponse struct {
	Data       []PinnedObject `json:"data"`
	TotalItems int            `json:"totalItems"`
	TotalPages int            `json:"totalPages"`
	PageNumber int            `json:"pageNumber"`
}

type PinnedObject struct {
	ID        string  `json:"id"`
	RequestID string  `json:"requestId"`
	Status    string  `json:"status"`
	CID       string  `json:"cid"`
	Name      string  `json:"name"`
	Origins   Origins `json:"origins"`
	Meta      struct {
		Test      string      `json:"test"`
		MoreValue interface{} `json:"morevalue"`
	}
	UUID        string `json:"uuid"`
	CreatedAT   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	ContentType string `json:"contentType"`
	Size        string `json:"size"`
	Source      string `json:"source"`
	ParentID    string `json:"parentId"`
	Type        string `json:"type"`
	Path        string `json:"path"`
}

type Origins []string

func (o *Origins) UnmarshalJSON(data []byte) error {
	var rawOrigins string

	if err := json.Unmarshal(data, &rawOrigins); err != nil {
		return err
	}
	*o = strings.Split(rawOrigins, ",")
	return nil
}

type PinningAPI struct {
	API *client.APIWrapper
}

func NewPinningAPI(apiToken, baseURL string) *PinningAPI {
	apiWrapper := client.NewAPIWrapper(apiToken, baseURL)
	return &PinningAPI{API: apiWrapper}
}

func (p *PinningAPI) CreatePinnedObject(cid, name string, origins []string, meta PinnedObjectPayloadMeta) (*PinnedObject, error) {
	p.API.SetBaseURL("https://api.quicknode.com/ipfs/rest/v1/pinning")
	endpoint := p.API.BaseURL

	payload := PinnedObjectPayload{
		CID:     cid,
		Name:    name,
		Origins: origins,
		Meta:    meta,
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

	resp, err := p.API.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to create pinned object: %s", body)
	}

	var pinnedObject PinnedObject
	err = json.Unmarshal(body, &pinnedObject)
	if err != nil {
		return nil, err
	}
	return &pinnedObject, nil
}

func (p *PinningAPI) GetAllPinnedObjects(pageNumber, resultsPerPage int) (PinnedObjectResponse, error) {
	p.API.SetBaseURL("https://api.quicknode.com/ipfs/rest/v1/pinning")
	endpoint := fmt.Sprintf("%s?pageNumber=%d&perPage=%d", p.API.BaseURL, pageNumber, resultsPerPage)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return PinnedObjectResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.API.Client.Do(req)
	if err != nil {
		return PinnedObjectResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PinnedObjectResponse{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return PinnedObjectResponse{}, fmt.Errorf("failed to get all pinned objects: %s", body)
	}

	var response PinnedObjectResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return PinnedObjectResponse{}, err
	}
	return response, nil
}
