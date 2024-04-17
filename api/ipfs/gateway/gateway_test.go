package gateway

import (
	"os"
	"testing"

	"github.com/jmtx1020/go_quicknode/client"
)

func TestCreateGateway(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/gateway")
	gatewayAPI := &GatewayAPI{API: apiWrapper}

	_, err := gatewayAPI.CreateGateway(
		"testing-api-gateway-4",
		true,
		false,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
