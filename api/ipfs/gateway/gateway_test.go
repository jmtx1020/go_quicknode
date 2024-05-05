package gateway

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/jmtx1020/go_quicknode/client"
)

const (
	charset = "0123456789"
	length  = 6
)

func TestCreateGateway(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/gateway")
	gatewayAPI := &GatewayAPI{API: apiWrapper}

	randomStr := randomString(length)

	_, err := gatewayAPI.CreateGateway(
		fmt.Sprintf("testing-api-%s", randomStr),
		false,
		false,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetAllGateways(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/gateway")
	gatewayAPI := &GatewayAPI{API: apiWrapper}

	_, err := gatewayAPI.GetAllGateways()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetGatewayByName(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/gateway")
	gatewayAPI := &GatewayAPI{API: apiWrapper}

	gateway1, err := gatewayAPI.GetAllGateways()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	gateway2, err := gatewayAPI.GetGetwayByName(gateway1[0].Name)
	if gateway1[0].Name != gateway2.Name {
		t.Errorf("Expected %v, got %v", gateway1[0].Name, gateway2.Name)
	}
}

func TestUpdateGatewayByName(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/gateway")
	gatewayAPI := &GatewayAPI{API: apiWrapper}

	gateways, err := gatewayAPI.GetAllGateways()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	testing_gateway := Gateway{}
	for _, gateway := range gateways {
		if strings.Contains(gateway.Name, "testing-api-") {
			testing_gateway = gateway
		}
	}

	updated_gateway, err := gatewayAPI.UpdateGatewayByName(testing_gateway.Name, true, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if updated_gateway.IsPrivate == testing_gateway.IsPrivate {
		t.Errorf("Expected %v, got %v", updated_gateway.IsPrivate, testing_gateway.IsPrivate)
	}
}

func TestDeleteGateway(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/ipfs/rest/v1/gateway")
	gatewayAPI := &GatewayAPI{API: apiWrapper}

	gateways, err := gatewayAPI.GetAllGateways()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	testing_gateway := ""
	for _, gateway := range gateways {
		if strings.Contains(gateway.Name, "testing-api-") {
			testing_gateway = gateway.Name
		}
	}

	err = gatewayAPI.DeleteGatewayByName(testing_gateway)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
