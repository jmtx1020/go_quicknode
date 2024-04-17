package account

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jmtx1020/go_quicknode/client"
)

type AccountUsage struct {
	Interval     string    `json:"interval"`
	StorageBytes string    `json:"storageBytes"`
	Gateways     string    `json:"gateways"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}

type AccountAPI struct {
	API *client.APIWrapper
}

func NewAccountAPI(apiToken, baseURL string) *AccountAPI {
	apiWrapper := client.NewAPIWrapper(apiToken, baseURL)
	return &AccountAPI{API: apiWrapper}
}

func (a *AccountAPI) GetAccountUsage(startDate, endDate, interval string) error {
	a.API.SetBaseURL("https://api.quicknode.com/ipfs/rest/v1/account")
	endpoint := fmt.Sprintf("%s/usage?startDate=%s&endDate=%s&interval=%s", a.API.BaseURL, startDate, endDate, interval)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := a.API.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error: %s", body)
	}

	return nil
}

func (a *AccountAPI) DailyAccountUsage() error {
	a.API.SetBaseURL("https://api.quicknode.com/ipfs/rest/v1/account")
	endpoint := fmt.Sprintf("%s/daily-usage", a.API.BaseURL)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := a.API.Client.Do(req)
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
