package account

import (
	"os"
	"testing"

	"github.com/jmtx1020/go_quicknode/client"
)

func TestGetAccountUsage(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/account")

	accountAPI := &AccountAPI{API: apiWrapper}

	err := accountAPI.GetAccountUsage("2024-01-17T22%3A18%3A08.000Z", "2024-02-17T22%3A18%3A08.000Z", "day")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDailyAccountUsage(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/account/daily-usage")

	accountAPI := &AccountAPI{API: apiWrapper}

	err := accountAPI.DailyAccountUsage()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
