package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jmtx1020/go_quicknode/api/destinations"
	"github.com/jmtx1020/go_quicknode/client"
)

type NotificationData struct {
	Data     []NotificationEvent   `json:"data"`
	PageInfo NotificationEventPage `json:"pageInfo"`
}

type NotificationEventPage struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type NotificationEvent struct {
	ID               string    `json:"id"`
	Destination_ID   string    `json:"destination_id"`
	NotificationID   string    `json:"notification_id"`
	BlockNumber      string    `json:"block_number"`
	TxMatchedIndexes string    `json:"tx_matched_indexes"`
	Delivered        bool      `json:"delivered"`
	DeliveredError   string    `json:"delivered_error"`
	DeliveredAt      time.Time `json:"delivered_at"`
	CreatedAt        time.Time `json:"created_at"`
	To               string    `json:"to"`
}

type Notification struct {
	ID           string                     `json:"id"`
	Name         string                     `json:"name"`
	Expression   string                     `json:"expression"`
	Network      string                     `json:"network"`
	Enabled      bool                       `json:"enabled"`
	Destinations []destinations.Destination `json:"destinations"`
	CreatedAt    time.Time                  `json:"created_at"`
	UpdatedAt    time.Time                  `json:"updated_at"`
}

type NotificationPayload struct {
	Name         string   `json:"name"`
	Expression   string   `json:"expression"`
	Network      string   `json:"network"`
	Destinations []string `json:"destinationIds"`
}

type NotificationValidationPayload struct {
	Expression string `json:"expression"`
	Network    string `json:"network"`
}

type NotificationAPI struct {
	API *client.APIWrapper
}

func NewNotificationAPI(apiToken, baseURL string) *NotificationAPI {
	apiWrapper := client.NewAPIWrapper(apiToken, baseURL)
	return &NotificationAPI{API: apiWrapper}
}

func (n *NotificationAPI) CreateNotification(name, expression, network string, destinatons []string) (*Notification, error) {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf(n.API.BaseURL)

	payload := NotificationPayload{
		Name:         name,
		Expression:   expression,
		Network:      network,
		Destinations: destinatons,
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

	resp, err := n.API.Client.Do(req)
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

	var notification Notification
	err = json.Unmarshal(body, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (n *NotificationAPI) GetAllNotifications() ([]Notification, error) {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf("%s", n.API.BaseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := n.API.Client.Do(req)
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

	var notifications []Notification
	err = json.Unmarshal(body, &notifications)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (n *NotificationAPI) GetNotificationByID(id string) (*Notification, error) {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf("%s/%s", n.API.BaseURL, id)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := n.API.Client.Do(req)
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

	var notification Notification
	err = json.Unmarshal(body, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (n *NotificationAPI) DeleteNotificationByID(id string) error {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf("%s/%s", n.API.BaseURL, id)

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := n.API.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error: %s", body)
	}

	return nil
}

func (n *NotificationAPI) UpdateNotificationByID(id, name, expression string, destinations []string) (*Notification, error) {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf("%s/%s", n.API.BaseURL, id)

	payload := NotificationPayload{
		Name:         name,
		Expression:   expression,
		Destinations: destinations,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.API.Client.Do(req)
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

	var notification Notification
	err = json.Unmarshal(body, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (n *NotificationAPI) ToggleNotificationByID(id string, enable bool) error {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")

	var toggle string
	if enable {
		toggle = "enable"
	} else {
		toggle = "disable"
	}

	endpoint := fmt.Sprintf("%s/%s/%s", n.API.BaseURL, id, toggle)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.API.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error: %s", body)
	}
	return nil
}

func (n *NotificationAPI) GetNotificationEventsByID(id string) (*NotificationData, error) {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf("%s/%s/events", n.API.BaseURL, id)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := n.API.Client.Do(req)
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

	var notificationData NotificationData
	err = json.Unmarshal(body, &notificationData)
	if err != nil {
		return nil, err
	}
	return &notificationData, nil
}

func (n *NotificationAPI) ValidateExpression(expression, network string) error {
	n.API.SetBaseURL("https://api.quicknode.com/quickalerts/rest/v1/notifications")
	endpoint := fmt.Sprintf("%s/validate", n.API.BaseURL)

	payload := NotificationValidationPayload{
		Expression: expression,
		Network:    network,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.API.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error: %s", body)
	}

	return nil
}
