package destinations

import (
	"fmt"
	"github.com/jmtx1020/go_quicknode/client"
	"testing"
)

func TestCreateDestination(t *testing.T) {
	apiToken := "QN_5f3e75f0de08436087af3b0191b3c6dd"

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &DestinationAPI{API: apiWrapper}

	destination, err := destinationAPI.CreateDestination(
		"testing-go-api",
		"https://us-central1-serious-truck-412423.cloudfunctions.net/function-1",
		"POST",
		"webhook",
		1,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	fmt.Println(destination)
}

func TestGetAllDestinationsIntegration(t *testing.T) {
	apiToken := "QN_5f3e75f0de08436087af3b0191b3c6dd"

	// Create a DestinationAPI instance with the real APIWrapper
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &DestinationAPI{API: apiWrapper}

	// Call the method being tested
	_, err := destinationAPI.GetAllDestinations()

	// Check for expected errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetDestinationByID(t *testing.T) {
	apiToken := "QN_5f3e75f0de08436087af3b0191b3c6dd"

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &DestinationAPI{API: apiWrapper}

	_, err := destinationAPI.GetDestinationByID("bfb935e9-b87b-4f99-ab62-f22f249d30c6")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDeleteDestinationByID(t *testing.T) {
	apiToken := "QN_5f3e75f0de08436087af3b0191b3c6dd"

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &DestinationAPI{API: apiWrapper}

	err := destinationAPI.DeleteDestinationByID("bfb935e9-b87b-4f99-ab62-f22f249d30c6")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
