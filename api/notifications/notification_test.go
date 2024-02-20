package notifications

import (
	"fmt"
	"github.com/jmtx1020/go_quicknode/api/destinations"
	"github.com/jmtx1020/go_quicknode/client"
	"os"
	"testing"
)

func TestCreateNotification(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	destinationAPI := &destinations.DestinationAPI{API: apiWrapper}
	allDestinations, err := destinationAPI.GetAllDestinations()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	_, err = notificationAPI.CreateNotification(
		"Test Notification",
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		"ethereum-mainnet",
		[]string{allDestinations[0].ID},
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetAllNotificationsIntegration(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	_, err := notificationAPI.GetAllNotifications()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetNotificationByID(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	allNotifications, err := notificationAPI.GetAllNotifications()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = notificationAPI.GetNotificationByID(allNotifications[0].ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDeleteNotificationByID(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	allNotifications, err := notificationAPI.GetAllNotifications()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = notificationAPI.DeleteNotificationByID(allNotifications[0].ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestUpdateNotificationByID(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")

	destinationAPI := &destinations.DestinationAPI{API: apiWrapper}
	allDestinations, err := destinationAPI.GetAllDestinations()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")
	not1, err := notificationAPI.CreateNotification(
		"Test Notification - 1",
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		"ethereum-mainnet",
		[]string{allDestinations[0].ID},
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	not2, err := notificationAPI.UpdateNotificationByID(
		not1.ID,
		"Test Notification - 2",
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		[]string{allDestinations[0].ID},
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = notificationAPI.DeleteNotificationByID(not2.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestToggleNotificationByID(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")

	// first create a destination
	destinationAPI := &destinations.DestinationAPI{API: apiWrapper}
	dest, err := destinationAPI.CreateDestination(
		"testing-go-api",
		"https://us-central1-serious-truck-412423.cloudfunctions.net/function-1",
		"POST",
		"webhook",
		1,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// create a notification, default behavior is enabled
	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")
	not1, err := notificationAPI.CreateNotification(
		"Test Notification",
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		"ethereum-mainnet",
		[]string{dest.ID},
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = notificationAPI.ToggleNotificationByID(not1.ID, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	not2, err := notificationAPI.GetNotificationByID(not1.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result := not2.Enabled == false
	if !result {
		t.Errorf("Expected %v, got %v", false, result)
	}

	// clean up
	err = notificationAPI.DeleteNotificationByID(not2.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetNotificationHistory(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")
	apiWrapper := client.NewAPIWrapper(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/destinations")
	// first create a destination
	destinationAPI := &destinations.DestinationAPI{API: apiWrapper}
	dest, err := destinationAPI.CreateDestination(
		"testing-go-api",
		"https://us-central1-serious-truck-412423.cloudfunctions.net/function-1",
		"POST",
		"webhook",
		1,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// create a notification in case one doesn't exist
	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	not1, err := notificationAPI.CreateNotification(
		"Test Notification",
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		"ethereum-mainnet",
		[]string{
			dest.ID,
		},
	)

	_, err = notificationAPI.GetNotificationEventsByID(not1.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// clean up
	err = notificationAPI.DeleteNotificationByID(not1.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestValidateExpression(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	err := notificationAPI.ValidateExpression(
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		"ethereum-mainnet",
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestBackTestExpression(t *testing.T) {
	apiToken := os.Getenv("QUICKNODE_API_TOKEN")

	notificationAPI := NewNotificationAPI(apiToken, "https://api.quicknode.com/quickalerts/rest/v1/notifications")

	bt_result, err := notificationAPI.BackTestExpression(
		"dHhfdG8gPT0gJzB4ZDhkQTZCRjI2OTY0YUY5RDdlRWQ5ZTAzRTUzNDE1RDM3YUE5NjA0NSc=",
		"ethereum-mainnet",
		12345678,
		10,
	)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	fmt.Println(bt_result)
}
