package destinations

import (
	"github.com/jmtx1020/go_quicknode/client"
	"os"
	"testing"
)

func TestCreateDestination(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &DestinationAPI{API: apiWrapper}

	_, err := destinationAPI.CreateDestination(
		"testing-go-api",
		"https://us-central1-serious-truck-412423.cloudfunctions.net/function-1",
		"POST",
		"webhook",
		1,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetAllDestinationsIntegration(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

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
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")

	destinationAPI := &DestinationAPI{API: apiWrapper}

	allDestinations, err := destinationAPI.GetAllDestinations()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = destinationAPI.GetDestinationByID(allDestinations[0].ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDeleteDestinationByID(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &DestinationAPI{API: apiWrapper}

	allDestinations, err := destinationAPI.GetAllDestinations()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = destinationAPI.DeleteDestinationByID(allDestinations[0].ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
